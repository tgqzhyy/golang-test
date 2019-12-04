/**
安装go-sql-driver/mysql包装
Go编程语言附带一个名为`database / sql`的便捷软件包，用于查询各种SQL数据库。这很有用，因为它将所有通用SQL功能抽象为一个API供您使用。Go不包括的是数据库驱动程序。在Go中，数据库驱动程序是一个软件包，用于实现特定数据库（在我们的情况下为MySQL）的低级详细信息。您可能已经猜到了，这对于保持向前兼容很有用。由于在创建所有Go软件包时，作者无法预见每个数据库将来都会投入使用，而支持每个可能的数据库将需要进行大量维护工作。

要安装MySQL数据库驱动程序，请转到您选择的终端并运行：

go get -u github.com/go-sql-driver/mysql
连接到MySQL数据库
安装所有必需的软件包后，我们需要检查的第一件事是，是否可以成功连接到MySQL数据库。如果您尚未运行MySQL数据库服务器，则可以使用Docker轻松启动新实例。以下是Docker MySQL映像的官方文档：https://hub.docker.com/_/mysql

要检查我们是否可以连接到数据库，请导入database/sql和go-sql-driver/mysql包，然后按如下所示打开连接：
import "database/sql"
import _ "go-sql-driver/mysql"


// Configure the database connection (always check errors)
db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")



// Initialize the first connection to the database, to see if everything works correctly.
// Make sure to check the error.
err := db.Ping()
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234567890@(127.0.0.1:3306)/ppgo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
          CREATE TABLE users (
              id INT AUTO_INCREMENT,
              username TEXT NOT NULL,
              password TEXT NOT NULL,
              created_at DATETIME,
              PRIMARY KEY (id)
          );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
		username := "johndoe1"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v\n", users)
	}

	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

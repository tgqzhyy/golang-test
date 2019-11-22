package mymysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbSelect *sql.DB

func Conn() *sql.DB {
	return dbSelect
}
func init() {
	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")

	dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	log.Println(dbcon)
	/*"root:root@tcp(localhost:3306)/test?charset=utf8"*/
	db,err :=sql.Open("mysql",dbcon)
	if err !=nil{
		panic(err)
	}
	if err:= db.Ping();err !=nil{
		panic(err)
	}
	dbSelect=db
}

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "golang-test/毒鸡汤/djt/routers"
	"time"
)
type Store struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}


func init() {

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")

	dbcon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	//fmt.Print(dbcon)
	orm.RegisterDataBase("default", "mysql", dbcon /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式
}
func main() {
	beego.Run()
}


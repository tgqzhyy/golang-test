package main

import (
	"github.com/astaxie/beego"
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

func main() {
	beego.Run()
}

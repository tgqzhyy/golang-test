package controllers

import (
	"github.com/astaxie/beego"
	"golang-test/毒鸡汤/djt/models/mymysql"
	"log"
)

type MainController struct {
	beego.Controller
}

var title string

func (c *MainController) Get() {
	query := "select title from soul order by rand( ) limit 1"
	dbSelect := mymysql.Conn()
	//defer dbSelect.Close()
	if err := dbSelect.QueryRow(query).Scan(&title); err != nil {
		log.Fatal(err)
	}
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "dujitang.tpl"
	c.Data["Statistical_code"] = beego.AppConfig.String("Statistical_code")
	c.Data["djtURL"] = beego.AppConfig.String("djtURL")
	//c.Data["djt"]="这是测试毒鸡汤"
	c.Data["djt"] = title

	//	<?php
	//	$sql="select * from soul order by rand( ) limit 1";
	//	$rs=mysql_query($sql);
	//	?>
	//	<?php
	//	while($rows=mysql_fetch_assoc($rs))
	//	{
	//	?>
	//
	//	<span id="sentence" style="font-size: 2rem;"><?php echo $rows["title"]?></span>
	//	<?php
	//	}
	//	?>
}

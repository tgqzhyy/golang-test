package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "dujitang.tpl"
	c.Data["Statistical_code"]=beego.AppConfig.String("Statistical_code")
	c.Data["djtURL"]=beego.AppConfig.String("djtURL")

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

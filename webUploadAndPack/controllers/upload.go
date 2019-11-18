package controllers

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Website"] = beego.AppConfig.String("Website")
	c.Data["Email"] = beego.AppConfig.String("Email")
	c.TplName = "upload.tpl"
}


func (c *UploadController) Post() {
	c.TplName="upload.tpl"
	f,h,err :=c.GetFile("file")
	if err !=nil{
		log.Fatal("getFile err:",err)
		c.Ctx.Redirect(302,"/erro")
		return
	}else {
		fileName :=h.Filename
		ExeFile :=strings.Split(fileName,".")
		layout :=strings.ToLower(ExeFile[len(ExeFile)-1])
		if layout !="exe" && layout !="dll"{
			c.Ctx.WriteString("请上传符合格式的内容(exe、dll)")
			return
		}
		fileNamePack := beego.AppConfig.String("SavaDiskPath")+strconv.FormatInt(time.Now().UnixNano(),10)+h.Filename
		err =c.SaveToFile("file",fileNamePack) //Save disk path
		if err!=nil {
			//c.Ctx.WriteString("File upload failed!")
			c.Ctx.Redirect(302,"/erro")
		}else {
			//c.Ctx.WriteString("File upload succed!")
			c.Ctx.Redirect(302,"/file")
			go ttt()
		}
	}
	defer f.Close()

}

//测试可以进行go程
func ttt() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1","A1","姓名")
	xlsx.SetCellValue("Sheet1","B1","年龄")
	xlsx.SetCellValue("Sheet1","A2","钩子")
	xlsx.SetCellValue("Sheet1","B2","199")
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err :=xlsx.SaveAs("MyXLSXFile.xlsx")
	if err != nil{
		fmt.Println(err)
	}
}

//func  AddPack(f string) {
//	cmd :=exec.Command("ls","-a","-l",">>","bb.txt")
//	stdout,err :=cmd.StdoutPipe()
//	if err !=nil{
//		log.Fatal(err)
//	}
//	defer  stdout.Close()
//	if err := cmd.Start();err != nil{
//		log.Fatal(err)
//		log.Fatal(f)
//	}
//
//	opBytes, err := ioutil.ReadAll(stdout)
//	if err != nil{
//		log.Fatal(err)
//	}
//	log.Println(string(opBytes))
//}

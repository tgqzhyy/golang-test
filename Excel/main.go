package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f :=excelize.NewFile()
	//创建一个工作表
	index :=f.NewSheet("fuck")
	//设置单元格格式
	f.SetCellValue("fuck","A2","这是一个test")
	f.SetCellValue("Sheet1","B2",100)
	//设置工作薄的默认工作表
	f.SetActiveSheet(index)
	//根据指定路径保存文件
	err :=f.SaveAs("./Book.xlsx")
	if err !=nil{
		fmt.Println(err)
	}
}

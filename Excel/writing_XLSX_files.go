package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet, err = file.AddSheet("Sheet2")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet, err = file.AddSheet("中文测试")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow() //增加行
	row.SetHeightCM(1) //设置每行的高度
	cell = row.AddCell() //增加列
	cell.Value = "I am a cell!" //设置内容
	cell =row.AddCell()
	cell.Value = "I am a cell444!"
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "2I am a cell!"
	cell =row.AddCell()
	cell.Value = "2I am a cell444!"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

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
	{
		// TOTO：貌似没有办法设置列宽
		// Sheet1 内容
		row = sheet.AddRow() //增加行
		row.SetHeightCM(3)   //设置每行的高度
		cell = row.AddCell() //增加列
		//style := xlsx.NewStyle()
		//style.Fill.FgColor = "FFFFFF00"
		////style.Fill.BgColor = "FFFFFF00"
		//style.Fill.PatternType = "solid"
		//cell.SetStyle(style)
		style := xlsx.NewStyle()
		style.Font.Bold = true
		style.Font.Color = "FF4DFF6B" //设置字体颜色
		style.Font.Name = "宋体"
		style.Font.Size = 18
		//style.Font.Family=2
		style.Alignment.WrapText = true //是否自动换行
		style.Alignment.Horizontal = "Center"
		cell.SetStyle(style)
		cell.Value = "I am a cell!" //设置内容
		cell = row.AddCell()
		cell.Value = "I am a cell444!"
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "2I am a cell!"
		cell = row.AddCell()
		cell.SetStyle(style)
		cell.Value = "2I am a cell444!"
	}
	sheet, err = file.AddSheet("Sheet2")
	if err != nil {
		fmt.Printf(err.Error())
	}
	sheet, err = file.AddSheet("中文测试")
	if err != nil {
		fmt.Printf(err.Error())
	}

	{
		// 中文测试表内容
		row = sheet.AddRow() //增加行
		row.SetHeightCM(2) //设置每行的高度
		cell = row.AddCell() //增加列
		cell.Value = "I am a cell!" //设置内容
		cell =row.AddCell()
		cell.Value = "I am a cell444!"
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "2I am a cell!"
		cell =row.AddCell()
		cell.Value = "2I am a cell444!"
	}
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

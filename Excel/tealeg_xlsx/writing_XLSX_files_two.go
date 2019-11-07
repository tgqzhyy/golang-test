package tealeg_xlsx

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"

	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"

	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

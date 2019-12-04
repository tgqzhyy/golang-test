package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "姓名")
	xlsx.SetCellValue("Sheet1", "B1", "年龄")
	xlsx.SetCellValue("Sheet1", "A2", "钩子")
	xlsx.SetCellValue("Sheet1", "B2", "199")
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell, _ := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	rows, _ := xlsx.GetRows("Sheet1")

	for _, row := range rows {
		for _, colcell := range row {
			fmt.Print(colcell, "\t")
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName :="MyXLSXFile.xlsx"
	xlFile,err := xlsx.OpenFile(excelFileName)
	if err!=nil {
		fmt.Printf("open failed:%s \n",err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet name:%s \n",sheet.Name)
		for _, row := range sheet.Rows{
			for _, cell := range  row.Cells{
				text := cell.String()
				fmt.Printf("%s\n",text)
			}
		}

	}
}

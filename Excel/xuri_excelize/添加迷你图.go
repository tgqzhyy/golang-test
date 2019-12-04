package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()
	e := f.AddSparkline("Sheet1", &excelize.SparklineOption{
		Location: []string{"A1", "A2", "A3"},
		Range:    []string{"Sheet2!A1:J1", "Sheet2!A2:J2", "Sheet2!A3:J3"},
		Markers:  true,
	})
	if e != nil {
		fmt.Println(e)
	}
	//根据指定路径保存文件
	err := f.SaveAs("./Book2.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

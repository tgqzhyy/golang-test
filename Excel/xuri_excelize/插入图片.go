package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	err = xlsx.AddPicture("Sheet1", "A2", "Excel/s7c15ctkgd.gif", "")
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	err = xlsx.AddPicture("Sheet1", "D2", "Excel/s7c15ctkgd.gif", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	err = xlsx.AddPicture("Sheet1", "H2", "Excel/s7c15ctkgd.gif", `{"x_offset": 15, "y_offset": 10, "print_obj": true, "lock_aspect_ratio": false, "locked": false}`)
	if err != nil {
		fmt.Println(err)
	}
	// Save the xlsx file with the origin path.
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}
}

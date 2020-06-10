package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"time"
)

/**
go get -u github.com/gen2brain/dlgs
dlgs is a cross-platform library for displaying dialogs and input boxes.
more information https://godoc.org/github.com/gen2brain/dlgs
*/
func main() {
	color, _, err := dlgs.Color("Pick color", "#BEBEBE")
	if err != nil {
		panic(err)
	}
	fmt.Println(color)

	date, _, err := dlgs.Date("Date", "Date of traveling:", time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(date)

	entry, _, err := dlgs.Entry("Entry", "Enter something here, anything:", "default text")
	if err != nil {
		panic(err)
	}
	fmt.Println(entry)

	_, err = dlgs.Error("Error", "Cannot divide by zero.")
	if err != nil {
		panic(err)
	}

	f, _, err := dlgs.File("Select file", "", false)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	// List 选择框
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)

	// Password 密码框
	passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)

	// Question 问答框
	yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(yes)
}

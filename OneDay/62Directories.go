package main

import (

	"os"
)

func check(e error) {
	if e !=nil{
		panic(e)
	}
}

func main() {
	err := os.Mkdir("subdir",0755)
	check(err)
}

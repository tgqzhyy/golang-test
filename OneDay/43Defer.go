package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tem/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func writeFile(f3 interface{}) {

}

func closeFile(f3 interface{}) {

}

func createFile(i string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(i)
	if err != nil {
		panic(err)
	}
	return f
}

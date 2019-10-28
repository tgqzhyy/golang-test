package main

import (
	"fmt"
	"io/ioutil"
	"os"

)

func check(e error) {
	if e !=nil{
		panic(e)
	}
}

func main() {
	f,err := ioutil.TempFile("","sample")
	check(err)

	fmt.Println("Temp file name:",f.Name())

	defer os.Remove(f.Name())


}

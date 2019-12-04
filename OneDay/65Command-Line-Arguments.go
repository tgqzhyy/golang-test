package main

import (
	"fmt"
	"os"
)

/**
go build 65Command-Line-Arguments.go
./65Command-Line-Arguments a b cd aad
[./65Command-Line-Arguments a b cd aad]
[a b cd aad]
cd

*/
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

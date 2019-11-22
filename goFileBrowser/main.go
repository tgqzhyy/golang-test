package main

import (
	"runtime"

	"golang-test/goFileBrowser/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}

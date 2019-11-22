package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err !=nil {
		panic(err)
	}
}
/**
panic: a problem

goroutine 1 [running]:
main.main()
        /home/landv/go/src/golang-test/OneDay/42Panic.go:6 +0x39

Process finished with exit code 2

 */

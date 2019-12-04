package main

import "fmt"

func zeroval(ival int) {
	ival = 4
	fmt.Println("内部：", ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 5
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

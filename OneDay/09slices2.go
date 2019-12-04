package main

import "fmt"

func main() {
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	a[0][0] = 9
	fmt.Println(a)
}

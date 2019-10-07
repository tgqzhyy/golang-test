package main

import "fmt"

//闭包
func intSeq() func() int{
	i :=0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt :=intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts :=intSeq()
	fmt.Println(newInts())
}
/**

Go支持匿名函数，可以形成闭包。当您想要内联定义函数而不必命名时，匿名函数很有用。
此函数intSeq返回另一个函数，我们在的正文中匿名定义intSeq。返回的函数关闭变量i以形成闭包。
我们调用intSeq，将结果（一个函数）分配给nextInt。此函数值捕获其自己的i值，每次调用时都会更新该值nextInt。
通过调用nextInt 几次来查看关闭的效果。
要确认状态对于该特定功能是唯一的，请创建并测试一个新功能。
1
2
3
1


*/

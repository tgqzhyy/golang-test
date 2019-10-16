package main

import "fmt"
//该ping功能仅接受用于发送值的通道。尝试在此通道上接收将是编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}
//该pong函数接受一个通道用于接收（pings），第二个通道用于发送（pongs）。
func pong(pings <-chan string, pongs chan<- string)  {
	msg :=<-pings
	pongs <- msg
}

func main() {
	pings :=make(chan string,1)
	pongs :=make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}

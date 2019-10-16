/**
Channel 是golang中的一个核心类型，你可以把它看成给一个管道，通过它发送核心单元就可以发送或者接受数据进行通讯
它的操作符是箭头<-
ch <- v //发送值V到Channel ch中
v := <-ch //冲channel ch中接受数据，并将数据赋值给V
箭头的指向就是数据的流向
就像map和slice数据类型一样channel必须先创建在使用：
ch :=make(chan int)

 */
package main

import "fmt"

func main() {
	messages :=make(chan string)

	go func() {messages <-"ping"}()

	msg :=<-messages
	fmt.Println(msg)
}

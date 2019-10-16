package main

import (
	"fmt"
	"time"
)
//这是我们将在goroutine中运行的函数。该 done通道将用于通知另一个goroutine该功能的工作已完成。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	//发送一个值以通知我们已完成。
	done <- true
}

func main() {
	//启动工作程序goroutine，为其提供通知的通道。
	done := make(chan bool,1)
	go worker(done)
	//封锁，直到我们在频道上收到工作人员的通知为止。
	//如果<- done从该程序中删除该行，则该程序将在worker偶数启动之前退出。
	 <-done
}

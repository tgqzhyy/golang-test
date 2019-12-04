package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动发起连接请求
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("dial err", err)
		return
	}
	defer conn.Close()
	//启动go进程，持续接收用户键盘输入
	go func() {
		//创建缓冲区
		buf := make([]byte, 1024)
		for {
			//持续接收放进缓冲区buf中
			n, err := os.Stdin.Read(buf)
			//错误处理
			if err != nil {
				fmt.Println("stdin err", err)
				return
			}
			//接收多少，写多少给服务器
			_, err = conn.Write(buf[:n])
			if err != nil {
				fmt.Println("write err", err)
				return
			}
		}
	}()
	//主进程接收服务器信息
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		//没有信息时，说明断开了连接,则主动结束进程
		if n == 0 {
			fmt.Println("与服务器连接断开")
			break
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		//打印接收到的信息
		fmt.Println(string(buf[:n]))
	}
}

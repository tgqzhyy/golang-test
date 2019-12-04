package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

//这是一个用来处理连接进来的用户的函数
func handleconn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	//给客户端发送连接成功的信号
	conn.Write([]byte("welcome to my server 中文测试\n"))
	for {
		//持续读取客户端数据，保存在buf缓冲区中，并处理
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("outtoleave", conn.RemoteAddr().String())
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("read err", err)
			return
		}
		//这里是处理数据的一个示范，把客户端发来的数据全部转化为大写
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//图片的第一步，net.listen开启监听，获得listener套接字
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	for {
		//开启图片中的第二步，阻塞等待用户连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		//每一个连接进来的用户都为其分配一个专属的go进程
		go handleconn(conn)
	}
}

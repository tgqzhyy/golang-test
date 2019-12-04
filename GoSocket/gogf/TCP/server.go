package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
)

/***
在这个示例中我们使用了Send和Recv来发送和接受数据。其中Recv方法会通过阻塞方式接受数据，
直到客户端“发送完毕一条数据”（执行一次Send，底层Sockettongue不带缓冲实现），或者关闭链接。
*/
func main() {
	gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()

		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte(">"), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()
}

/***
执行之后我们使用telnet工具来进行测试：
landv@landv-PC:~$ telnet 127.0.0.1 8999
Trying 127.0.0.1...
Connected to 127.0.0.1.
Escape character is '^]'.
aaa
>aaa
aaa
>aaa
ccc
>ccc
zhongg
>zhongg
每一个客户端发起的TCP链接，TCPServer都会创建一个goroutine进行处理，直到TCP链接断开。
由于goroutine比较轻量级，因此可以支撑很高的并发量。
*/

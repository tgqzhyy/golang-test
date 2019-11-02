/**
https://www.jianshu.com/p/dbc62a879081
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

//数据包的类型
const(
	HEART_BEAT_PACKET =0x00
	REPORT_PACKET=0X01
)

var(
	server= "127.0.0.1:8081"
)
//这里是包的结构体
type Packet struct {
	PacketType 	byte
	PacketContent	[]byte
}

//心跳包，这里用json来序列化，也可以用github上的gogo/probobuf包
//具体见(https://github.com/gogo/protobuf)
type HeartPacket struct {
	Version string`json:"version"`
	Timestamp int64`json:"timestamp"`
}
//正式上传的数据包
type ReportPacket struct {
	Content string`json:"content"`
	Rand 	int`json:"rand"`
	Timestamp int64`json:"timestamp"`
}
//与服务器相关的资源都放在这里面
type TcpServer struct {
	listener *net.TCPListener
	hawkServer *net.TCPAddr
}

func checkErr(err error) {
	if err !=nil{
		fmt.Println(err)
		os.Exit(-1)
	}
}
func main() {
	//类似与初始化套接字，绑定端口
	hawkServer,err :=net.ResolveTCPAddr("tcp",server)
	checkErr(err)
	//监听
	listen,err :=net.ListenTCP("tcp",hawkServer)
	checkErr(err)
	defer listen.Close()
	tcpServer :=&TcpServer{
		listener:   listen,
		hawkServer: hawkServer,
	}
	fmt.Println("start server successfull....")
	//开始接收请求
	for{
		conn,err :=tcpServer.listener.Accept()
		fmt.Println("accept tcp client %s",conn.RemoteAddr().String())
		checkErr(err)
		//每次建立一个连接就放到单独的协程内做处理
		go Handle(conn)
	}
}
//处理函数，这是一个状态机
//根据数据包来做解析
//数据包的格式为|0xFF|0xFF|len(高)|len(低)|Data|CRC高16位|0xFF|0xFE
//其中len为data的长度，实际长度为len(高)*256+len(低)
//CRC为32位CRC，取了最高16位共2Bytes
//0Xff|0xFF和0xFF|0xFE类似于前导码
func Handle(conn net.Conn) {
	// close connection before exit
	defer conn.Close()
	// 状态机状态
	state :=0x00
	// 数据包长度
	length :=uint16(0)
	// crc校验和
	crc16 :=uint16(0)
	var recvBuffer []byte
	// 游标
	cursor := uint16(0)
	bufferReader :=bufio.NewReader(conn)

	//状态机处理数据
	for{
		recvByte,err :=bufferReader.ReadByte()
		if err !=nil{
			//这里因为做了心跳，所以就没有加deadline时间，如果客户端断开连接
			//这里ReadByte方法返回一个io.EOF的错误，具体可以考虑文档
			if err ==io.EOF{
				fmt.Printf("client %s is close!\n",conn.RemoteAddr().String())
			}
			//在这里直接退出goroutine,关闭有defer操作完成
			return
		}
		//进入状态机，根据不同的状态来处理
		switch state{
		case 0x00:
			if recvByte==0xFF{
				state =0x01
				//初始化状态机
				recvBuffer =nil
				length=0
				crc16 =0
			}else {
				state =0x00
			}
			break
		case 0x01:
			if recvByte == 0xFF{
				state = 0x02
			}else {
				state = 0x00
			}
			break
		case 0x02:
			length += uint16(recvByte) * 256
			state = 0x03
			break
		case 0x03:
			length += uint16(recvByte)
		}

	}

}

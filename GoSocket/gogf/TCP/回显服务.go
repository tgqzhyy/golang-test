package main

import (
	"fmt"
	"gitee.com/ha666/golibs/gtime"
	"github.com/gogf/gf/net/gtcp"
	"github.com/wudaoluo/goutil/glog"
	"time"
)

func main() {
	// Server
	go gtcp.NewServer("localhost:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for{
			data, err :=conn.Recv(-1)
			if len(data)>0{
				if err := conn.Send(append([]byte("S:> "),data...));err !=nil{
					fmt.Println(err)
				}
			}
			if err !=nil{
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	// Client

	for{
		if conn,err :=gtcp.NewConn("localhost:8999"); err ==nil{
			if b,err :=conn.SendRecv([]byte(gtime.Datetime()),-1); err ==nil{
				fmt.Println(string(b),conn.LocalAddr(),conn.RemoteAddr())
			}else {
				fmt.Println(err)
			}
			conn.Close()
		}else {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}

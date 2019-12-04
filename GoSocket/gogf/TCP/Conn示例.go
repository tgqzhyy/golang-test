package main

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/util/gconv"
	"github.com/wudaoluo/goutil/glog"
	"time"
)

func main() {
	// Server
	go gtcp.NewServer("localhost:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				fmt.Println("S:", time.Now(), string(data))
			}
			if err != nil {
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	conn, err := gtcp.NewConn("localhost:8999")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}

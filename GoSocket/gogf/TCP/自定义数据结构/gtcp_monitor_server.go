package main

import (
	"encoding/json"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"golang-test/GoSocket/gogf/TCP/自定义数据结构/types"
)

func main() {
	// 服务端，接受客户端数据并格式化为指定数据结构，打印
	glog.Println("############SERVER#############")
	gtcp.NewServer("localhost:8999", func(conn *gtcp.Conn) {

		defer conn.Close()
		for{
			data,err :=conn.RecvPkg()
			if err !=nil{
				if err.Error() =="EOF"{
					glog.Println("client closed")
				}
				break
			}
			info := &types.NodeInfo{}
			if err :=json.Unmarshal(data,info); err !=nil{
				glog.Errorf("invalid package structrue:%s \n",err.Error())
			}else {
				glog.Println(info)
				conn.SendPkg([]byte("ok"))
			}
		}
	}).Run()
}

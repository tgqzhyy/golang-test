package main

import (
	"encoding/json"
	"gitee.com/ha666/golibs/gtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"golang-test/GoSocket/gogf/TCP/自定义数据结构/types"
)

func main() {
	// 数据上报客户端
	conn, err := gtcp.NewConn("localhost:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 试用JSON格式化数据字段
	info, err := json.Marshal(types.NodeInfo{
		Cpu:  float32(66.66),
		Host: "localhost",
		Ip: g.Map{
			"eth0": "192.168.1.100",
			"eth1": "114.114.10.11",
		},
		MemUsed:  15560320,
		MemTotal: 16333788,
		Time:     int(gtime.Second()),
	})
	if err != nil {
		panic(err)
	}
	// 使用 SendRecvPkg 发送消息包并接受返回
	if result, err := conn.SendRecvPkg(info); err != nil {
		if err.Error() == "EOF" {
			glog.Println("server closed")
		}
	} else {
		glog.Println(string(result))
	}
}

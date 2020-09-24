package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

/**
需要用的库https://github.com/kardianos/service
start, stop, restart, install, uninstall 可用于操作服务
sudo fuckRegisServicesLinux   install   -- 安装服务
sudo fuckRegisServicesLinux   start     -- 启动服务
sudo fuckRegisServicesLinux   stop     -- 停止服务

service fuckRegisServicesLinux start
*/
var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Println("开始服务")
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	//fmt.Println("先打印一行文字再说")
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Println("停止服务")
	return nil
}

func main() {
	// 服务的配置信息
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}
	// Interface 接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}

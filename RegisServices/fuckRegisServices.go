package main

import (
	"github.com/kardianos/service"
	"log"
	"os"
)

var logging service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	logging.Info("开始服务")

	go p.run()
	return nil
}
func (p *program) Stop(s service.Service) error {
	logging.Info("停止服务")
	return nil
}
func (p *program) run() {
	// 这里放置程序要执行的代码……
	logging.Info("run my ccc")
}

func main() {
	//服务的配置信息
	cfg := &service.Config{
		Name:        "simple_test",
		DisplayName: "a simple_test service",
		Description: "This is an simple_test Go service.",
	}
	// Interface 接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	errs := make(chan error, 5)
	logging, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(os.Args) == 2 { //如果有命令则执行
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else { //否则说明是方法启动了
		err = s.Run()
		if err != nil {
			logging.Error(err)
		}
	}
	if err != nil {
		logging.Error(err)
	}
}

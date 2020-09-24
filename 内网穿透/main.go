package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const APP_VERSION = "20200522"

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal("err : %s", err)
	}
	Clean(&service.S)
}
func Clean(s *services.Service) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP, //Linux Signal及Golang中的信号处理,https://colobu.com/2015/10/09/Linux-Signals/
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt,stoping services...")
			(*s).Clean()
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}

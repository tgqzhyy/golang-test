package main

import (
	"log"
	"os"
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

}

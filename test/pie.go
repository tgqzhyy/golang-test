package main

import (
	"time"

	"github.com/ha666/golibs"
	"github.com/ha666/logs"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 300)
		logs.Info(golibs.GetPwd(32))
	}
}

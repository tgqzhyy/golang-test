package main

import (
	"fmt"
	"github.com/KyleBanks/goggles/server"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/core/errgroup"
	"github.com/kataras/iris/sessions"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	gin.SetMode("debug")
	config := cfg.instance()
	if config {
		fmt.Printf("参数错误")
		return
	}
	// 后台服务状态
	adminStatus := config.Status.Admin
	// api服务状态
	apiStatus := config.Status.Api

	if adminStatus {
		store := memstore.NewStore([]byte("secret"))
		admin := server.New("admin", config.Admin.Address, gin.Logger(), sessions.Sessions("mysession", store))
		admin.Template("template").Static(config.Admin.ServerRoot)
		admin.Start(g)
	}

	if apiStatus {
		api := server.New("api", config.Api.Address, gin.Recovery(), gin.Logger())
		api.Start(g)
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
	var state int32 = 1
	sc := make(chan os.Signal, 1)                                                       // 创建监听退出chan
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // 监听指定信号
	// SIGHUP:终端控制进程结束（终端连接断开）；SIGINT：用户发送INTR字符串（Ctrl+C）触发
	// SIGTERM:结束程序（可以被捕获、阻塞或忽略）；SIGQUIT：用户发送QUIT字符（Ctrl+/）触发
EXIT:
	for {
		sig := <-sc
		fmt.Println("获取到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0) //原子操作,函数atomic.StoreInt32会接受两个参数。第一个参数的类型是*int 32类型的，其含义同样是指向被操作值的指针。而第二个参数则是int32类型的，它的值应该代表欲存储的新值。其它的同类函数也会有类似的参数声明列表。
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))

}

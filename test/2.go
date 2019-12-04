package main

import "log"



func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("test:","http://landv.cnblogs.com")
	//2019/12/02 10:44:10 test: http://landv.cnblogs.com 不加SetFlags
	//2019/12/02 10:44:28 2.go:7: test: http://landv.cnblogs.com 加SetFlags

}

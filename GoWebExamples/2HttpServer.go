package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	/**
	服务静态资产
	为了提供JavaScript，CSS和图片等静态资源，我们使用内置的http.FileServer并将其指向url路径。为了使文件服务器正常工作，需要知道从何处提供文件。我们可以这样做
	*/
	fs := http.FileServer(http.Dir("static/"))
	/**
	一旦我们的文件服务器到位，我们只需要指向它的url路径，就像处理动态请求一样。需要注意的一件事：为了正确提供文件，我们需要剥离一部分url路径。通常，这是我们文件所在目录的名称。
	*/
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Print("http://localhost:8091")
	http.ListenAndServe(":8091", nil)
}

/**
landv@landv-PC:~/go/src/golang-test/GoWebExamples$ go run 2HttpServer.go
2019/10/25 08:50:26 http://localhost:8091

*/

/**
fuck,这种简单的代码，居然运行不起来，检查了N多遍代码输入错误问题，结果重启一遍goland就能运行。
当什么都检查不出来的时候，请保存重启，然后抽根烟冷静一下。
这个是缓存的问题，重启就解决。
Go的html/template软件包为HTML模板提供了丰富的模板语言。它主要用于Web应用程序中，以结构化的方式在客户端的浏览器中显示数据。Go的模板语言的一大优势是自动转义数据。无需担心XSS攻击，因为Go会解析HTML模板并在将其显示给浏览器之前转义所有输入。

 */
package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	//
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     []Todo{
				{Title:"Task1",Done:false},
				{Title:"Task2",Done:true},
				{Title:"Task3",Done:true},
			},
		}
		tmpl.Execute(writer,data)
	})
	log.Print("http://localhost:8090")
	http.ListenAndServe(":8090",nil)



}

//package main
//
//import (
//	"html/template"
//	"net/http"
//)
//
//type Todo struct {
//	Title string
//	Done  bool
//}
//
//
//
//type TodoPageData struct {
//	PageTitle string
//	Todos     []Todo
//}
//
//func tmpll(w http.ResponseWriter, r *http.Request) {
//	tl, err :=template.ParseFiles("layout.html")
//	if err !=nil{
//		panic(err)
//	}
//	data := TodoPageData{
//		PageTitle: "My TODO list",
//		Todos:     []Todo{
//			{Title:"Task1",Done:false},
//			{Title:"Task2",Done:true},
//			{Title:"Task3",Done:true},
//		},
//	}
//	tl.Execute(w,data)
//}
//
//func main() {
//	server :=http.Server{
//		Addr:"127.0.0.1:8080",
//	}
//
//	http.HandleFunc("/",tmpll)
//	server.ListenAndServe()
//}

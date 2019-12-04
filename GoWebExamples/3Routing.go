package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/**
Go的net/http软件包为HTTP协议提供了许多功能。它做得不好的一件事是复杂的请求路由，例如将请求url分割成单个参数。幸运的是，有一个非常流行的软件包，以Go社区中良好的代码质量而闻名。在此示例中，您将看到如何使用gorilla/mux包创建具有命名参数，GET / POST处理程序和域限制的路由。
安装gorilla/mux包装
gorilla/mux是适合Go的默认HTTP路由器的软件包。它具有许多功能，可以提高编写Web应用程序时的生产率。它也符合Go的默认请求处理程序签名func (w http.ResponseWriter, r *http.Request)，因此该包可以与其他HTTP库（例如中间件或现有应用程序）混合使用。使用go get命令从GitHub安装软件包，如下所示：

go get -u github.com/gorilla/mux
创建一个新的路由器
首先创建一个新的请求路由器。路由器是Web应用程序的主要路由器，以后将作为参数传递给服务器。它将接收所有HTTP连接，并将其传递给您将在其上注册的请求处理程序。您可以这样创建新的路由器：

r := mux.NewRouter()
注册请求处理程序
一旦有了新的路由器，就可以像往常一样注册请求处理程序。唯一的区别是，http.HandleFunc(...)您无需HandleFunc像调用那样在路由器上进行呼叫，就像这样：r.HandleFunc(...)。
URL参数
gorilla/mux路由器的最大优势是能够从请求URL中提取分段。例如，这是您的应用程序中的URL：

/books/go-programming-blueprint/page/10
该URL具有两个动态段：
书名（go-programming-blueprint）
页面（10）
要使请求处理程序与上述URL匹配，您可以使用URL模式中的占位符替换动态分段，如下所示：

r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    // get the book
    // navigate to the page
})
最后一件事是从这些段中获取数据。该包带有功能mux.Vars(r)它接受http.Request作为参数并返回地图上的节段。

func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    vars["title"] // the book title slug
    vars["page"] // the page
}
设置HTTP服务器的路由器
有没有想过在什么样nil的http.ListenAndServe(":80", nil)换货？它是HTTP服务器的主路由器的参数。默认情况下为nil，这意味着使用net/http软件包的默认路由器。要使用自己的路由器，请将替换为router nil 的变量r。

http.ListenAndServe(":80", r)
*/

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book :%s on page %s\n", title, page)
	})
	defer http.ListenAndServe(":8090", r)
	log.Print("http://localhost:8090")
}

/**
http://localhost:8090/books/22/page/88
You've requested the book :22 on page 88
*/

/**
该功能gorilla/mux路由器
方法
将请求处理程序限制为特定的HTTP方法。

r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

主机名和子域
将请求处理程序限制为特定的主机名或子域。

r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

方案
将请求处理程序限制为http / https。

r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

路径前缀和子路由器
将请求处理程序限制为特定的路径前缀。

bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
*/

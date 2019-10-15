package main

import "github.com/zserge/webview"
//各种骚操作，emmm这种还是比较适合我滴。直接嵌入本地化网站就好了。
func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example",
		"https://www.cnblogs.com/landv/", 800, 600, true)
}

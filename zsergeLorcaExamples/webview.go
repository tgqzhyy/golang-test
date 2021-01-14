package main

import "github.com/webview/webview"

//各种骚操作，emmm这种还是比较适合我滴。直接嵌入本地化网站就好了。
func main() {
	// Open wikipedia in a 800x600 resizable window
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://www.cnblogs.com/landv/")
	w.Run()
}

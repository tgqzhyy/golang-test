package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello 中文测试")
	w.SetContent(widget.NewLabel("Hello Fyne! 中文测试")) // TODO 需要设置字体位置

	w.ShowAndRun()
}

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)
/**
export FYNE_FONT=/usr/share/fonts/fonts-cesi/CESI_SS_GB18030.TTF
echo $FYNE_FONT
go run main.go
 */
var b fyne.Size
func main() {

	app := app.New()
	w := app.NewWindow("标题中文测试")

	w.CenterOnScreen()
	b.Height=150
	b.Width=300
	w.Resize(b)
	w.SetContent(widget.NewVBox(
		widget.NewLabel("内容中文测试"),
		widget.NewButton("按钮中文测试", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

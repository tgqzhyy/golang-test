package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)
var b fyne.Size
func main() {

	app := app.New()
	w := app.NewWindow("Hello")
	w.CenterOnScreen()
	b.Height=150
	b.Width=300
	w.Resize(b)
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

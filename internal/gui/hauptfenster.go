package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Test() {
	a := app.New()
	w := a.NewWindow("Hello world")

	w.SetContent(widget.NewLabel("Hello world"))
	w.ShowAndRun()

}

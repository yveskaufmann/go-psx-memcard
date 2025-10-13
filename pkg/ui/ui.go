package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Start() error {
	a := app.New()
	w := a.NewWindow("PSX Memory Card Manager")

	view := NewManagerWindowView(w)

	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(view.Container())
	w.ShowAndRun()

	return nil
}

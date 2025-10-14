package ui

import (
	"fyne.io/fyne/v2/app"
)

func Start() error {
	a := app.New()
	w := a.NewWindow("PSX Memory Card Manager")

	view := NewManagerWindowView(w)

	w.SetContent(view.Container())
	w.ShowAndRun()

	return nil
}

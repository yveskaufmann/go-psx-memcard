package ui

import (
	"com.yv35.memcard/pkg/dig"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func newApp() fyne.App {
	return app.NewWithID("com.yv35.PSXMemoryCardManager")
}

func newWindow(a fyne.App) fyne.Window {
	return a.NewWindow("PSX Memory Card Manager")
}

func Start() error {

	dig.Provide(newApp)
	dig.Provide(newWindow)
	dig.Provide(NewManagerWindowView)

	return dig.Invoke(func(window fyne.Window, view *ManagerWindowView) {
		window.SetContent(view.Container())
		window.ShowAndRun()
	})
}

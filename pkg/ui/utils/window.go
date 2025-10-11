package utils

import "fyne.io/fyne/v2"

func GetPrimaryWindow() fyne.Window {
	windows := fyne.CurrentApp().Driver().AllWindows()
	if len(windows) > 0 {
		return windows[0]
	}
	return nil
}

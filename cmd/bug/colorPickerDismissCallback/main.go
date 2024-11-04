package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"image/color"
)

func main() {
	fapp := app.NewWithID("app-with-unique-id")
	win := fapp.NewWindow("pick color")

	cp := dialog.NewColorPicker("", "", func(c color.Color) {
		fmt.Println("Color callback called with:", c)
		win.Close()
		fapp.Quit()
	}, win)

	cp.Show()
	win.Show()
	fapp.Run()
}

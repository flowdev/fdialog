package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"image/color"
	"os"
)

func main() {
	fapp := app.NewWithID("app-with-unique-id")
	win := fapp.NewWindow("pick color")
	win.Resize(fyne.NewSize(500, 500))

	cp := dialog.NewColorPicker("", "", func(c color.Color) {
		fmt.Println("Color callback called!")
		win.Close()
		fapp.Quit()
	}, win)
	if len(os.Args) >= 2 && os.Args[1] == "1" {
		cp.Advanced = true
		cp.SetColor(color.White)
	}
	if len(os.Args) >= 3 && os.Args[2] == "2" {
		cp.Refresh() // update the picker internal UI
	}
	cp.Resize(fyne.NewSize(500, 500)) // this might crash!!!

	cp.Show()
	win.Show()
	fapp.Run()
}

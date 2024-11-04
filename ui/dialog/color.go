package dialog

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"image/color"
	"log"
)

func runPickColor(colorDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	callback := confirmCallback(colorDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
	title, _ := colorDescr["title"].(string) // title is optional with zero value as default
	group, _ := colorDescr[ui.KeyGroup].(string)

	picker := dialog.NewColorPicker(title, "", func(c color.Color) {
		fmt.Println("Color callback called!")
		if c == nil {
			fmt.Println("c == nil")
			callback(false)
			return
		}
		fmt.Println("c != nil")
		hex := colorToString(c)
		ui.StoreValueByFullName(hex, fullName, group)
		callback(true)
	}, win)
	picker.Advanced, _ = colorDescr["advanced"].(bool)
	value := colorDescr["initialColor"]
	if value != nil {
		if c, ok := parseColor(value.(string), fullName); ok {
			picker.Advanced = true
			picker.SetColor(c)
		}
	}
	picker.Refresh() // update the picker internal UI

	value = colorDescr["buttonText"]
	if value != nil {
		picker.SetDismissText(value.(string))
	}

	width, height := run.GetSize(colorDescr)
	if width > 0 && height > 0 {
		picker.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		switch keyEvent.Name {
		case fyne.KeyEscape:
			callback(false)
		}
	})

	picker.Show()
}

func colorToString(c color.Color) string {
	nrgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	if nrgba.A == 0xff {
		return fmt.Sprintf("#%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
	}
	return fmt.Sprintf("#%02x%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B, nrgba.A)
}

func parseColor(s string, fullName string) (color.Color, bool) {
	var c color.NRGBA
	var err error
	if len(s) == 7 {
		c.A = 0xFF
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	} else {
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	}
	if err != nil {
		log.Printf(`ERROR: for %q: converting color: %v`, fullName, err)
		return color.NRGBA{}, false
	}
	return c, true
}

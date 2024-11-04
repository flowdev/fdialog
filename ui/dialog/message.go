package dialog

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
)

func runInfo(infoDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	title, _ := infoDescr["title"].(string)  // title is optional with zero value as default
	message := infoDescr["message"].(string) // message is required
	info := dialog.NewInformation(title, message, win)
	if children, ok := infoDescr[ui.KeyChildren]; ok {
		callback := closeCallback(children.(ui.CommandsDescr), fullName, win, uiDescr)
		info.SetOnClosed(callback)
	}

	value := infoDescr["buttonText"]
	if value != nil {
		info.SetDismissText(value.(string))
	}

	width, height := run.GetSize(infoDescr)
	if width > 0 && height > 0 {
		info.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		switch keyEvent.Name {
		case fyne.KeyReturn, fyne.KeyEnter, fyne.KeySpace, fyne.KeyEscape:
			win.Close()
		}
	})

	info.Show()
}

func runError(errorDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	message := errorDescr["message"].(string) // message is required
	errorDialog := dialog.NewError(errors.New(message), win)
	if children, ok := errorDescr[ui.KeyChildren]; ok {
		callback := closeCallback(children.(ui.CommandsDescr), fullName, win, uiDescr)
		errorDialog.SetOnClosed(callback)
	}

	value := errorDescr["buttonText"]
	if value != nil {
		errorDialog.SetDismissText(value.(string))
	}

	width, height := run.GetSize(errorDescr)
	if width > 0 && height > 0 {
		errorDialog.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		switch keyEvent.Name {
		case fyne.KeyReturn, fyne.KeyEnter, fyne.KeySpace, fyne.KeyEscape:
			win.Close()
		}
	})

	ui.StoreExitCode(0) // error has been noted; so all is OK
	errorDialog.Show()
}

func runConfirmation(cnfDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	callback := confirmCallback(cnfDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
	title, _ := cnfDescr["title"].(string)  // title is optional with zero value as default
	message := cnfDescr["message"].(string) // message is required
	cnf := dialog.NewConfirm(title, message, callback, win)

	value := cnfDescr["confirmText"]
	if value != nil {
		cnf.SetConfirmText(value.(string))
	}
	value = cnfDescr["dismissText"]
	if value != nil {
		cnf.SetDismissText(value.(string))
	}

	width, height := run.GetSize(cnfDescr)
	if width > 0 && height > 0 {
		cnf.Resize(fyne.NewSize(width, height))
	}
	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		switch keyEvent.Name {
		case fyne.KeyReturn, fyne.KeyEnter, fyne.KeySpace:
			callback(true)
		case fyne.KeyEscape:
			callback(false)
		}
	})

	cnf.Show()
}

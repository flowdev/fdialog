package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"strings"
)

func runOpenFile(ofDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	outputKey, _ := ofDescr[ui.AttrOutputKey].(string)
	id, _ := ofDescr[ui.AttrID].(string)
	group, _ := ofDescr[ui.AttrGroup].(string)
	callback := run.BooleanCallback(ofDescr[ui.AttrChildren].(ui.CommandsDescr),
		ui.NameChoose, ui.NameCancel, fullName, win, uiDescr)
	ofDialog := dialog.NewFileOpen(func(frd fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			callback(false)
			return
		}
		if frd == nil {
			callback(false)
			return
		}
		fileName := strings.TrimPrefix(frd.URI().String(), "file://")
		ui.StoreValue(fileName, outputKey, id, fullName, group)
		callback(true)
	}, win)

	extAttr := ofDescr["extensions"]
	if extAttr != nil {
		ofDialog.SetFilter(storage.NewExtensionFileFilter(ui.AnysToStrings(extAttr)))
	}

	value := ofDescr["confirmText"]
	if value != nil {
		ofDialog.SetConfirmText(value.(string))
	}
	value = ofDescr["dismissText"]
	if value != nil {
		ofDialog.SetDismissText(value.(string))
	}

	width, height := run.GetSize(ofDescr)
	if width > 0 && height > 0 {
		ofDialog.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			win.Close()
		}
	})

	ofDialog.Show()
}

func runSaveFile(sfDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	outputKey, _ := sfDescr[ui.AttrOutputKey].(string)
	id, _ := sfDescr[ui.AttrID].(string)
	group, _ := sfDescr[ui.AttrGroup].(string)
	callback := run.BooleanCallback(sfDescr[ui.AttrChildren].(ui.CommandsDescr),
		ui.NameChoose, ui.NameCancel, fullName, win, uiDescr)
	sfDialog := dialog.NewFileSave(func(fwr fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			callback(false)
			return
		}
		if fwr == nil {
			callback(false)
			return
		}

		fileName := strings.TrimPrefix(fwr.URI().String(), "file://")
		ui.StoreValue(fileName, outputKey, id, fullName, group)
		callback(true)
	}, win)

	extAttr := sfDescr["extensions"]
	if extAttr != nil {
		sfDialog.SetFilter(storage.NewExtensionFileFilter(ui.AnysToStrings(extAttr)))
	}

	value := sfDescr["chooseText"]
	if value != nil {
		sfDialog.SetConfirmText(value.(string))
	}
	value = sfDescr["cancelText"]
	if value != nil {
		sfDialog.SetDismissText(value.(string))
	}

	width, height := run.GetSize(sfDescr)
	if width > 0 && height > 0 {
		sfDialog.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			win.Close()
		}
	})

	sfDialog.Show()
}

func runOpenFolder(ofDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	outputKey, _ := ofDescr[ui.AttrOutputKey].(string)
	id, _ := ofDescr[ui.AttrID].(string)
	group, _ := ofDescr[ui.AttrGroup].(string) // group is optional with zero value as default
	callback := run.BooleanCallback(ofDescr[ui.AttrChildren].(ui.CommandsDescr),
		ui.NameChoose, ui.NameCancel, fullName, win, uiDescr)

	ofDialog := dialog.NewFolderOpen(func(fold fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if fold == nil {
			callback(false)
			return
		}
		folderName := strings.TrimPrefix(fold.String(), "file://")
		ui.StoreValue(folderName, outputKey, id, fullName, group)
		callback(true)
	}, win)

	value := ofDescr["chooseText"]
	if value != nil {
		ofDialog.SetConfirmText(value.(string))
	}
	value = ofDescr["cancelText"]
	if value != nil {
		ofDialog.SetDismissText(value.(string))
	}

	width, height := run.GetSize(ofDescr)
	if width > 0 && height > 0 {
		ofDialog.Resize(fyne.NewSize(width, height))
	}

	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			win.Close()
		}
	})

	ofDialog.Show()
}

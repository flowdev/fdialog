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
	outputKey, _ := ofDescr[ui.KeyOutputKey].(string)
	id, _ := ofDescr[ui.KeyID].(string)
	group, _ := ofDescr[ui.KeyGroup].(string)
	callback := confirmCallback(ofDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
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
		extSlice := strings.Split(extAttr.(string), ",")
		for i := 0; i < len(extSlice); i++ {
			extSlice[i] = strings.TrimSpace(extSlice[i])
		}
		ofDialog.SetFilter(storage.NewExtensionFileFilter(extSlice))
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
	outputKey, _ := sfDescr[ui.KeyOutputKey].(string)
	id, _ := sfDescr[ui.KeyID].(string)
	group, _ := sfDescr[ui.KeyGroup].(string)
	callback := confirmCallback(sfDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
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
		extSlice := strings.Split(extAttr.(string), ",")
		for i := 0; i < len(extSlice); i++ {
			extSlice[i] = strings.TrimSpace(extSlice[i])
		}
		sfDialog.SetFilter(storage.NewExtensionFileFilter(extSlice))
	}

	value := sfDescr["confirmText"]
	if value != nil {
		sfDialog.SetConfirmText(value.(string))
	}
	value = sfDescr["dismissText"]
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
	outputKey, _ := ofDescr[ui.KeyOutputKey].(string)
	id, _ := ofDescr[ui.KeyID].(string)
	group, _ := ofDescr[ui.KeyGroup].(string) // group is optional with zero value as default
	callback := confirmCallback(ofDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)

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

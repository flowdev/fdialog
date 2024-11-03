package dialog

import (
	"errors"
	"log"
	"math"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
)

const KeywordDialog = "dialog"

func RegisterAll() error {
	// -----------------------------------------------------------------------
	// Register Validators
	//

	err := ui.RegisterValidKeyword(KeywordDialog, "info", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("info"),
			},
			"title": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"message": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"buttonText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: false,
				Validate: valid.ChildrenValidator(0, 1),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordDialog, "error", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("error"),
			},
			"message": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"buttonText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: false,
				Validate: valid.ChildrenValidator(0, 1),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordDialog, "confirmation", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("confirmation"),
			},
			"title": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"message": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"dismissText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"confirmText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: true,
				Validate: valid.ChildrenValidator(2, 2),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordDialog, "openFile", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("openFile"),
			},
			"extensions": {
				Validate: valid.StringValidator(2, 0, nil),
			},
			"dismissText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"confirmText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: true,
				Validate: valid.ChildrenValidator(2, 2),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordDialog, "saveFile", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("saveFile"),
			},
			"extensions": {
				Validate: valid.StringValidator(2, 0, nil),
			},
			"dismissText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"confirmText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: true,
				Validate: valid.ChildrenValidator(2, 2),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordDialog, "openFolder", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("openFolder"),
			},
			"dismissText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"confirmText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Required: true,
				Validate: valid.ChildrenValidator(2, 2),
			},
		},
	})
	if err != nil {
		return err
	}

	// -----------------------------------------------------------------------
	// Register Runners
	//

	err = ui.RegisterRunKeyword(KeywordDialog, "dlg", runDialog)
	if err != nil {
		return err
	}
	return nil
}

func runDialog(dialogDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	dlg := dialogDescr[ui.KeyType]

	switch dlg {
	case "info":
		runInfo(dialogDescr, fullName, win, uiDescr)
	case "error":
		runError(dialogDescr, fullName, win, uiDescr)
	case "confirmation":
		runConfirmation(dialogDescr, fullName, win, uiDescr)
	case "openFile":
		runOpenFile(dialogDescr, fullName, win, uiDescr)
	case "saveFile":
		runSaveFile(dialogDescr, fullName, win, uiDescr)
	case "openFolder":
		runOpenFolder(dialogDescr, fullName, win, uiDescr)
	default:
		log.Printf(`ERROR: for %q: unknown dialog type %q`, fullName, dlg)
	}
}

func runInfo(infoDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	_ = fullName // currently not used but might change

	value := infoDescr["title"]
	title := ""
	if value != nil {
		title = value.(string)
	}
	message := infoDescr["message"].(string) // message is required
	info := dialog.NewInformation(title, message, win)
	if children, ok := infoDescr[ui.KeyChildren]; ok {
		callback := closeCallback(children.(ui.CommandsDescr), fullName, win, uiDescr)
		info.SetOnClosed(callback)
	}

	value = infoDescr["buttonText"]
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
	_ = fullName // currently not used but might change

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

func closeCallback(childDescr ui.CommandsDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) func() {
	defaultCallback := func() {
		return
	}
	actClose, _ := childDescr.Get("close")
	if actClose == nil { // action is optional
		return defaultCallback
	}
	keyword := actClose[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		log.Printf("ERROR: for %q: close action is not an action but a %q", fullName, keyword)
		return defaultCallback
	}

	return func() {
		run.Action(actClose, ui.FullNameFor(fullName, "close"), win, uiDescr)
	}
}

func runConfirmation(cnfDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	callback := confirmCallback(cnfDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
	title := ""
	value := cnfDescr["title"]
	if value != nil {
		title = value.(string)
	}
	message := cnfDescr["message"].(string) // message is required
	cnf := dialog.NewConfirm(title, message, callback, win)

	value = cnfDescr["confirmText"]
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

func confirmCallback(
	childrenDescr ui.CommandsDescr,
	fullName string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) func(bool) {
	defaultCallback := func(_ bool) {
		return
	}

	actConfirm, _ := childrenDescr.Get("confirm")
	if actConfirm == nil {
		log.Printf("ERROR: for %q: confirm action is missing", fullName)
		return defaultCallback
	}
	keyword := actConfirm[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		log.Printf("ERROR: for %q: confirm action is not an action but a %q", fullName, keyword)
		return defaultCallback
	}

	actDismiss, _ := childrenDescr.Get("dismiss")
	if actDismiss == nil {
		log.Printf("ERROR: for %q: dismiss action is missing", fullName)
		return defaultCallback
	}
	keyword = actDismiss[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		log.Printf("ERROR: for %q: dismiss action is not an action but a %q", fullName, keyword)
		return defaultCallback
	}

	return func(confirmed bool) {
		if confirmed {
			run.Action(actConfirm, ui.FullNameFor(fullName, "confirm"), win, uiDescr)
		} else {
			run.Action(actDismiss, ui.FullNameFor(fullName, "dismiss"), win, uiDescr)
		}
	}
}

func runOpenFile(ofDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	_, _ = fullName, uiDescr

	group, _ := ofDescr[ui.KeyGroup].(string)
	callback := confirmCallback(ofDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
	ofDialog := dialog.NewFileOpen(func(frd fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if frd == nil {
			callback(false)
			return
		}
		fileName := strings.TrimPrefix(frd.URI().String(), "file://")
		ui.StoreValueByFullName(fileName, fullName, group)
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
		ui.StoreValueByFullName(fileName, fullName, group)
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
	_, _ = fullName, uiDescr

	group, _ := ofDescr[ui.KeyGroup].(string)
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
		ui.StoreValueByFullName(folderName, fullName, group)
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

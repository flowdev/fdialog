package dialog

import (
	"errors"
	"fmt"
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
			//ui.KeyChildren: {
			//	Required: true,
			//	Validate: valid.ChildrenValidator(2, 2),
			//},
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

func runDialog(dialogDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	var err error
	dlg := dialogDescr[ui.KeyType]

	switch dlg {
	case "info":
		err = runInfo(dialogDescr, fullName, win, uiDescr)
	case "error":
		err = runError(dialogDescr, fullName, win, uiDescr)
	case "confirmation":
		err = runConfirmation(dialogDescr, fullName, win, uiDescr)
	case "openFile":
		err = runOpenFile(dialogDescr, fullName, win, uiDescr)
	default:
		err = fmt.Errorf(`for %q: unknown dialog type %q`, ui.DisplayName(fullName), dlg)
	}
	return err
}

func runInfo(infoDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	_ = fullName // currently not used but might change

	value := infoDescr["title"]
	title := ""
	if value != nil {
		title = value.(string)
	}
	message := infoDescr["message"].(string) // message is required
	info := dialog.NewInformation(title, message, win)
	if children, ok := infoDescr[ui.KeyChildren]; ok {
		callback, err := messageCallback(children.(ui.CommandsDescr), fullName, win, uiDescr)
		if err != nil {
			return err
		}
		if callback != nil {
			info.SetOnClosed(callback)
		}
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

	ui.StoreExitCode(0) // info has been noted; so all is OK
	info.Show()
	return nil
}

func runError(errorDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	_ = fullName // currently not used but might change

	message := errorDescr["message"].(string) // message is required
	errorDialog := dialog.NewError(errors.New(message), win)
	if children, ok := errorDescr[ui.KeyChildren]; ok {
		callback, err := messageCallback(children.(ui.CommandsDescr), fullName, win, uiDescr)
		if err != nil {
			return err
		}
		if callback != nil {
			errorDialog.SetOnClosed(callback)
		}
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
	return nil
}

func messageCallback(
	childrenDescr ui.CommandsDescr,
	fullName []string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) (func(), error) {

	actClose := childrenDescr["close"]
	if actClose == nil { // action is optional
		return nil, nil
	}
	keyword := actClose[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		return nil, fmt.Errorf("for %q: close action is not an action but a %q",
			ui.DisplayName(fullName), keyword)
	}

	return func() {
		err := run.Action(actClose, append(fullName, "close"), win, uiDescr)
		if err != nil {
			log.Printf("ERROR: Can't run close action: %v", err)
		}
	}, nil
}

func runConfirmation(
	cnfDescr ui.AttributesDescr,
	fullName []string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) error {

	callback, err := confirmCallback(cnfDescr[ui.KeyChildren].(ui.CommandsDescr), fullName, win, uiDescr)
	if err != nil {
		return err
	}

	value := cnfDescr["title"]
	title := ""
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

	ui.StoreExitCode(1) // closing the window => dismissed
	cnf.Show()
	return nil
}

func confirmCallback(
	childrenDescr ui.CommandsDescr,
	fullName []string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) (func(bool), error) {

	actConfirm := childrenDescr["confirm"]
	if actConfirm == nil {
		return nil, fmt.Errorf("for %q: confirm action is missing", ui.DisplayName(fullName))
	}
	keyword := actConfirm[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		return nil, fmt.Errorf("for %q: confirm action is not an action but a %q",
			ui.DisplayName(fullName), keyword)
	}

	actDismiss := childrenDescr["dismiss"]
	if actDismiss == nil {
		return nil, fmt.Errorf("for %q: dismiss action is missing", ui.DisplayName(fullName))
	}
	keyword = actDismiss[ui.KeyKeyword].(string)
	if keyword != ui.KeywordAction {
		return nil, fmt.Errorf("for %q: dismiss action is not an action but a %q",
			ui.DisplayName(fullName), keyword)
	}
	return func(confirmed bool) {
		if confirmed {
			err := run.Action(actConfirm, append(fullName, "confirm"), win, uiDescr)
			if err != nil {
				log.Printf("ERROR: Can't run confirm action: %v", err)
			}
		} else {
			err := run.Action(actDismiss, append(fullName, "dismiss"), win, uiDescr)
			if err != nil {
				log.Printf("ERROR: Can't run dismiss action: %v", err)
			}
		}
	}, nil
}

func runOpenFile(
	ofDescr ui.AttributesDescr,
	fullName []string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) error {
	_, _ = fullName, uiDescr
	ofDialog := dialog.NewFileOpen(func(frd fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if frd == nil {
			fmt.Println("<CLOSED>")
			ui.ExitApp(1)
		}
		fmt.Println("file to open:", strings.TrimPrefix(frd.URI().String(), "file://"))

		ui.ExitApp(0)
	}, win)

	ofDialog.SetOnClosed(func() {
		fmt.Println("dialog closed, exiting?")
	})

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

	ofDialog.Show()
	return nil
}

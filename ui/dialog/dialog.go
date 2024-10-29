package dialog

import (
	"errors"
	"fmt"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
	"log"
	"math"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

func runDialog(dialogDescr map[string]any, fullName []string, win fyne.Window, uiDescr map[string]map[string]any) error {
	var err error
	dlg := dialogDescr[ui.KeyType]

	switch dlg {
	case "info":
		err = runInfo(dialogDescr, fullName, win)
	case "error":
		err = runError(dialogDescr, fullName, win)
	case "confirmation":
		err = runConfirmation(dialogDescr, fullName, win, uiDescr)
	case "openFile":
		err = runOpenFile(dialogDescr, fullName, win, uiDescr)
	default:
		err = fmt.Errorf(`for %q: unknown dialog type %q`, ui.DisplayName(fullName), dlg)
	}
	return err
}

func runOpenFile(
	ofDescr map[string]any,
	fullName []string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
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

	width := float64(0)
	height := float64(0)
	if _, ok := ofDescr["width"]; ok {
		width = ofDescr["width"].(float64)
	}
	if _, ok := ofDescr["height"]; ok {
		height = ofDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		ofSize := fyne.NewSize(float32(width), float32(height))
		ofDialog.Resize(ofSize)
	}

	ofDialog.Show()
	return nil
}

func runConfirmation(
	cnfDescr map[string]any,
	fullName []string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
) error {

	callback, err := confirmCallback(cnfDescr[ui.KeyChildren].(map[string]map[string]any), fullName, win, uiDescr)
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

	width := float64(0)
	height := float64(0)
	if _, ok := cnfDescr["width"]; ok {
		width = cnfDescr["width"].(float64)
	}
	if _, ok := cnfDescr["height"]; ok {
		height = cnfDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		cnfSize := fyne.NewSize(float32(width), float32(height))
		cnf.Resize(cnfSize)
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
	childrenDescr map[string]map[string]any,
	fullName []string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
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

func runError(errorDescr map[string]any, fullName []string, win fyne.Window) error {
	_ = fullName                              // currently not used but might change
	message := errorDescr["message"].(string) // message is required
	errorDialog := dialog.NewError(errors.New(message), win)
	errorDialog.SetOnClosed(func() {
		os.Exit(0) // error has been noted
	})

	value := errorDescr["buttonText"]
	if value != nil {
		errorDialog.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := errorDescr["width"]; ok {
		width = errorDescr["width"].(float64)
	}
	if _, ok := errorDescr["height"]; ok {
		height = errorDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		infoSize := fyne.NewSize(float32(width), float32(height))
		errorDialog.Resize(infoSize)
	}

	ui.StoreExitCode(0) // error has been noted; so all is OK
	errorDialog.Show()
	return nil
}

func runInfo(infoDescr map[string]any, fullName []string, win fyne.Window) error {
	_ = fullName // currently not used but might change
	value := infoDescr["title"]
	title := ""
	if value != nil {
		title = value.(string)
	}
	message := infoDescr["message"].(string) // message is required
	info := dialog.NewInformation(title, message, win)
	info.SetOnClosed(func() {
		os.Exit(0) // info has been noted
	})

	value = infoDescr["buttonText"]
	if value != nil {
		info.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := infoDescr["width"]; ok {
		width = infoDescr["width"].(float64)
	}
	if _, ok := infoDescr["height"]; ok {
		height = infoDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		infoSize := fyne.NewSize(float32(width), float32(height))
		info.Resize(infoSize)
	}

	ui.StoreExitCode(0) // info has been noted; so all is OK
	info.Show()
	return nil
}

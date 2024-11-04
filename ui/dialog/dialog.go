package dialog

import (
	"fyne.io/fyne/v2"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
	"log"
	"math"
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

	err = ui.RegisterValidKeyword(KeywordDialog, "pickColor", ui.ValidAttributesType{
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
				Validate: valid.ExactStringValidator("pickColor"),
			},
			"title": {
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
			"advanced": {
				Validate: valid.BoolValidator(),
			},
			"initialColor": {
				Validate: valid.StringValidator(4, 9, ui.ColorRegex),
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
	case "pickColor":
		runPickColor(dialogDescr, fullName, win, uiDescr)
	default:
		log.Printf(`ERROR: for %q: unknown dialog type %q`, fullName, dlg)
	}
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

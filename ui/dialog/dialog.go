package dialog

import (
	"fyne.io/fyne/v2"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
	"log"
	"math"
	"regexp"
)

const KeywordDialog = "dialog"

var extensionRegex = regexp.MustCompile(`^\..+$`)

func RegisterAll() error {
	// -----------------------------------------------------------------------
	// Register Validators
	//

	err := ui.RegisterValidKeyword(KeywordDialog, "info", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
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
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
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
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
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
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("openFile"),
			},
			"extensions": {
				Validate: valid.ListValidator(1, math.MaxInt,
					valid.StringValidator(2, 0, extensionRegex)),
			},
			"cancelText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"chooseText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.AttrOutputKey: valid.ValidateOutputKey,
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("saveFile"),
			},
			"extensions": {
				Validate: valid.ListValidator(1, math.MaxInt,
					valid.StringValidator(2, 0, extensionRegex)),
			},
			"cancelText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"chooseText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.AttrOutputKey: valid.ValidateOutputKey,
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("openFolder"),
			},
			"cancelText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"chooseText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.AttrOutputKey: valid.ValidateOutputKey,
			ui.AttrChildren: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordDialog),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("pickColor"),
			},
			"title": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"cancelText": {
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
				Validate: valid.StringValidator(7, 9, ui.ColorRegex),
			},
			ui.AttrOutputKey: valid.ValidateOutputKey,
			ui.AttrChildren: {
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
	dlg := dialogDescr[ui.AttrType]

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

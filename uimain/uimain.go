package uimain

import (
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/ui/dialog"
	"github.com/flowdev/fdialog/valid"
	"math"
)

func RegisterEverything() error {
	err := RegisterBase()
	if err != nil {
		return err
	}
	err = dialog.RegisterAll()
	if err != nil {
		return err
	}
	return nil
}

func RegisterBase() error {
	// -----------------------------------------------------------------------
	// Register Validators
	//

	err := ui.RegisterValidKeyword(ui.KeywordWindow, "", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordWindow),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			"title": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"width": {
				Validate: valid.FloatValidator(50.0, math.MaxFloat32),
			},
			"height": {
				Validate: valid.FloatValidator(80.0, math.MaxFloat32),
			},
			ui.KeyChildren: {
				Validate: valid.ChildrenValidator(0, math.MaxInt),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(ui.KeywordLink, "", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordLink),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			"destination": {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.LinkRegex),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(ui.KeywordAction, "exit", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.KeyKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordAction),
			},
			ui.KeyName: {
				Required: true,
				Validate: valid.StringValidator(1, 0, ui.NameRegex),
			},
			ui.KeyType: {
				Required: true,
				Validate: valid.ExactStringValidator("exit"),
			},
			"code": {
				Validate: valid.IntValidator(0, 125),
			},
		},
	})
	if err != nil {
		return err
	}

	// -----------------------------------------------------------------------
	// Register Runners
	//

	// Keywords:
	err = ui.RegisterRunKeyword(ui.KeywordWindow, "win", run.Window)
	if err != nil {
		return err
	}
	err = ui.RegisterRunKeyword(ui.KeywordLink, "lnk", run.Link)
	if err != nil {
		return err
	}
	err = ui.RegisterRunKeyword(ui.KeywordAction, "act", run.Action)
	if err != nil {
		return err
	}

	// Actions:
	err = ui.RegisterAction("exit", run.Exit)
	if err != nil {
		return err
	}
	return nil
}

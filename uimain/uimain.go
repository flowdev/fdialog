package uimain

import (
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/ui/dialog"
	"github.com/flowdev/fdialog/ui/widget"
	"github.com/flowdev/fdialog/valid"
	"log"
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
	err = widget.RegisterAll()
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordWindow),
			},
			ui.AttrType: {
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
			"appId": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"exitCode": {
				Validate: valid.IntValidator(0, 125),
			},
			ui.AttrChildren: {
				Validate: valid.ChildrenValidator(0, math.MaxInt),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(ui.KeywordLink, "", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordLink),
			},
			ui.AttrType: {
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
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordAction),
			},
			ui.AttrType: {
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

	err = ui.RegisterValidKeyword(ui.KeywordAction, "close", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordAction),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("close"),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(ui.KeywordAction, "group", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordAction),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("group"),
			},
			ui.AttrChildren: {
				Validate: valid.ChildrenValidator(1, math.MaxInt),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(ui.KeywordAction, "write", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(ui.KeywordAction),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("write"),
			},
			"fullName": {
				Validate: valid.StringValidator(1, 0, ui.LinkRegex),
			},
			ui.AttrOutputKey: valid.ValidateOutputKey,
		},
		Validate: func(attrs ui.AttributesDescr, parent string) bool {
			_, okGroup := attrs[ui.AttrGroup].(string)
			_, okID := attrs[ui.AttrID].(string)
			_, okName := attrs["fullName"].(string)
			_, okOutKey := attrs[ui.AttrOutputKey].(string)
			ok := okGroup || (okID && okOutKey) || (okName && okOutKey)
			if !ok {
				log.Printf(`ERROR: for %q: attribute "group" or attributes: `+
					`"outputKey" and one "id" or "fullName" are required`,
					parent)
			}
			return ok
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
	err = ui.RegisterAction("close", run.Close)
	if err != nil {
		return err
	}
	err = ui.RegisterAction("group", run.Group)
	if err != nil {
		return err
	}
	err = ui.RegisterAction("write", run.Write)
	if err != nil {
		return err
	}
	return nil
}

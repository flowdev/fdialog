package uimain

import (
	"github.com/flowdev/fdialog/dialog"
	"github.com/flowdev/fdialog/run"
)

func RegisterEverything() error {
	err := run.RegisterAll()
	if err != nil {
		return err
	}
	err = dialog.RegisterAll()
	if err != nil {
		return err
	}
	return nil
}

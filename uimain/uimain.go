package uimain

import (
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui/dialog"
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

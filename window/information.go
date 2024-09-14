// This package has been copied from fyne.io/fyne/v2/dialog
// So we have to update and maintain this forever. :(
package window

import (
	"unicode"
	"unicode/utf8"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/lang"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createInformationDialog(title, message string, icon fyne.Resource, parent fyne.App) Dialog {
	d := newTextDialog(title, message, icon, parent)
	d.dismiss = &widget.Button{
		Text:     lang.L("OK"),
		OnTapped: d.Hide,
	}
	d.create(container.NewGridWithColumns(1, d.dismiss))
	return d
}

// NewInformation creates a dialog window for the specified app for user information.
// The title is used for the dialog window and message is the content.
// After creation you should call Show().
func NewInformation(title, message string, parent fyne.App) Dialog {
	return createInformationDialog(title, message, theme.InfoIcon(), parent)
}

// NewError creates a dialog window for the specified app for an application error.
// The message is extracted from the provided error (should not be nil).
// After creation you should call Show().
func NewError(err error, parent fyne.App) Dialog {
	dialogText := err.Error()
	r, size := utf8.DecodeRuneInString(dialogText)
	if r != utf8.RuneError {
		dialogText = string(unicode.ToUpper(r)) + dialogText[size:]
	}
	return createInformationDialog(lang.L("Error"), dialogText, theme.ErrorIcon(), parent)
}

package parse

import (
	"fmt"
	"io"

	"github.com/flowdev/fdialog/ui"
)

// UIDescription parses a UI description in the given format and gives the content back suitable for validation.
// Supported formats are: (H)JSON and UIDL
// An error is returned if the steam can't be read or unmarshalled or a data type doesn't match.
func UIDescription(input io.Reader, name string, format string) (ui.CommandsDescr, error) {
	var uiDescr ui.CommandsDescr
	var err error

	switch format {
	case "json":
		uiDescr, err = JSON(input, name)
	case "uidl":
		uiDescr, err = UIDL(input, name)
	default:
		err = fmt.Errorf("unknown UI description format: %s", format)
	}
	if err != nil {
		return nil, err
	}
	return uiDescr, err
}

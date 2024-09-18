package parse

import (
	"fmt"
	"io"
)

// UIDescription parses a UI description in the given format and gives the content back suitable for validation.
// Supported formats are: (H)JSON and UIDL
// An error is returned if the steam can't be read or unmarshalled or a data type doesn't match.
func UIDescription(input io.Reader, name string, format string, strict bool) (map[string]map[string]any, error) {
	var uiDescr map[string]map[string]any
	var err error

	switch format {
	case "json":
		uiDescr, err = ParseJSON(input, name)
	case "uidl":
		uiDescr, err = ParseUIDL(input, name)
	default:
		err = fmt.Errorf("unknown UI description format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	err = Validate(uiDescr, strict)
	if err != nil {
		return nil, fmt.Errorf("error validating uiDescr from file %q in format %q: %w", name, format, err)
	}
	return uiDescr, err
}

package parse

import (
	"fmt"
	"io"
)

// UIDescription parses a UI description in the given format and gives the content back suitable for validation.
// Supported formats are: (H)JSON and UIDL
// An error is returned if the steam can't be read or unmarshalled or a data type doesn't match.
func UIDescription(input io.Reader, format string) (map[string]map[string]any, error) {
	switch format {
	case "json":
		return ParseJSON(input)
	default:
		return nil, fmt.Errorf("unknown format: %s", format)
	}
}

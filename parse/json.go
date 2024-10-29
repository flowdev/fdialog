package parse

import (
	"fmt"
	"github.com/flowdev/fdialog/ui"
	hjson "github.com/hjson/hjson-go/v4"
	"io"
)

// ParseJSON parses (H)JSON from a Reader and gives the content back suitable for validation.
// An error is returned if the stream can't be unmarshalled or a data type doesn't match.
func ParseJSON(input io.Reader, name string) (map[string]map[string]any, error) {
	data := make(map[string]map[string]any)
	inputData, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	// Decode with default options and check for errors.
	err = hjson.UnmarshalWithOptions(inputData, data, hjson.DecoderOptions{
		DisallowUnknownFields: true,
		DisallowDuplicateKeys: true,
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON file %v: %w", name, err)
	}

	err = cleanAllChildren(data)
	if err != nil {
		return nil, fmt.Errorf("error cleaning JSON data from file %v: %w", name, err)
	}

	return data, nil
}

func cleanAllChildren(data map[string]map[string]any) error {
	for keyword, value := range data {
		if err := cleanKeyword(value, []string{keyword}); err != nil {
			return err
		}
	}
	return nil
}

func cleanKeyword(data map[string]any, fullName []string) error {
	if rawChildren, ok := data[ui.KeyChildren]; ok {
		mapChildren, ok := rawChildren.(map[string]any)
		if !ok {
			return fmt.Errorf("for %q: expected attribute children to contain a map[string]any, got %T",
				ui.DisplayName(fullName), rawChildren)
		}

		cleanChildren := make(map[string]map[string]any, len(mapChildren))
		for key, avalue := range mapChildren {
			cleanChild, ok := avalue.(map[string]any)
			if !ok {
				return fmt.Errorf("for %q: expected value of keyword to be a map[string]any, got %T",
					ui.DisplayName(append(fullName, key)), avalue)
			}
			cleanChildren[key] = cleanChild
		}
		if err := cleanAllChildren(cleanChildren); err != nil {
			return err
		}
		data[ui.KeyChildren] = cleanChildren
	}
	return nil
}

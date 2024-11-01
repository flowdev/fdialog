package parse

import (
	"fmt"
	"github.com/flowdev/fdialog/ui"
	hjson "github.com/hjson/hjson-go/v4"
	"io"
)

// JSON parses (H)JSON from a Reader and gives the content back suitable for validation.
// An error is returned if the stream can't be unmarshalled or a data type doesn't match.
func JSON(input io.Reader, name string) (ui.CommandsDescr, error) {
	data := make(ui.CommandsDescr)
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

	err = cleanAllChildren(data, "")
	if err != nil {
		return nil, fmt.Errorf("error cleaning JSON data from file %v: %w", name, err)
	}

	return data, nil
}

func cleanAllChildren(data ui.CommandsDescr, parent string) error {
	for name, value := range data {
		if err := cleanCommand(value, ui.FullNameFor(parent, name)); err != nil {
			return err
		}
	}
	return nil
}

func cleanCommand(data ui.AttributesDescr, fullName string) error {
	if rawChildren, ok := data[ui.KeyChildren]; ok {
		mapChildren, ok := rawChildren.(ui.AttributesDescr)
		if !ok {
			return fmt.Errorf("for %q: expected attribute children to contain a map[string]any, got %T",
				fullName, rawChildren)
		}

		cleanChildren := make(ui.CommandsDescr, len(mapChildren))
		for key, avalue := range mapChildren {
			cleanChild, ok := avalue.(ui.AttributesDescr)
			if !ok {
				return fmt.Errorf("for %q: expected value of keyword to be a map[string]any, got %T",
					ui.FullNameFor(fullName, key), avalue)
			}
			cleanChildren[key] = cleanChild
		}
		if err := cleanAllChildren(cleanChildren, fullName); err != nil {
			return err
		}
		data[ui.KeyChildren] = cleanChildren
	}
	return nil
}

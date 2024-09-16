package parse

import (
	"fmt"
	"io"

	hjson "github.com/hjson/hjson-go/v4"
)

// ParseJSON parses (H)JSON from a Reader and gives it back suitable for validation.
// An error is returned if the steam can't be unmarshalled or a data type doesn't match.
func ParseJSON(input io.Reader) (map[string]map[string]any, error) {
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
		return nil, err
	}

	err = cleanAllChildren(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func cleanAllChildren(data map[string]map[string]any) error {
	for _, value := range data {
		if err := cleanKeyword(value); err != nil {
			return err
		}
	}
	return nil
}

func cleanKeyword(data map[string]any) error {
	if rawChildren, ok := data[KeyChildren]; ok {
		mapChildren, ok := rawChildren.(map[string]any)
		if !ok {
			return fmt.Errorf("expected attribute cleanChildren to contain a map[string]any, got %T", rawChildren)
		}

		cleanChildren := make(map[string]map[string]any, len(mapChildren))
		for key, avalue := range mapChildren {
			cleanChild, ok := avalue.(map[string]any)
			if !ok {
				return fmt.Errorf("expected value of a keyword to be a map[string]any, got %T", avalue)
			}
			cleanChildren[key] = cleanChild
		}
		if err := cleanAllChildren(cleanChildren); err != nil {
			return err
		}
		data[KeyChildren] = cleanChildren
	}
	return nil
}

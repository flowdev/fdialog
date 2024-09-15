package parse

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
)

const (
	keyKeyword = "keyword"
	keyName    = "name"
	keyType    = "type"
)

type keywordType struct {
	keyword string
	typ     string
}

type keywordValueType struct {
	attributes map[string]attributeValueType
	validate   func(v any) error
}

type attributeValueType struct {
	kind     reflect.Kind
	required bool
	validate func(v any) error
}

var validKeywords = map[keywordType]keywordValueType{
	keywordType{"window", "info"}: {
		attributes: map[string]attributeValueType{
			keyKeyword: {
				validate: exactStringValidator("window", "info", keyKeyword, "window"),
			},
			"name": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			keyType: {
				validate: exactStringValidator("window", "info", keyType, "info"),
			},
			"title": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"message": {
				kind:     reflect.String,
				required: true,
				validate: stringValidator(1, 0, nil),
			},
			"buttonText": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"width": {
				kind:     reflect.Int,
				validate: intValidator(50, math.MaxInt64),
			},
			"height": {
				kind:     reflect.Int,
				validate: intValidator(80, math.MaxInt64),
			},
		},
	},
	keywordType{"window", "error"}: {
		attributes: map[string]attributeValueType{
			keyKeyword: {
				validate: exactStringValidator("window", "error", keyKeyword, "window"),
			},
			"name": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			keyType: {
				validate: exactStringValidator("window", "error", keyType, "error"),
			},
			"message": {
				kind:     reflect.String,
				required: true,
				validate: stringValidator(1, 0, nil),
			},
			"buttonText": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"width": {
				kind:     reflect.Int,
				validate: intValidator(50, math.MaxInt64),
			},
			"height": {
				kind:     reflect.Int,
				validate: intValidator(80, math.MaxInt64),
			},
		},
	},
	keywordType{"window", "confirmation"}: {
		attributes: map[string]attributeValueType{
			keyKeyword: {
				validate: exactStringValidator("window", "confirmation", keyKeyword, "window"),
			},
			"name": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			keyType: {
				validate: exactStringValidator("window", "confirmation", keyType, "confirmation"),
			},
			"title": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"message": {
				kind:     reflect.String,
				required: true,
				validate: stringValidator(1, 0, nil),
			},
			"dismissText": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"confirmText": {
				kind:     reflect.String,
				validate: stringValidator(1, 0, nil),
			},
			"width": {
				kind:     reflect.Int,
				validate: intValidator(50, math.MaxInt64),
			},
			"height": {
				kind:     reflect.Int,
				validate: intValidator(80, math.MaxInt64),
			},
		},
	},
}

// Validate validates the data from a whole UI description file independent of its format.
// If strict is true additional attributes are errors.
func Validate(uiDescr []map[string]any, strict bool) error {
	var errs []error

	for i, keywordMap := range uiDescr {
		keyword, name, typ, err := keywordNameType(i, keywordMap)
		if err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, validateKeyword(i, keyword, name, typ, keywordMap, strict))
		}
	}
	return errors.Join(errs...)
}

func validateKeyword(
	i int,
	keyword, name, typ string,
	valueMap map[string]any,
	strict bool,
) error {
	keywordTypeValidationData, ok := validKeywords[keywordType{keyword: keyword, typ: typ}]
	if !ok {
		return fmt.Errorf("for %d-th keyword map is the combination of keyword %q and type %q not supported",
			i, keyword, typ)
	}

	return validateAttributes(i, keyword, name, typ, valueMap, keywordTypeValidationData.attributes, strict)
}

func validateAttributes(
	i int,
	keyword, name, typ string,
	valueMap map[string]any,
	attributes map[string]attributeValueType,
	strict bool,
) error {

	validatedAttributes := make(map[string]bool, len(attributes))
	errs := make([]error, len(attributes))

	for attrName, attribute := range attributes {
		if vv, ok := valueMap[attrName]; ok {
			validatedAttributes[attrName] = true
			err := attribute.validate(vv)
			if err != nil {
				errs = append(errs, fmt.Errorf("for the %d-th keyword %q, name %q, type %q and attribute %q: %v",
					i, keyword, name, typ, attrName, err.Error()))
			}
		} else if attribute.required {
			errs = append(errs, fmt.Errorf("for the %d-th keyword %q, name %q and type %q is attribute %q required",
				i, keyword, name, typ, attrName))
		}
	}

	if len(validatedAttributes) != len(valueMap) {
		keysTooMuch := make([]string, len(valueMap)-len(validatedAttributes))

		for k := range valueMap {
			_, ok := validatedAttributes[k]
			if !ok {
				keysTooMuch = append(keysTooMuch, k)
			}
		}

		err := fmt.Errorf("for the %d-th keyword %q, name %q and type %q are these given attributes too much: %q",
			i, keyword, name, typ, keysTooMuch)
		if strict {
			errs = append(errs, err)
		} else {
			_, err = fmt.Fprintln(os.Stderr, "Warning:", err.Error())
			if err != nil {
				// can't do much here
			}
		}
	}
	return errors.Join(errs...)
}

func keywordNameType(i int, keywordMap map[string]any) (keyword, name, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[keyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "", "", fmt.Errorf(
			"expecting a string value for the %d-th keyword map and key %q, got a %s",
			i, keyKeyword, rkeyword.Kind(),
		)
	}
	keyword = rkeyword.String()

	rname := reflect.ValueOf(keywordMap[keyName])
	if rname.Kind() != reflect.String {
		return "", "", "", fmt.Errorf(
			"expecting a string value for the %d-th keyword map, keyword %q and key %q, got a %s",
			i, keyword, keyName, rname.Kind(),
		)
	}
	name = rname.String()

	rtype := reflect.ValueOf(keywordMap[keyType])
	if rtype.Kind() != reflect.String {
		return "", "", "", fmt.Errorf(
			"expecting a string value for the %d-th keyword map, keyword %q, name %q and key %q, got a %s",
			i, keyword, name, keyType, rtype.Kind(),
		)
	}
	typ = rtype.String()

	return keyword, name, typ, nil
}

func stringValidator(minLen, maxLen int, regex *regexp.Regexp) func(v any) error {
	return func(v any) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			return fmt.Errorf("expecting a string value, got %s", rv.Kind())
		}
		s := rv.String()

		if minLen > 0 && len(s) < minLen {
			return fmt.Errorf("string too short (min %d > actual %d)", minLen, len(s))
		}
		if maxLen > 0 && len(s) > maxLen {
			return fmt.Errorf("string too long (max %d < actual %d)", maxLen, len(s))
		}

		if regex != nil && !regex.MatchString(s) {
			return fmt.Errorf("string %q does not match pattern %q", s, regex.String())
		}

		return nil
	}
}
func exactStringValidator(keyword, typ, attribute string, expected string) func(v any) error {
	return func(v any) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			return fmt.Errorf("for keyword %q, type %q expecting attribute %s to be of type string, got %q",
				keyword, typ, attribute, rv.Kind())
		}
		s := rv.String()

		if s != expected {
			return fmt.Errorf("for keyword %q, type %q expecting attribute %s to have value %q, got %q",
				keyword, typ, attribute, expected, s)
		}

		return nil
	}
}

func intValidator(minVal, maxVal int64) func(v any) error {
	return func(v any) error {
		var i int64
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Float64 {
			f := rv.Float()
			i = int64(f)
			if f != float64(i) {
				return fmt.Errorf("expecting an int64 (or a float64 convertable to it), got %f", f)
			}
		} else if rv.Kind() != reflect.Int64 {
			return fmt.Errorf("expecting a int64 value, got %s", rv.Kind())
		} else {
			i = rv.Int()
		}

		if i < minVal {
			return fmt.Errorf("integer value too small (min %d > actual %d)", minVal, i)
		}
		if i > maxVal {
			return fmt.Errorf("integer value too big (max %d < actual %d)", maxVal, i)
		}
		return nil
	}
}

func floatValidator(minVal, maxVal float64) func(v any) error {
	return func(v any) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Float64 {
			return fmt.Errorf("expecting a float64 value, got %s", rv.Kind())
		}
		f := rv.Float()

		if !math.IsNaN(f) && f < minVal {
			return fmt.Errorf("float value too small (min %f > actual %f)", minVal, f)
		}
		if !math.IsNaN(f) && f > maxVal {
			return fmt.Errorf("float value too big (max %f < actual %f)", maxVal, f)
		}
		return nil
	}
}

func boolValidator() func(v any) error {
	return func(v any) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Bool {
			return fmt.Errorf("expecting a boolean value, got %s", rv.Kind())
		}
		return nil
	}
}

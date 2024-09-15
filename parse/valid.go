package parse

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"strings"
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
	required bool
	validate func(v any, strict bool, parent string) error
}

var validKeywords map[keywordType]keywordValueType

// need this init function, or we would get an initialization cycle :(
func init() {
	validKeywords = map[keywordType]keywordValueType{
		keywordType{"window", ""}: {
			attributes: map[string]attributeValueType{
				keyKeyword: {
					required: true,
					validate: exactStringValidator("window", "", keyKeyword, "window"),
				},
				keyName: {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				keyType: {
					validate: stringValidator(1, 0, nil),
				},
				"title": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: intValidator(50, math.MaxInt64),
				},
				"height": {
					validate: intValidator(80, math.MaxInt64),
				},
				"children": {
					validate: childrenValidator(0, math.MaxInt),
				},
			},
		},
		keywordType{"dialog", "info"}: {
			attributes: map[string]attributeValueType{
				keyKeyword: {
					required: true,
					validate: exactStringValidator("dialog", "info", keyKeyword, "dialog"),
				},
				keyName: {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				keyType: {
					required: true,
					validate: exactStringValidator("dialog", "info", keyType, "info"),
				},
				"title": {
					validate: stringValidator(1, 0, nil),
				},
				"message": {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				"buttonText": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: intValidator(50, math.MaxInt64),
				},
				"height": {
					validate: intValidator(80, math.MaxInt64),
				},
			},
		},
		keywordType{"dialog", "error"}: {
			attributes: map[string]attributeValueType{
				keyKeyword: {
					required: true,
					validate: exactStringValidator("dialog", "error", keyKeyword, "dialog"),
				},
				keyName: {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				keyType: {
					required: true,
					validate: exactStringValidator("dialog", "error", keyType, "error"),
				},
				"message": {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				"buttonText": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: intValidator(50, math.MaxInt64),
				},
				"height": {
					validate: intValidator(80, math.MaxInt64),
				},
			},
		},
		keywordType{"dialog", "confirmation"}: {
			attributes: map[string]attributeValueType{
				keyKeyword: {
					required: true,
					validate: exactStringValidator("dialog", "confirmation", keyKeyword, "dialog"),
				},
				keyName: {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				keyType: {
					required: true,
					validate: exactStringValidator("dialog", "confirmation", keyType, "confirmation"),
				},
				"title": {
					validate: stringValidator(1, 0, nil),
				},
				"message": {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				"dismissText": {
					validate: stringValidator(1, 0, nil),
				},
				"confirmText": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: intValidator(50, math.MaxInt64),
				},
				"height": {
					validate: intValidator(80, math.MaxInt64),
				},
				"children": {
					validate: childrenValidator(2, 2),
				},
			},
		},
		keywordType{"action", "exit"}: {
			attributes: map[string]attributeValueType{
				keyKeyword: {
					required: true,
					validate: exactStringValidator("action", "exit", keyKeyword, "action"),
				},
				keyName: {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				keyType: {
					required: true,
					validate: exactStringValidator("action", "exit", keyType, "exit"),
				},
				"code": {
					validate: intValidator(0, 127),
				},
			},
		},
	}
}

// Validate validates the data from a whole UI description file independent of its format.
// If strict is true additional attributes are errors.
// The keys of the first level map are the names of the windows, containers, ...
// The keys of the second level map are the attributes of the keyword map.
// Mandatory key for keyword maps is: "keyword"
// The key "type" is expected for most keywords but not for all.
func Validate(uiDescr map[string]map[string]any, strict bool) error {
	return validateRecursivMap(uiDescr, 0, 0, strict, "")
}

func validateRecursivMap(m map[string]map[string]any, minLen, maxLen int, strict bool, parent string) error {
	if minLen > 0 && len(m) < minLen {
		return fmt.Errorf("for keyword map %q: expecting at least %d map elements, got %d",
			parent, minLen, len(m))
	}
	if maxLen > 0 && len(m) > maxLen {
		return fmt.Errorf("for keyword map %q: expecting at most %d map elements, got %d",
			parent, maxLen, len(m))
	}

	var errs []error

	for name, keywordMap := range m {
		keywordMap[keyName] = name
		keyword, typ, err := getKeywordType(keywordMap, joinParentName(parent, name))
		if err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, validateKeyword(keyword, joinParentName(parent, name), typ, keywordMap, strict))
		}
	}
	return errors.Join(errs...)
}

func validateKeyword(
	keyword, name, typ string,
	valueMap map[string]any,
	strict bool,
) error {
	keywordTypeValidationData, ok := validKeywords[keywordType{keyword: keyword, typ: typ}]
	if !ok {
		return fmt.Errorf("for the keyword map %q is the combination of keyword %q and type %q not supported",
			name, keyword, typ)
	}

	return validateAttributes(keyword, name, typ, valueMap, keywordTypeValidationData.attributes, strict)
}

func validateAttributes(
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
			err := attribute.validate(vv, strict, name)
			if err != nil {
				errs = append(errs, fmt.Errorf("for the keyword %q, name %q, type %q and attribute %q: %v",
					keyword, name, typ, attrName, err.Error()))
			}
		} else if attribute.required {
			errs = append(errs, fmt.Errorf("for the keyword %q, name %q and type %q is attribute %q required",
				keyword, name, typ, attrName))
		}
	}

	if len(validatedAttributes) != len(valueMap) {
		keysTooMuch := make([]string, 0, len(valueMap)-len(validatedAttributes))

		for k := range valueMap {
			_, ok := validatedAttributes[k]
			if !ok {
				keysTooMuch = append(keysTooMuch, k)
			}
		}

		err := fmt.Errorf("for the keyword %q, name %q and type %q are these given attributes too much: %q",
			keyword, name, typ, keysTooMuch)
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

func getKeywordType(keywordMap map[string]any, name string) (keyword, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[keyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "", fmt.Errorf(
			"expecting a string value for the keyword map %q and key %q, got a %s",
			name, keyKeyword, rkeyword.Kind(),
		)
	}
	keyword = rkeyword.String()

	typ = "" // this it the intentional default
	atype, ok := keywordMap[keyType]
	if ok {
		rtype := reflect.ValueOf(atype)
		if rtype.Kind() != reflect.String {
			return "", "", fmt.Errorf(
				"expecting a string value for the keyword %q, name %q and key %q, got a %s",
				keyword, name, keyType, rtype.Kind(),
			)
		}
		typ = rtype.String()
	}

	return keyword, typ, nil
}

func stringValidator(minLen, maxLen int, regex *regexp.Regexp) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
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
func exactStringValidator(keyword, typ, attribute string, expected string) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
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

func intValidator(minVal, maxVal int64) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
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

func floatValidator(minVal, maxVal float64) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
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

func boolValidator() func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Bool {
			return fmt.Errorf("expecting a boolean value, got %s", rv.Kind())
		}
		return nil
	}
}

func childrenValidator(minLen, maxLen int) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Map {
			return fmt.Errorf("expecting a map value, got %s", rv.Kind())
		}

		m, ok := v.(map[string]map[string]any)
		if !ok {
			return fmt.Errorf("expecting a map[string]map[string]any value, got %T", v)
		}

		if len(m) < minLen {
			return fmt.Errorf("expecting at least %d map elements, got %d", minLen, len(m))
		}
		if len(m) > maxLen {
			return fmt.Errorf("expecting at most %d map elements, got %d", maxLen, len(m))
		}

		return validateRecursivMap(m, minLen, maxLen, strict, parent)
	}
}

func joinParentName(parent, name string) string {
	if parent == "" {
		return name
	}

	b := strings.Builder{}
	b.WriteString(parent)
	b.WriteByte('.')
	b.WriteString(name)
	return b.String()
}

package parse

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"reflect"
	"regexp"
	"strings"
)

// Reserved attribute names:
const (
	KeyKeyword  = "keyword"
	KeyName     = "name"
	KeyChildren = "children"
	KeyType     = "type" // type is used like an ordinary attribute, but it has special semantics
)

// Recognized keywords:
const (
	KeywordWindow = "window"
	KeywordDialog = "dialog"
	KeywordAction = "action"
	KeywordLink   = "link"
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
	nameRegex := regexp.MustCompile(`^[\pL\pN_]+$`)
	linkRegex := regexp.MustCompile(`^[\pL\pN_]+(?:[.][\pL\pN_]+)*$`)

	validKeywords = map[keywordType]keywordValueType{
		keywordType{KeywordWindow, ""}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordWindow, "", KeyKeyword, KeywordWindow),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					validate: stringValidator(1, 0, nameRegex),
				},
				"title": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: floatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: floatValidator(80.0, math.MaxFloat32),
				},
				KeyChildren: {
					validate: childrenValidator(0, math.MaxInt),
				},
			},
		},
		keywordType{KeywordLink, ""}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordLink, "", KeyKeyword, KeywordLink),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					validate: stringValidator(1, 0, nameRegex),
				},
				"destination": {
					required: true,
					validate: stringValidator(1, 0, linkRegex),
				},
			},
		},
		keywordType{KeywordDialog, "info"}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "info", KeyKeyword, KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "info", KeyType, "info"),
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
					validate: floatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: floatValidator(80.0, math.MaxFloat32),
				},
			},
		},
		keywordType{KeywordDialog, "error"}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "error", KeyKeyword, KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "error", KeyType, "error"),
				},
				"message": {
					required: true,
					validate: stringValidator(1, 0, nil),
				},
				"buttonText": {
					validate: stringValidator(1, 0, nil),
				},
				"width": {
					validate: floatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: floatValidator(80.0, math.MaxFloat32),
				},
			},
		},
		keywordType{KeywordDialog, "confirmation"}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "confirmation", KeyKeyword, KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator(KeywordDialog, "confirmation", KeyType, "confirmation"),
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
					validate: floatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: floatValidator(80.0, math.MaxFloat32),
				},
				KeyChildren: {
					required: true,
					validate: childrenValidator(2, 2),
				},
			},
		},
		keywordType{KeywordAction, "exit"}: {
			attributes: map[string]attributeValueType{
				KeyKeyword: {
					required: true,
					validate: exactStringValidator(KeywordAction, "exit", KeyKeyword, KeywordAction),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator(KeywordAction, "exit", KeyType, "exit"),
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
		keywordMap[KeyName] = name
		keyword, typ, err := getKeywordType(keywordMap, JoinParentName(parent, name))
		if err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, validateKeyword(keyword, JoinParentName(parent, name), typ, keywordMap, strict))
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
	if !ok && typ != "" { // try empty type; will error later if not supported
		keywordTypeValidationData, ok = validKeywords[keywordType{keyword: keyword, typ: ""}]
	}
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

		err := fmt.Errorf("for the keyword %q, name %q and type %q are these attributes too much: %q",
			keyword, name, typ, keysTooMuch)
		if strict {
			errs = append(errs, err)
		} else {
			slog.Warn(err.Error())
		}
	}
	return errors.Join(errs...)
}

func getKeywordType(keywordMap map[string]any, name string) (keyword, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[KeyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "", fmt.Errorf(
			"expecting a string value for the keyword map %q and key %q, got a %s",
			name, KeyKeyword, rkeyword.Kind(),
		)
	}
	keyword = rkeyword.String()

	typ = "" // this it the intentional default
	atype, ok := keywordMap[KeyType]
	if ok {
		rtype := reflect.ValueOf(atype)
		if rtype.Kind() != reflect.String {
			return "", "", fmt.Errorf(
				"expecting a string value for the keyword %q, name %q and key %q, got a %s",
				keyword, name, KeyType, rtype.Kind(),
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

func JoinParentName(parent, name string) string {
	if parent == "" {
		return name
	}

	b := strings.Builder{}
	b.WriteString(parent)
	b.WriteByte('.')
	b.WriteString(name)
	return b.String()
}

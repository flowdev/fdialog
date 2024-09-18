package parse

import (
	"errors"
	"fmt"
	"log"
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
					validate: exactStringValidator(KeywordWindow),
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
					validate: exactStringValidator(KeywordLink),
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
					validate: exactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator("info"),
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
					validate: exactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator("error"),
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
					validate: exactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator("confirmation"),
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
					validate: exactStringValidator(KeywordAction),
				},
				KeyName: {
					required: true,
					validate: stringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: exactStringValidator("exit"),
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
	return validateRecursiveMap(uiDescr, strict, "")
}

func validateRecursiveMap(m map[string]map[string]any, strict bool, parent string) error {
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
	keyword, fullName, typ string,
	valueMap map[string]any,
	strict bool,
) error {
	keywordTypeValidationData, ok := validKeywords[keywordType{keyword: keyword, typ: typ}]
	if !ok && typ != "" { // try empty type; will error later if not supported
		keywordTypeValidationData, ok = validKeywords[keywordType{keyword: keyword, typ: ""}]
	}
	if !ok {
		return fmt.Errorf("for %q: the combination of keyword %q and type %q is not supported",
			fullName, keyword, typ)
	}

	return validateAttributes(fullName, valueMap, keywordTypeValidationData.attributes, strict)
}

func validateAttributes(
	fullName string,
	valueMap map[string]any,
	attributes map[string]attributeValueType,
	strict bool,
) error {

	validatedAttributes := make(map[string]bool, len(attributes))
	errs := make([]error, len(attributes))

	for attrName, attribute := range attributes {
		if vv, ok := valueMap[attrName]; ok {
			validatedAttributes[attrName] = true
			err := attribute.validate(vv, strict, fullName)
			if err != nil {
				errs = append(errs, fmt.Errorf("for %q, attribute %q: %v",
					fullName, attrName, err.Error()))
			}
		} else if attribute.required {
			errs = append(errs, fmt.Errorf("for %q, is attribute %q required",
				fullName, attrName))
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

		err := fmt.Errorf("for %q: these attributes are too much: %s", fullName, keysTooMuch)
		if strict {
			errs = append(errs, err)
		} else {
			log.Printf("WARNING: %v", err)
		}
	}
	return errors.Join(errs...)
}

func getKeywordType(keywordMap map[string]any, fullName string) (keyword, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[KeyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "",
			fmt.Errorf("for %q: expecting the keyword to be a string, got a %s",
				fullName, rkeyword.Kind())
	}
	keyword = rkeyword.String()

	typ = "" // this it the intentional default
	atype, ok := keywordMap[KeyType]
	if ok {
		rtype := reflect.ValueOf(atype)
		if rtype.Kind() != reflect.String {
			return "", "",
				fmt.Errorf("for %q: expecting the type attribute to be a string, got a %s",
					fullName, rtype.Kind())
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
func exactStringValidator(expected string) func(v any, strict bool, parent string) error {
	return func(v any, strict bool, parent string) error {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			return fmt.Errorf("expecting a string value, got %q", rv.Kind())
		}
		s := rv.String()

		if s != expected {
			return fmt.Errorf("expecting value to be %q, got %q",
				expected, s)
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
			return fmt.Errorf("expecting an int64 value, got %s", rv.Kind())
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

		return validateRecursiveMap(m, strict, parent)
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

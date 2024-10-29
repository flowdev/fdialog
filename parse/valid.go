package parse

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"

	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
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

type KeywordType struct {
	keyword string
	typ     string
}

type KeywordValueType struct {
	attributes map[string]AttributeValueType
	validate   func(v any) error
}

type AttributeValidator func(v any, strict bool, parent []string) (any, error)
type AttributeValueType struct {
	required bool
	validate AttributeValidator
}

var validKeywords map[KeywordType]KeywordValueType

// need this init function, or we would get an initialization cycle :(
func init() {
	nameRegex := regexp.MustCompile(`^[\pL\pN_]+$`)
	linkRegex := regexp.MustCompile(`^[\pL\pN_]+(?:[.][\pL\pN_]+)*$`)

	validKeywords = map[KeywordType]KeywordValueType{
		KeywordType{KeywordWindow, ""}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordWindow),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				"title": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"width": {
					validate: valid.FloatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: valid.FloatValidator(80.0, math.MaxFloat32),
				},
				KeyChildren: {
					validate: ChildrenValidator(0, math.MaxInt),
				},
			},
		},
		KeywordType{KeywordLink, ""}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordLink),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				"destination": {
					required: true,
					validate: valid.StringValidator(1, 0, linkRegex),
				},
			},
		},
		KeywordType{KeywordDialog, "info"}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: valid.ExactStringValidator("info"),
				},
				"title": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"message": {
					required: true,
					validate: valid.StringValidator(1, 0, nil),
				},
				"buttonText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"width": {
					validate: valid.FloatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: valid.FloatValidator(80.0, math.MaxFloat32),
				},
			},
		},
		KeywordType{KeywordDialog, "error"}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: valid.ExactStringValidator("error"),
				},
				"message": {
					required: true,
					validate: valid.StringValidator(1, 0, nil),
				},
				"buttonText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"width": {
					validate: valid.FloatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: valid.FloatValidator(80.0, math.MaxFloat32),
				},
			},
		},
		KeywordType{KeywordDialog, "confirmation"}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: valid.ExactStringValidator("confirmation"),
				},
				"title": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"message": {
					required: true,
					validate: valid.StringValidator(1, 0, nil),
				},
				"dismissText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"confirmText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"width": {
					validate: valid.FloatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: valid.FloatValidator(80.0, math.MaxFloat32),
				},
				KeyChildren: {
					required: true,
					validate: ChildrenValidator(2, 2),
				},
			},
		},
		KeywordType{KeywordDialog, "openFile"}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordDialog),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: valid.ExactStringValidator("openFile"),
				},
				"extensions": {
					validate: valid.StringValidator(2, 0, nil),
				},
				"dismissText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"confirmText": {
					validate: valid.StringValidator(1, 0, nil),
				},
				"width": {
					validate: valid.FloatValidator(50.0, math.MaxFloat32),
				},
				"height": {
					validate: valid.FloatValidator(80.0, math.MaxFloat32),
				},
				//KeyChildren: {
				//	required: true,
				//	validate: ChildrenValidator(2, 2),
				//},
			},
		},
		KeywordType{KeywordAction, "exit"}: {
			attributes: map[string]AttributeValueType{
				KeyKeyword: {
					required: true,
					validate: valid.ExactStringValidator(KeywordAction),
				},
				KeyName: {
					required: true,
					validate: valid.StringValidator(1, 0, nameRegex),
				},
				KeyType: {
					required: true,
					validate: valid.ExactStringValidator("exit"),
				},
				"code": {
					validate: valid.IntValidator(0, 125),
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
	_, err := validateRecursiveMap(uiDescr, strict, nil)
	return err
}

func validateRecursiveMap(m map[string]map[string]any, strict bool, parent []string) (any, error) {
	var errs []error

	for name, keywordMap := range m {
		keywordMap[KeyName] = name
		fullName := append(parent, name)
		keyword, typ, err := GetKeywordType(keywordMap, fullName)
		if err != nil {
			errs = append(errs, err)
		} else {
			errs = append(errs, validateKeyword(keyword, fullName, typ, keywordMap, strict))
		}
	}
	return m, errors.Join(errs...)
}

func validateKeyword(
	keyword string, fullName []string, typ string,
	valueMap map[string]any,
	strict bool,
) error {
	keywordTypeValidationData, ok := validKeywords[KeywordType{keyword: keyword, typ: typ}]
	if !ok && typ != "" { // try empty type; will error later if not supported
		keywordTypeValidationData, ok = validKeywords[KeywordType{keyword: keyword, typ: ""}]
	}
	if !ok {
		return fmt.Errorf("for %q: the combination of keyword %q and type %q is not supported",
			ui.DisplayName(fullName), keyword, typ)
	}

	return validateAttributes(fullName, valueMap, keywordTypeValidationData.attributes, strict)
}

func validateAttributes(
	fullName []string,
	valueMap map[string]any,
	attributes map[string]AttributeValueType,
	strict bool,
) error {

	validatedAttributes := make(map[string]bool, len(attributes))
	errs := make([]error, len(attributes))

	for attrName, attribute := range attributes {
		if vv, ok := valueMap[attrName]; ok {
			validatedAttributes[attrName] = true
			vv, err := attribute.validate(vv, strict, fullName)
			if err != nil {
				errs = append(errs, fmt.Errorf("for %q, attribute %q: %v",
					ui.DisplayName(fullName), attrName, err.Error()))
			}
			valueMap[attrName] = vv
		} else if attribute.required {
			errs = append(errs, fmt.Errorf("for %q, is attribute %q required",
				ui.DisplayName(fullName), attrName))
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

		err := fmt.Errorf("for %q: these attributes aren't recognized: %s", ui.DisplayName(fullName), keysTooMuch)
		if strict {
			errs = append(errs, err)
		} else {
			log.Printf("WARNING: %v", err)
		}
	}
	return errors.Join(errs...)
}

func GetKeywordType(keywordMap map[string]any, fullName []string) (keyword, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[KeyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "",
			fmt.Errorf("for %q: expecting the keyword to be a string, got a %s",
				ui.DisplayName(fullName), rkeyword.Kind())
	}
	keyword = rkeyword.String()

	typ = "" // this it the intentional default
	atype, ok := keywordMap[KeyType]
	if ok {
		rtype := reflect.ValueOf(atype)
		if rtype.Kind() != reflect.String {
			return "", "",
				fmt.Errorf("for %q: expecting the type attribute to be a string, got a %s",
					ui.DisplayName(fullName), rtype.Kind())
		}
		typ = rtype.String()
	}

	return keyword, typ, nil
}

func ChildrenValidator(minLen, maxLen int) func(v any, strict bool, parent []string) (any, error) {
	return func(v any, strict bool, parent []string) (any, error) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Map {
			return v, fmt.Errorf("expecting a map value, got %s", rv.Kind())
		}

		m, ok := v.(map[string]map[string]any)
		if !ok {
			return v, fmt.Errorf("expecting a map[string]map[string]any value, got %T", v)
		}

		if len(m) < minLen {
			return v, fmt.Errorf("expecting at least %d map elements, got %d", minLen, len(m))
		}
		if len(m) > maxLen {
			return v, fmt.Errorf("expecting at most %d map elements, got %d", maxLen, len(m))
		}

		return validateRecursiveMap(m, strict, parent)
	}
}

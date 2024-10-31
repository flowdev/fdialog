package valid

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"

	"github.com/flowdev/fdialog/ui"
)

// UIDescription validates the data from a whole UI description file independent of its format.
// If strict is true additional attributes are errors.
// The keys of the first level map are the names of the windows, containers, ...
// The keys of the second level map are the attributes of the keyword map.
// Mandatory key for keyword maps is: "keyword"
// The key "type" is expected for most keywords but not for all.
func UIDescription(uiDescr ui.CommandsDescr, strict bool) error {
	ensureLongKeywords(uiDescr)
	_, err := validateRecursiveMap(uiDescr, strict, nil)
	return err
}

func ensureLongKeywords(descr ui.CommandsDescr) {
	for _, attributesDescr := range descr {
		ui.EnsureLongKeyword(attributesDescr)
		if achildren, ok := attributesDescr[ui.KeyChildren]; ok {
			if children, ok := achildren.(ui.CommandsDescr); ok {
				ensureLongKeywords(children)
			}
		}
	}
}

// ---------------------------------------------------------------------------
// AttributeValidator(s): func(v any, strict bool, parent []string) (any, error)
//

func StringValidator(minLen, maxLen int, regex *regexp.Regexp) ui.AttributeValidator {
	return func(v any, strict bool, parent []string) (any, error) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			return v, fmt.Errorf("expecting a string value, got %s", rv.Kind())
		}
		s := rv.String()

		if minLen > 0 && len(s) < minLen {
			return s, fmt.Errorf("string too short (min %d > actual %d)", minLen, len(s))
		}
		if maxLen > 0 && len(s) > maxLen {
			return s, fmt.Errorf("string too long (max %d < actual %d)", maxLen, len(s))
		}

		if regex != nil && !regex.MatchString(s) {
			return s, fmt.Errorf("string %q does not match pattern %q", s, regex.String())
		}

		return s, nil
	}
}
func ExactStringValidator(expected string) ui.AttributeValidator {
	return func(v any, strict bool, parent []string) (any, error) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			return v, fmt.Errorf("expecting a string value, got %q", rv.Kind())
		}
		s := rv.String()

		if s != expected {
			return s, fmt.Errorf("expecting value to be %q, got %q",
				expected, s)
		}

		return s, nil
	}
}

func IntValidator(minVal, maxVal int64) ui.AttributeValidator {
	return func(v any, strict bool, parent []string) (any, error) {
		var i int64
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Float64 {
			f := rv.Float()
			i = int64(f)
			if f != float64(i) {
				return v, fmt.Errorf("expecting an int64 (or a float64 convertable to it), got %f", f)
			}
		} else if rv.Kind() != reflect.Int64 {
			return v, fmt.Errorf("expecting an int64 value, got %s", rv.Kind())
		} else {
			i = rv.Int()
		}

		if i < minVal {
			return i, fmt.Errorf("integer value too small (min %d > actual %d)", minVal, i)
		}
		if i > maxVal {
			return i, fmt.Errorf("integer value too big (max %d < actual %d)", maxVal, i)
		}
		return i, nil
	}
}

func FloatValidator(minVal, maxVal float64) ui.AttributeValidator {
	return func(v any, strict bool, parent []string) (any, error) {
		var f float64
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Float64 {
			if rv.Kind() != reflect.Int64 {
				return v, fmt.Errorf("expecting a float64 value, got %s", rv.Kind())
			}
			f = float64(rv.Int()) // treat ints as floats as they are automatically recognized
		} else {
			f = rv.Float()
		}

		if !math.IsNaN(f) && f < minVal {
			return f, fmt.Errorf("float value too small (min %f > actual %f)", minVal, f)
		}
		if !math.IsNaN(f) && f > maxVal {
			return f, fmt.Errorf("float value too big (max %f < actual %f)", maxVal, f)
		}
		return f, nil
	}
}

func BoolValidator() func(v any, strict bool, parent []string) (any, error) {
	return func(v any, strict bool, parent []string) (any, error) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Bool {
			return v, fmt.Errorf("expecting a boolean value, got %s", rv.Kind())
		}
		return v, nil
	}
}

func ChildrenValidator(minLen, maxLen int) ui.AttributeValidator {
	return func(v any, strict bool, parent []string) (any, error) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Map {
			return v, fmt.Errorf("expecting a map value, got %s", rv.Kind())
		}

		m, ok := v.(ui.CommandsDescr)
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

// ---------------------------------------------------------------------------
// Helpers
//

func validateRecursiveMap(m ui.CommandsDescr, strict bool, parent []string) (any, error) {
	var errs []error

	for name, keywordMap := range m {
		keywordMap[ui.KeyName] = name
		fullName := append(parent, name)
		keyword, typ, err := getKeywordType(keywordMap, fullName)
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
	valueMap ui.AttributesDescr,
	strict bool,
) error {
	keywordTypeValidationData, ok := ui.KeywordValidData(keyword, typ)
	if !ok && typ != "" { // try empty type; will error later if not supported
		keywordTypeValidationData, ok = ui.KeywordValidData(keyword, "")
	}
	if !ok {
		return fmt.Errorf("for %q: the combination of keyword %q and type %q is not supported",
			ui.DisplayName(fullName), keyword, typ)
	}

	return validateAttributes(fullName, valueMap, keywordTypeValidationData.Attributes, strict)
}

func validateAttributes(
	fullName []string,
	valueMap ui.AttributesDescr,
	attributes map[string]ui.AttributeValueType,
	strict bool,
) error {

	validatedAttributes := make(map[string]bool, len(attributes))
	errs := make([]error, len(attributes))

	for attrName, attribute := range attributes {
		if vv, ok := valueMap[attrName]; ok {
			validatedAttributes[attrName] = true
			vv, err := attribute.Validate(vv, strict, fullName)
			if err != nil {
				errs = append(errs, fmt.Errorf("for %q, attribute %q: %v",
					ui.DisplayName(fullName), attrName, err.Error()))
			}
			valueMap[attrName] = vv
		} else if attribute.Required {
			errs = append(errs, fmt.Errorf("for %q, is attribute %q Required",
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

func getKeywordType(keywordMap ui.AttributesDescr, fullName []string) (keyword, typ string, err error) {
	rkeyword := reflect.ValueOf(keywordMap[ui.KeyKeyword])
	if rkeyword.Kind() != reflect.String {
		return "", "",
			fmt.Errorf("for %q: expecting the keyword to be a string, got a %s",
				ui.DisplayName(fullName), rkeyword.Kind())
	}
	keyword = rkeyword.String()

	typ = "" // this is the intentional default
	atype, ok := keywordMap[ui.KeyType]
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

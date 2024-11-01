package valid

import (
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
// The mandatory key for keyword maps is: "keyword"
// The key "type" is expected for most keywords but not for all.
func UIDescription(uiDescr ui.CommandsDescr, strict bool) bool {
	ok1 := PreprocessUIDescription(uiDescr, "")
	_, ok2 := validateRecursiveMap(uiDescr, strict, "")
	return ok1 && ok2
}

func PreprocessUIDescription(descr ui.CommandsDescr, parent string) bool {
	ok := true
	for name, attributesDescr := range descr {
		fullName := ui.FullNameFor(parent, name)
		ok = ok && ui.PreprocessAttributesDescription(attributesDescr, fullName)
		if children, ok2 := attributesDescr[ui.KeyChildren].(ui.CommandsDescr); ok2 {
			ok = ok && PreprocessUIDescription(children, fullName)
		}
	}
	return ok
}

// ---------------------------------------------------------------------------
// AttributeValidator(s): func(v any, strict bool, parent []string) (any, error)
//

func StringValidator(minLen, maxLen int, regex *regexp.Regexp) ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		ok := true
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			log.Printf("ERROR: for %q: expecting a string value, got %s", parent, rv.Kind())
			return v, false
		}
		s := rv.String()

		if minLen > 0 && len(s) < minLen {
			log.Printf("ERROR: for %q: string too short (min %d > actual %d)", parent, minLen, len(s))
			ok = false
		}
		if maxLen > 0 && len(s) > maxLen {
			log.Printf("ERROR: for %q: string too long (max %d < actual %d)", parent, maxLen, len(s))
			ok = false
		}

		if regex != nil && !regex.MatchString(s) {
			log.Printf("ERROR: for %q: string %q does not match pattern %q", parent, s, regex.String())
			ok = false
		}

		return s, ok
	}
}
func ExactStringValidator(expected string) ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.String {
			log.Printf("ERROR: for %q: expecting a string value, got %q", parent, rv.Kind())
			return v, false
		}
		s := rv.String()

		if s != expected {
			log.Printf("ERROR: for %q: expecting value to be %q, got %q", parent, expected, s)
			return s, false
		}

		return s, true
	}
}

func IntValidator(minVal, maxVal int64) ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		var i int64
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Float64 {
			f := rv.Float()
			i = int64(f)
			if f != float64(i) {
				log.Printf("ERROR: for %q: expecting an int64 (or a float64 convertable to it), got %f",
					parent, f)
				return v, false
			}
		} else if rv.Kind() != reflect.Int64 {
			log.Printf("ERROR: for %q: expecting an int64 value, got %s", parent, rv.Kind())
			return v, false
		} else {
			i = rv.Int()
		}

		ok := true
		if i < minVal {
			log.Printf("ERROR: for %q: integer value too small (min %d > actual %d)", parent, minVal, i)
			ok = false
		}
		if i > maxVal {
			log.Printf("ERROR: for %q: integer value too big (max %d < actual %d)", parent, maxVal, i)
			ok = false
		}
		return i, ok
	}
}

func FloatValidator(minVal, maxVal float64) ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		var f float64
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Float64 {
			if rv.Kind() != reflect.Int64 {
				log.Printf("ERROR: for %q: expecting a float64 value, got %s", parent, rv.Kind())
				return v, false
			}
			f = float64(rv.Int()) // treat ints as floats as they are automatically recognized
		} else {
			f = rv.Float()
		}

		ok := true
		if math.IsNaN(f) {
			log.Printf("ERROR: for %q: float value expected, got NaN (Not a Number)", parent)
			ok = false
		}
		if f < minVal {
			log.Printf("ERROR: for %q: float value too small (min %f > actual %f)", parent, minVal, f)
			ok = false
		}
		if f > maxVal {
			log.Printf("ERROR: for %q: float value too big (max %f < actual %f)", parent, maxVal, f)
			ok = false
		}
		return f, ok
	}
}

func BoolValidator() ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Bool {
			log.Printf("ERROR: for: %q: expecting a boolean value, got %s", parent, rv.Kind())
			return v, false
		}
		return v, true
	}
}

func ChildrenValidator(minLen, maxLen int) ui.AttributeValidator {
	return func(v any, strict bool, parent string) (any, bool) {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Map {
			log.Printf("ERROR: for %q: expecting a map value, got %s", parent, rv.Kind())
			return v, false
		}

		m, ok := v.(ui.CommandsDescr)
		if !ok {
			log.Printf("ERROR: for %q: expecting a map[string]map[string]any value, got %T", parent, v)
			return v, false
		}

		if len(m) < minLen {
			log.Printf("ERROR: for %q: expecting at least %d map elements, got %d", parent, minLen, len(m))
			ok = false
		}
		if len(m) > maxLen {
			log.Printf("ERROR: for %q: expecting at most %d map elements, got %d", parent, maxLen, len(m))
			ok = false
		}

		val, ok2 := validateRecursiveMap(m, strict, parent)
		return val, ok2 && ok
	}
}

// ---------------------------------------------------------------------------
// Helpers
//

func validateRecursiveMap(m ui.CommandsDescr, strict bool, parent string) (any, bool) {
	ok := true
	for name, keywordMap := range m {
		keywordMap[ui.KeyName] = name
		fullName := ui.FullNameFor(parent, name)
		keyword, typ, ok2 := getKeywordType(keywordMap, fullName)
		ok = ok && ok2
		ok = ok && validateKeyword(keyword, fullName, typ, keywordMap, strict)
	}
	return m, ok
}

func validateKeyword(keyword string, fullName string, typ string, valueMap ui.AttributesDescr, strict bool) bool {
	keywordTypeValidationData, ok := ui.KeywordValidData(keyword, typ)
	if !ok && typ != "" { // try empty type; will error later if not supported
		keywordTypeValidationData, ok = ui.KeywordValidData(keyword, "")
	}
	if !ok {
		log.Printf("ERROR: for %q: the combination of keyword %q and type %q is not supported",
			fullName, keyword, typ)
		return false
	}

	return validateAttributes(fullName, valueMap, keywordTypeValidationData.Attributes, strict)
}

func validateAttributes(
	fullName string,
	valueMap ui.AttributesDescr,
	attributes map[string]ui.AttributeValueType,
	strict bool,
) bool {
	validatedAttributes := make(map[string]bool, len(attributes))
	ok := true

	for attrName, attribute := range attributes {
		if vv, ok2 := valueMap[attrName]; ok2 {
			validatedAttributes[attrName] = true
			vv, ok3 := attribute.Validate(vv, strict, ui.FullNameFor(fullName, attrName))
			if !ok3 {
				ok = false
			}
			valueMap[attrName] = vv
		} else if attribute.Required {
			log.Printf("for %q, is attribute %q Required", fullName, attrName)
			ok = false
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

		err := fmt.Errorf("for %q: these attributes aren't recognized: %s", fullName, keysTooMuch)
		if strict {
			log.Println("ERROR:", err)
			ok = false
		} else {
			log.Println("WARNING:", err)
		}
	}
	return ok
}

func getKeywordType(keywordMap ui.AttributesDescr, fullName string) (keyword, typ string, ok bool) {
	rkeyword := reflect.ValueOf(keywordMap[ui.KeyKeyword])
	if rkeyword.Kind() != reflect.String {
		log.Printf("ERROR: for %q: expecting the keyword to be a string, got a %s", fullName, rkeyword.Kind())
		return "", "", false
	}
	keyword = rkeyword.String()

	typ = "" // this is the intentional default
	atype, ok := keywordMap[ui.KeyType]
	if ok {
		rtype := reflect.ValueOf(atype)
		if rtype.Kind() != reflect.String {
			log.Printf("ERROR: for %q: expecting the type attribute to be a string, got a %s",
				fullName, rtype.Kind())
			return "", "", false
		}
		typ = rtype.String()
	}

	return keyword, typ, true
}

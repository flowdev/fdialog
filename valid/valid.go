package valid

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"

	"github.com/flowdev/fdialog/ui"
)

var validateName = StringValidator(1, 0, ui.NameRegex)
var validateID = StringValidator(1, 0, ui.NameRegex)
var validateGroup = StringValidator(1, 0, ui.LinkRegex)
var ValidateOutputKey = ui.AttributeValueType{
	Validate: StringValidator(1, 0, ui.LinkRegex),
}

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
	for name, attrs := range descr.All() {
		fullName := ui.FullNameFor(parent, name)
		attrs[ui.KeyName] = name
		ok = ok && ui.PreprocessAttributesDescription(attrs, fullName)
		if children, ok2 := attrs[ui.KeyChildren].(ui.CommandsDescr); ok2 {
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
		if rv.Kind() != reflect.Ptr {
			log.Printf(`ERROR: for %q: expecting a pointer value for "children", got %s`, parent, rv.Kind())
			return v, false
		}

		m, ok := v.(ui.CommandsDescr)
		if !ok {
			log.Printf("ERROR: for %q: expecting a omap.OrderedMap[string, map[string]any] value, got %T",
				parent, v)
			return v, false
		}

		if m.Len() < minLen {
			log.Printf("ERROR: for %q: expecting at least %d map elements, got %d", parent, minLen, m.Len())
			ok = false
		}
		if m.Len() > maxLen {
			log.Printf("ERROR: for %q: expecting at most %d map elements, got %d", parent, maxLen, m.Len())
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
	for name, attrs := range m.All() {
		fullName := ui.FullNameFor(parent, name)
		keyword, typ, ok2 := getKeywordType(attrs, fullName)
		ok = ok && ok2
		ok = ok && validateKeyword(keyword, fullName, typ, attrs, strict)
	}
	return m, ok
}

func validateKeyword(keyword string, fullName string, typ string, valueMap ui.AttributesDescr, strict bool) bool {
	commandValidationData, ok := ui.KeywordValidData(keyword, typ)
	if !ok && typ != "" { // try empty type; will error later if not supported
		commandValidationData, ok = ui.KeywordValidData(keyword, "")
	}
	if !ok {
		log.Printf("ERROR: for %q: the combination of keyword %q and type %q is not supported",
			fullName, keyword, typ)
		return false
	}

	if validate := commandValidationData.Validate; validate != nil {
		ok = validate(valueMap, fullName)
	}
	return validateAttributes(valueMap, commandValidationData.Attributes, strict, fullName) && ok
}

func validateAttributes(
	valueMap ui.AttributesDescr,
	attributes map[string]ui.AttributeValueType,
	strict bool,
	parent string,
) bool {
	validatedAttributes := make(map[string]bool, len(attributes))
	ok := true

	for attrName, attribute := range attributes {
		if value, ok2 := valueMap[attrName]; ok2 {
			validatedAttributes[attrName] = true
			fullName := parent
			if attrName != ui.KeyChildren {
				fullName = ui.FullNameFor(parent, attrName)
			}
			v, ok3 := attribute.Validate(value, strict, fullName)
			ok = ok && ok3
			valueMap[attrName] = v
		} else if attribute.Required {
			log.Printf("for %q: attribute %q is required", parent, attrName)
			ok = false
		}
	}

	if value, ok2 := valueMap[ui.KeyName]; ok2 {
		validatedAttributes[ui.KeyName] = true
		fullName := ui.FullNameFor(parent, ui.KeyName)
		_, ok3 := validateName(value, strict, fullName)
		ok = ok && ok3
	} else {
		log.Printf(`for %q: attribute "name" is required`, parent)
		ok = false
	}

	if len(validatedAttributes) != len(valueMap) {
		unknownKeys := make([]string, 0, len(valueMap)-len(validatedAttributes))

	forLoop:
		for k, v := range valueMap {
			_, ok := validatedAttributes[k]
			if !ok {
				switch k {
				case ui.KeyID:
					validateID(v, strict, ui.FullNameFor(parent, k))
					continue forLoop // id is always allowed
				case ui.KeyGroup:
					validateGroup(v, strict, ui.FullNameFor(parent, k))
					continue forLoop // group is always allowed
				case ui.KeyName:
					validateName(v, strict, ui.FullNameFor(parent, k))
					continue forLoop // name is always required
				}
				unknownKeys = append(unknownKeys, k)
			}
		}

		if len(unknownKeys) > 0 {
			err := fmt.Errorf("for %q: these attributes are unknown: %s", parent, unknownKeys)
			if strict {
				log.Println("ERROR:", err)
				ok = false
			} else {
				log.Println("WARNING:", err)
			}
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

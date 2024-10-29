package ui

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"fyne.io/fyne/v2"
)

// Reserved attribute names:
const (
	KeyKeyword  = "keyword"
	KeyName     = "name"
	KeyChildren = "children"
	KeyType     = "type" // type is used like an ordinary attribute, but it has special semantics
)

// Basic keywords:
const (
	KeywordWindow = "window"
	KeywordAction = "action"
	KeywordLink   = "link"
)

const WinMain = "main"

// ---------------------------------------------------------------------------
//  Run Types & Data

type RunFunction func(
	detailDescr map[string]any,
	fullName []string,
	win fyne.Window,
	completeDescr map[string]map[string]any,
) error

var keywordShortToLong = make(map[string]string, 64)
var keywordLongToShort = make(map[string]string, 64)
var keywordMap = make(map[string]RunFunction, 64)

var actionMap = make(map[string]RunFunction, 32)

// idMap maps an ID to a full name path.
var idMap = make(map[string][]string, 32)

// ---------------------------------------------------------------------------
//  Validation Types & Data

type ValidKeywordType struct {
	Keyword string
	Type    string
}

type ValidAttributesType struct {
	Attributes map[string]AttributeValueType
	Validate   func(attrs map[string]any) error
}

type AttributeValidator func(v any, strict bool, parent []string) (any, error)
type AttributeValueType struct {
	Required bool
	Validate AttributeValidator
}

// validKeywords is the big map used for keyword validation
var validKeywords = make(map[ValidKeywordType]ValidAttributesType, 64)

var NameRegex = regexp.MustCompile(`^[\pL\pN_]+$`)
var LinkRegex = regexp.MustCompile(`^[\pL\pN_]+(?:[.][\pL\pN_]+)*$`)

// ---------------------------------------------------------------------------
//  Registration & Access

// RegisterRunKeyword registers a keyword with long and (potentially) short name
// and run function.
func RegisterRunKeyword(longKW, shortKW string, runFunc RunFunction) error {
	if longKW == "" && shortKW == "" {
		return errors.New("unable to register empty keyword")
	}
	if runFunc == nil {
		return fmt.Errorf("unable to register keyword (%q/%q) with nil function", longKW, shortKW)
	}

	if len(longKW) < len(shortKW) {
		longKW, shortKW = shortKW, longKW // swap
	}

	if _, ok := keywordMap[longKW]; ok {
		return fmt.Errorf("keyword with long name %q exists already", longKW)
	}
	keywordMap[longKW] = runFunc

	if shortKW != "" && shortKW != longKW {
		if _, ok := keywordMap[shortKW]; ok {
			return fmt.Errorf("keyword with short name %q exists already", shortKW)
		}
		keywordLongToShort[longKW] = shortKW
		keywordShortToLong[shortKW] = longKW
	}
	return nil
}

// KeywordRunFunc returns the run function for a registered keyword.
// It returns `false` if nothing was found.
func KeywordRunFunc(keyword string) (runFunc RunFunction, ok bool) {
	runFunc, ok = keywordMap[keyword]
	return runFunc, ok
}

// RegisterAction registers an action with a name and a run function.
func RegisterAction(name string, runFunc RunFunction) error {
	if name == "" {
		return errors.New("unable to register an empty action")
	}
	if runFunc == nil {
		return fmt.Errorf("unable to register action %q with nil function", name)
	}

	actionMap[name] = runFunc
	return nil
}

// ActionRunFunc returns the run function for a registered action.
// It returns `false` if nothing was found.
func ActionRunFunc(name string) (runFunc RunFunction, ok bool) {
	runFunc, ok = actionMap[name]
	return runFunc, ok
}

func RegisterValidKeyword(keyword, typ string, validKWMap ValidAttributesType) error {
	_, ok := validKeywords[ValidKeywordType{Keyword: keyword, Type: typ}]
	if ok {
		return fmt.Errorf("keyword %q with type %q already exists", keyword, typ)
	}
	validKeywords[ValidKeywordType{Keyword: keyword, Type: typ}] = validKWMap
	return nil
}

// KeywordValidData returns the validation data for a registered
// keyword, type combination.
// It returns `false` if nothing was found.
func KeywordValidData(keyword, typ string) (ValidAttributesType, bool) {
	validFunc, ok := validKeywords[ValidKeywordType{Keyword: keyword, Type: typ}]
	return validFunc, ok
}

// RegisterID registers an ID as a shortcut for the fullName.
func RegisterID(id string, fullName []string) error {
	if _, ok := idMap[id]; ok {
		return fmt.Errorf("ID %q already exists", id)
	}
	idMap[id] = fullName
	return nil
}

// FullNameForID returns the fullName for an ID.
// It returns `false` if nothing was found.
func FullNameForID(id string) ([]string, bool) {
	fullName, ok := idMap[id]
	return fullName, ok
}

// ---------------------------------------------------------------------------
//  Helpers

func SameFullName(a []string, b ...string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

func DisplayName(fullName []string) string {
	return strings.Join(fullName, ".")
}

func FullName(displayName string) []string {
	return strings.Split(displayName, ".")
}

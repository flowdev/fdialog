package ui

import (
	"errors"
	"fmt"
	"log"
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
	KeyID       = "id"
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

type AttributesDescr map[string]any

type CommandsDescr map[string]AttributesDescr

type RunFunction func(
	detailDescr AttributesDescr,
	fullName string,
	win fyne.Window,
	completeDescr CommandsDescr,
)

var keywordShortToLong = make(map[string]string, 64)
var keywordLongToShort = make(map[string]string, 64)
var keywordMap = make(map[string]RunFunction, 64)

var actionMap = make(map[string]RunFunction, 32)

// idMap maps an ID to a full name path.
var idMap = make(map[string]string, 32)

// valueMap maps a fullName to an input value
// fullNames are the display names with '.' inside.
var valueMap = make(map[string]map[string]any)

// ---------------------------------------------------------------------------
//  Validation Types & Data

type ValidKeywordType struct {
	Keyword string
	Type    string
}

type ValidAttributesType struct {
	Attributes map[string]AttributeValueType
	Validate   func(attrs AttributesDescr) bool
}

type AttributeValidator func(v any, strict bool, parent string) (any, bool)
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
		if _, ok := keywordShortToLong[shortKW]; ok {
			return fmt.Errorf("keyword with short name %q exists already", shortKW)
		}
		keywordLongToShort[longKW] = shortKW
		keywordShortToLong[shortKW] = longKW
	}
	return nil
}

// RunFuncForKeyword returns the run function for a registered keyword.
// It returns `false` if nothing was found.
func RunFuncForKeyword(keyword string) (runFunc RunFunction, ok bool) {
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
// The keyword might be the shortened variant.
// It returns `false` if nothing was found.
func KeywordValidData(keyword, typ string) (ValidAttributesType, bool) {
	longKW, ok := keywordShortToLong[keyword]
	if !ok {
		longKW = keyword
	}
	validFunc, ok := validKeywords[ValidKeywordType{Keyword: longKW, Type: typ}]
	return validFunc, ok
}

// RegisterID registers an ID as a shortcut for the fullName.
func RegisterID(id string, fullName string) error {
	if _, ok := idMap[id]; ok {
		return fmt.Errorf("ID %q already exists", id)
	}
	idMap[id] = fullName
	return nil
}

// FullNameForID returns the full display name for an ID.
// It returns `false` if nothing was found.
func FullNameForID(id string) (string, bool) {
	fullName, ok := idMap[id]
	return fullName, ok
}

func GetValueByID(id, group string) (any, bool) {
	grpMap, ok := valueMap[group]
	if !ok {
		return nil, false
	}
	v, ok := grpMap[idMap[id]]
	return v, ok
}

func GetValueByFullName(fullName, group string) (any, bool) {
	grpMap, ok := valueMap[group]
	if !ok {
		return nil, false
	}
	v, ok := grpMap[fullName]
	return v, ok
}

func GetValueGroup(group string) (map[string]any, bool) {
	grpMap, ok := valueMap[group]
	return grpMap, ok
}

func StoreValueByID(value any, id, group string, parent string) {
	if name, ok := FullNameForID(id); ok {
		grpMap, ok := valueMap[group]
		if !ok {
			grpMap = make(map[string]any)
			valueMap[group] = grpMap
		}
		grpMap[name] = value
		return
	}
	log.Printf(`ERROR: for %q: unknown ID: %q`, parent, id)
}

func StoreValueByFullName(value any, fullName, group string) {
	grpMap, ok := valueMap[group]
	if !ok {
		grpMap = make(map[string]any)
		valueMap[group] = grpMap
	}
	grpMap[fullName] = value
}

// ---------------------------------------------------------------------------
//  Helpers

// PreprocessAttributesDescription prepares an attributes description for
// validation and running. Specifically it:
//   - converts known short keywords to their long counterparts and
//   - registers all IDs.
//
// `false` is returned if an error occurs.
func PreprocessAttributesDescription(descr AttributesDescr, fullName string) bool {
	var err error
	if shortKW, ok := descr[KeyKeyword].(string); ok {
		if longKW, ok := keywordShortToLong[shortKW]; ok {
			descr[KeyKeyword] = longKW
		}
	}
	if id, ok := descr[KeyID].(string); ok {
		err = RegisterID(id, fullName)
		if err != nil {
			log.Printf("ERROR: for %q: %v", fullName, err)
		}
	}
	return err == nil
}

func SplitName(fullName string) []string {
	return strings.Split(fullName, ".")
}

func FullNameFor(parent, name string) string {
	if parent == "" {
		return name
	}
	return parent + "." + name
}

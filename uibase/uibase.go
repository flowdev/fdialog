package uibase

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
)

const WinMain = "main"

type KeywordFunction func(
	keywordDescr map[string]any,
	fullName string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
) error

// idMap maps an ID to a full name path.
var idMap map[string][]string = make(map[string][]string, 32)

var keywordShortToLong map[string]string = make(map[string]string, 64)
var keywordLongToShort map[string]string = make(map[string]string, 64)
var keywordMap map[string]KeywordFunction = make(map[string]KeywordFunction, 64)

var actionMap map[string]KeywordFunction = make(map[string]KeywordFunction, 32)

func RegisterKeyword(longKW, shortKW string, kwFunc KeywordFunction) error {
	if longKW == "" && shortKW == "" {
		return errors.New("unable to register empty keyword")
	}
	if kwFunc == nil {
		return fmt.Errorf("unable to register keyword (%q/%q) with nil function", longKW, shortKW)
	}

	if len(longKW) < len(shortKW) {
		longKW, shortKW = shortKW, longKW // swap
	}

	keywordMap[longKW] = kwFunc
	if shortKW != "" && shortKW != longKW {
		keywordMap[shortKW] = kwFunc
		keywordLongToShort[longKW] = shortKW
		keywordShortToLong[shortKW] = longKW
	}
	return nil
}

func KeywordFunc(keyword string) (kwFunc KeywordFunction, ok bool) {
	kwFunc, ok = keywordMap[keyword]
	return kwFunc, ok
}
func RegisterAction(name string, fn KeywordFunction) error {
	if name == "" {
		return errors.New("unable to register an empty action")
	}
	if fn == nil {
		return fmt.Errorf("unable to register action %q with nil function", name)
	}

	actionMap[name] = fn
	return nil
}

func ActionFunc(name string) (fn KeywordFunction, ok bool) {
	fn, ok = keywordMap[name]
	return fn, ok
}

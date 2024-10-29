package ui

import (
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
)

const WinMain = "main"

type RunFunction func(
	detailDescr map[string]any,
	fullName []string,
	win fyne.Window,
	completeDescr map[string]map[string]any,
) error

// idMap maps an ID to a full name path.
var idMap map[string][]string = make(map[string][]string, 32)

var keywordShortToLong map[string]string = make(map[string]string, 64)
var keywordLongToShort map[string]string = make(map[string]string, 64)
var keywordMap map[string]RunFunction = make(map[string]RunFunction, 64)

var actionMap map[string]RunFunction = make(map[string]RunFunction, 32)

// RegisterKeyword registers a keyword with long and (potentially) short name
// and run function.
func RegisterKeyword(longKW, shortKW string, runFunc RunFunction) error {
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
		keywordMap[shortKW] = runFunc
		keywordLongToShort[longKW] = shortKW
		keywordShortToLong[shortKW] = longKW
	}
	return nil
}

func KeywordRunFunc(keyword string) (runFunc RunFunction, ok bool) {
	runFunc, ok = keywordMap[keyword]
	return runFunc, ok
}

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

func ActionRunFunc(name string) (runFunc RunFunction, ok bool) {
	runFunc, ok = actionMap[name]
	return runFunc, ok
}

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
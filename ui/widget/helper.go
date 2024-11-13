package widget

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/flowdev/fdialog/ui"
	"log"
	"regexp"
	"strings"
)

// -----------------------------------------------------------------------
// Validators
//

func StringValidator(child ui.AttributesDescr, fullName string) fyne.StringValidator {
	i64, _ := child["minLen"].(int64)
	minLen := int(i64)
	i64, _ = child["maxLen"].(int64)
	maxLen := int(i64)
	regex, _ := child["regexp"].(string)
	failText, _ := child["failText"].(string)
	vErr := errors.New(failText)

	re, err := regexp.Compile(regex)
	if err != nil {
		log.Printf("ERROR: for %q: illegal regular expression for validation: %q", fullName, regex)
	}
	return func(s string) error {
		if minLen > 0 && len(s) < minLen {
			if failText != "" {
				return vErr
			}
			return fmt.Errorf("string too short (min %d > actual %d)", minLen, len(s))
		}
		if maxLen > 0 && len(s) > maxLen {
			if failText != "" {
				return vErr
			}
			return fmt.Errorf("string too long (max %d < actual %d)", maxLen, len(s))
		}
		if regex == "" {
			return nil
		}
		if err != nil {
			return err
		}
		if !re.MatchString(s) {
			if failText != "" {
				return vErr
			}
			return fmt.Errorf("string %q does not match pattern %q", s, regex)
		}
		return nil
	}
}

// -----------------------------------------------------------------------
// Helpers
//

func linesToSlice(a any) []string {
	str, ok := a.(string)
	if !ok {
		return nil
	}
	sl := strings.Split(str, "\n")
	for i := 0; i < len(sl); i++ {
		sl[i] = strings.TrimSpace(sl[i])
	}
	return sl
}

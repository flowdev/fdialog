package valid

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
)

func StringValidator(minLen, maxLen int, regex *regexp.Regexp) func(v any, strict bool, parent []string) (any, error) {
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
func ExactStringValidator(expected string) func(v any, strict bool, parent []string) (any, error) {
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

func IntValidator(minVal, maxVal int64) func(v any, strict bool, parent []string) (any, error) {
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

func FloatValidator(minVal, maxVal float64) func(v any, strict bool, parent []string) (any, error) {
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

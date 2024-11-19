package parse

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/valyala/fastjson"

	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/x/omap"
)

var jsonParser = &fastjson.ParserPool{}

// JSON parses JSON from a Reader and gives the content back suitable
// for validation.
// An error is returned if the stream can't be unmarshalled or a data type
// doesn't match.
// We are rather strict about the JSON standard as this is meant for fast
// machine to machine communication.
// Please use UIDL as a human friendly format.
func JSON(input io.Reader, name string) (ui.CommandsDescr, error) {
	inputData, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON from %q: %w", name, err)
	}
	parser := jsonParser.Get()
	defer jsonParser.Put(parser)
	val, err := parser.ParseBytes(inputData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON from %q: %w", name, err)
	}
	data, err := convertJSONCommands(val, "")
	if err != nil {
		return nil, fmt.Errorf("error converting JSON data from %q: %w", name, err)
	}
	return data, nil
}

func convertJSONCommands(val *fastjson.Value, parent string) (ui.CommandsDescr, error) {
	obj, err := val.Object()
	if err != nil {
		return nil, fmt.Errorf("for %q: error converting JSON object: %v", parent, err)
	}
	errs := make([]error, 0, 32)
	data := omap.New[string, ui.AttributesDescr](8)
	obj.Visit(func(k []byte, v *fastjson.Value) {
		name := string(k)
		attrs, err := convertJSONAttributes(v, ui.FullNameFor(parent, name))
		if err != nil {
			errs = append(errs, err)
		} else {
			if ok := data.Add(name, attrs); !ok {
				errs = append(errs, fmt.Errorf("for %q: command with name %q exists already", parent, name))
			}
		}
	})
	if len(errs) == 0 {
		return data, nil
	}
	return data, errors.Join(errs...)
}

func convertJSONAttributes(val *fastjson.Value, fullName string) (ui.AttributesDescr, error) {
	obj, err := val.Object()
	if err != nil {
		return nil, fmt.Errorf("for %q: error converting object: %w", fullName, err)
	}
	errs := make([]error, 0, 32)
	attrs := make(map[string]any)
	obj.Visit(func(k []byte, v *fastjson.Value) {
		name := string(k)
		attr, err := convertJSONValue(v, ui.FullNameFor(fullName, name))
		if err != nil {
			errs = append(errs, err)
		} else {
			attrs[name] = attr
		}
	})
	if len(errs) == 0 {
		return attrs, nil
	}
	return attrs, errors.Join(errs...)
}

func convertJSONValue(val *fastjson.Value, parent string) (any, error) {
	switch val.Type() {
	case fastjson.TypeFalse:
		return false, nil
	case fastjson.TypeTrue:
		return true, nil
	case fastjson.TypeString:
		return convertJSONString(val, parent)
	case fastjson.TypeObject:
		return convertJSONCommands(val, parent)
	case fastjson.TypeArray:
		return convertJSONArray(val, parent)
	case fastjson.TypeNumber:
		return convertJSONNumber(val, parent)
	default:
		return nil, fmt.Errorf("for %q: unable to convert JSON value type %s", parent, val.Type().String())
	}
}

func convertJSONSubValue(val *fastjson.Value, parent string) (any, error) {
	switch val.Type() {
	case fastjson.TypeFalse:
		return false, nil
	case fastjson.TypeTrue:
		return true, nil
	case fastjson.TypeString:
		return convertJSONString(val, parent)
	case fastjson.TypeNumber:
		return convertJSONNumber(val, parent)
	default:
		return nil, fmt.Errorf("for %q: unable to convert JSON sub-value type %s", parent, val.Type().String())
	}
}

func convertJSONString(val *fastjson.Value, fullName string) (any, error) {
	bs, err := val.StringBytes()
	if err != nil {
		return "", fmt.Errorf("for %q: converting string: %w", fullName, err)
	}
	if len(bs) == 0 {
		return "", nil
	}
	return string(bs), nil
}

func convertJSONNumber(val *fastjson.Value, fullName string) (any, error) {
	f, err := val.Float64()
	if err != nil {
		return nil, fmt.Errorf("for %q: %w", fullName, err)
	}
	i := int64(f)
	if float64(i) == f { // use int64 if representation is exact
		return i, nil
	}
	return f, nil
}

func convertJSONArray(val *fastjson.Value, fullName string) (any, error) {
	arr, err := val.Array()
	if err != nil {
		return nil, fmt.Errorf("for %q: error converting array: %w", fullName, err)
	}
	result := make([]any, len(arr))
	errs := make([]error, 0, len(arr))
	for i, v := range arr {
		converted, err := convertJSONSubValue(v, ui.FullNameFor(fullName, strconv.Itoa(i)))
		errs = append(errs, err)
		result[i] = converted
	}
	return result, errors.Join(errs...)
}

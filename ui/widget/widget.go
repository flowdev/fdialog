package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/flowdev/fdialog/run"
	"github.com/flowdev/fdialog/ui"
	"github.com/flowdev/fdialog/valid"
	"log"
	"math"
	"reflect"
	"regexp"
)

const KeywordForm = "form"
const KeywordFormItem = "item"

type Creator func(attrs ui.AttributesDescr, values map[string]any, outputKey, fullName string) fyne.CanvasObject

var URLRegex = regexp.MustCompile(`^http(s?)://[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(/?)([a-zA-Z0-9\-.?,'/\\+&%$#_]*)?$`)

var ScrollBarsRegex = regexp.MustCompile(`^both|horizontal|vertical|none$`)

var widgetMap = make(map[string]Creator, 64)

func RegisterAll() error {
	// -----------------------------------------------------------------------
	// Register Validators
	//

	err := ui.RegisterValidKeyword(KeywordForm, "", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordForm),
			},
			"submitText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"cancelText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			ui.AttrChildren: {
				Required: true,
				Validate: valid.ChildrenValidator(3, math.MaxInt),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "entry", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("entry"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"placeHolder": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"minLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"maxLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"regexp": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"failText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "multiLineEntry", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("multiLineEntry"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"placeHolder": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"minLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"maxLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"regexp": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"failText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "passwordEntry", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("passwordEntry"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"placeHolder": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"minLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"maxLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"regexp": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"failText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "checkBox", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("checkBox"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"subLabel": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "checkGroup", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("checkGroup"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"options": {
				Required: true,
				Validate: valid.ListValidator(1, math.MaxInt,
					valid.StringValidator(1, 0, nil)),
			},
			"initiallySelected": {
				Validate: valid.ListValidator(0, math.MaxInt,
					valid.StringValidator(1, 0, nil)),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "hyperlink", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("hyperlink"),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"text": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"url": {
				Required: true,
				Validate: valid.StringValidator(8, 0, URLRegex),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "radioGroup", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("radioGroup"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"options": {
				Required: true,
				Validate: valid.ListValidator(2, math.MaxInt,
					valid.StringValidator(1, 0, nil)),
			},
			"initiallySelected": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"horizontal": {
				Validate: valid.BoolValidator(),
			},
			"required": {
				Validate: valid.BoolValidator(),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "richText", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("richText"),
			},
			"label": {
				Validate: valid.StringValidator(0, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"text": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"scroll": {
				Validate: valid.StringValidator(1, 0, ScrollBarsRegex),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "select", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("select"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"placeHolder": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"options": {
				Required: true,
				Validate: valid.ListValidator(0, math.MaxInt,
					valid.StringValidator(1, 0, nil)),
			},
			"initiallySelected": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "selectEntry", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("selectEntry"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"placeHolder": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"options": {
				Required: true,
				Validate: valid.ListValidator(0, math.MaxInt,
					valid.StringValidator(1, 0, nil)),
			},
			"minLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"maxLen": {
				Validate: valid.IntValidator(0, math.MaxInt),
			},
			"regexp": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"failText": {
				Validate: valid.StringValidator(1, 0, nil),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "separator", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("separator"),
			},
		},
	})
	if err != nil {
		return err
	}

	err = ui.RegisterValidKeyword(KeywordFormItem, "slider", ui.ValidAttributesType{
		Attributes: map[string]ui.AttributeValueType{
			ui.AttrKeyword: {
				Required: true,
				Validate: valid.ExactStringValidator(KeywordFormItem),
			},
			ui.AttrType: {
				Required: true,
				Validate: valid.ExactStringValidator("slider"),
			},
			ui.AttrOutputKey: {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"label": {
				Required: true,
				Validate: valid.StringValidator(1, 0, nil),
			},
			"hint": {
				Validate: valid.StringValidator(1, 0, nil),
			},
			"disabled": {
				Validate: valid.BoolValidator(),
			},
			"min": {
				Validate: valid.FloatValidator(-math.MaxFloat64, math.MaxFloat64),
			},
			"max": {
				Validate: valid.FloatValidator(-math.MaxFloat64, math.MaxFloat64),
			},
			"step": {
				Validate: valid.FloatValidator(0.0, math.MaxFloat64),
			},
			"initialValue": {
				Validate: valid.FloatValidator(-math.MaxFloat64, math.MaxFloat64),
			},
		},
	})
	if err != nil {
		return err
	}

	// -----------------------------------------------------------------------
	// Register Runners
	//

	err = ui.RegisterRunKeyword(KeywordForm, "frm", runForm)
	if err != nil {
		return err
	}

	// -----------------------------------------------------------------------
	// Register Widgets
	//

	err = RegisterWidget(createEntry, "entry")
	if err != nil {
		return err
	}
	err = RegisterWidget(createMultiLineEntry, "multiLineEntry")
	if err != nil {
		return err
	}
	err = RegisterWidget(createPasswordEntry, "passwordEntry")
	if err != nil {
		return err
	}
	err = RegisterWidget(createCheckBox, "checkBox")
	if err != nil {
		return err
	}
	err = RegisterWidget(createCheckGroup, "checkGroup")
	if err != nil {
		return err
	}
	err = RegisterWidget(createHyperlink, "hyperlink")
	if err != nil {
		return err
	}
	err = RegisterWidget(createRadioGroup, "radioGroup")
	if err != nil {
		return err
	}
	err = RegisterWidget(createSelect, "select")
	if err != nil {
		return err
	}
	err = RegisterWidget(createRichText, "richText")
	if err != nil {
		return err
	}
	err = RegisterWidget(createSeparator, "separator")
	if err != nil {
		return err
	}
	err = RegisterWidget(createSlider, "slider")
	if err != nil {
		return err
	}
	err = RegisterWidget(createSelectEntry, "selectEntry")
	if err != nil {
		return err
	}

	return nil
}

func runForm(formDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	callback := run.BooleanCallback(formDescr[ui.AttrChildren].(ui.CommandsDescr),
		ui.NameSubmit, ui.NameCancel, fullName, win, uiDescr)
	group, _ := formDescr[ui.AttrName].(string) // default value
	if g, ok := formDescr[ui.AttrGroup].(string); ok {
		group = g
	}

	values := make(map[string]any)
	children := formDescr[ui.AttrChildren].(ui.CommandsDescr)
	items := make([]*widget.FormItem, 0, children.Len()-2)
	for _, child := range children.All() {
		if child[ui.AttrKeyword] != KeywordFormItem {
			continue
		}
		name, _ := child[ui.AttrName].(string)
		outputKey := name // default value
		if o, ok := child[ui.AttrOutputKey].(string); ok {
			outputKey = o
		}
		typ := child[ui.AttrType].(string)
		var wdgt fyne.CanvasObject
		if creator, ok := widgetMap[typ]; !ok {
			log.Printf("ERROR: for %q: widget of type %q isn't registered", fullName, typ)
			continue
		} else {
			wdgt = creator(child, values, outputKey, ui.FullNameFor(fullName, name))
		}
		if disabled, _ := child["disabled"].(bool); disabled {
			if w, ok := wdgt.(fyne.Disableable); ok {
				w.Disable()
			}
		}
		label, _ := child["label"].(string)
		hint, _ := child["hint"].(string)
		items = append(items, &widget.FormItem{
			Text:     label,
			Widget:   wdgt,
			HintText: hint,
		})
	}
	form := widget.NewForm(items...)
	if submitText, ok := formDescr["submitText"].(string); ok {
		form.SubmitText = submitText
	}
	if cancelText, ok := formDescr["cancelText"].(string); ok {
		form.CancelText = cancelText
	}
	form.OnSubmit = func() {
		// store all values:
		for k, v := range values {
			// we have to convert any pointer, not simply pointer to any
			a := reflect.ValueOf(v).Elem().Interface()
			ui.StoreValue(a, k, "", fullName, group)
		}
		callback(true)
	}
	form.OnCancel = func() {
		callback(false)
	}
	win.SetContent(form)
}

func RegisterWidget(w Creator, typ string) error {
	if _, ok := widgetMap[typ]; ok {
		return fmt.Errorf("widget of type %q exists already", typ)
	}
	widgetMap[typ] = w
	return nil
}

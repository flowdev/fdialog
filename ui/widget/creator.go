package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flowdev/fdialog/ui"
	"log"
	"net/url"
	"strings"
)

func createEntry(attrs ui.AttributesDescr, values map[string]any, outputKey, fullName string) fyne.CanvasObject {
	entry := widget.NewEntry()
	if ph, _ := attrs["placeHolder"].(string); ph != "" {
		entry.SetPlaceHolder(ph)
	}
	entry.Validator = StringValidator(attrs, fullName)
	values[outputKey] = &entry.Text
	return entry
}

func createMultiLineEntry(attrs ui.AttributesDescr, values map[string]any, outputKey, fullName string) fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	if ph, _ := attrs["placeHolder"].(string); ph != "" {
		entry.SetPlaceHolder(ph)
	}
	entry.Validator = StringValidator(attrs, fullName)
	values[outputKey] = &entry.Text
	return entry
}

func createPasswordEntry(attrs ui.AttributesDescr, values map[string]any, outputKey, fullName string) fyne.CanvasObject {
	entry := widget.NewPasswordEntry()
	if ph, _ := attrs["placeHolder"].(string); ph != "" {
		entry.SetPlaceHolder(ph)
	}
	entry.Validator = StringValidator(attrs, fullName)
	values[outputKey] = &entry.Text
	return entry
}

func createCheckBox(attrs ui.AttributesDescr, values map[string]any, outputKey, _ string) fyne.CanvasObject {
	subLabel, _ := attrs["subLabel"].(string)
	box := widget.NewCheck(subLabel, nil)
	values[outputKey] = &box.Checked
	return box
}

func createCheckGroup(attrs ui.AttributesDescr, values map[string]any, outputKey, _ string) fyne.CanvasObject {
	cg := widget.NewCheckGroup(ui.AnysToStrings(attrs["options"]), func(selected []string) {
		result := strings.Join(selected, "\n")
		values[outputKey] = &result
	})
	cg.SetSelected(ui.AnysToStrings(attrs["initiallySelected"]))
	values[outputKey] = &cg.Selected
	return cg
}

func createHyperlink(attrs ui.AttributesDescr, _ map[string]any, _, fullName string) fyne.CanvasObject {
	text, _ := attrs["text"].(string)
	surl, _ := attrs["url"].(string)
	link, err := url.Parse(surl)
	if err != nil {
		log.Printf("ERROR: for %q: %v", fullName, err)
	}
	return widget.NewHyperlink(text, link)
}

func createRadioGroup(attrs ui.AttributesDescr, values map[string]any, outputKey, _ string) fyne.CanvasObject {
	fmt.Println("createRadioGroup: called")
	rg := widget.NewRadioGroup(ui.AnysToStrings(attrs["options"]), nil)
	horizontal, _ := attrs["horizontal"].(bool)
	rg.Horizontal = horizontal
	required, _ := attrs["required"].(bool)
	rg.Required = required
	if initial, ok := attrs["initiallySelected"].(string); ok {
		rg.SetSelected(initial)
	}
	values[outputKey] = &rg.Selected
	return rg
}

func createRichText(attrs ui.AttributesDescr, _ map[string]any, _, _ string) fyne.CanvasObject {
	text, _ := attrs["text"].(string)
	rt := widget.NewRichTextFromMarkdown(text)
	rt.Wrapping = fyne.TextWrapWord
	scroll, _ := attrs["scroll"].(string)
	switch scroll {
	case "both":
		rt.Scroll = container.ScrollBoth
	case "horizontal":
		rt.Scroll = container.ScrollHorizontalOnly
	case "vertical":
		rt.Scroll = container.ScrollVerticalOnly
	case "none":
		rt.Scroll = container.ScrollNone
	}
	return rt
}

func createSelect(attrs ui.AttributesDescr, values map[string]any, outputKey, _ string) fyne.CanvasObject {
	sel := widget.NewSelect(ui.AnysToStrings(attrs["options"]), nil)
	if ph, ok := attrs["placeHolder"].(string); ok {
		sel.PlaceHolder = ph
	}
	if initial, ok := attrs["initiallySelected"].(string); ok {
		sel.SetSelected(initial)
	}
	values[outputKey] = &sel.Selected
	return sel
}

func createSelectEntry(attrs ui.AttributesDescr, values map[string]any, outputKey, fullName string) fyne.CanvasObject {
	sel := widget.NewSelectEntry(ui.AnysToStrings(attrs["options"]))
	if ph, _ := attrs["placeHolder"].(string); ph != "" {
		sel.SetPlaceHolder(ph)
	}
	sel.Validator = StringValidator(attrs, fullName)
	values[outputKey] = &sel.Text
	return sel
}

func createSeparator(_ ui.AttributesDescr, _ map[string]any, _, _ string) fyne.CanvasObject {
	return widget.NewSeparator()
}

func createSlider(attrs ui.AttributesDescr, values map[string]any, outputKey, _ string) fyne.CanvasObject {
	minv, _ := attrs["min"].(float64) // default min is 0
	maxv, ok := attrs["max"].(float64)
	if !ok {
		maxv = 100.0 // default max is 100
	}
	slider := widget.NewSlider(minv, maxv)
	if initial, ok := attrs["initialValue"].(float64); ok {
		slider.SetValue(initial)
	}
	if step, ok := attrs["step"].(float64); ok { // default step is 1
		slider.Step = step
	}
	values[outputKey] = &slider.Value
	return slider
}

package run

import (
	"fyne.io/fyne/v2"
	"github.com/valyala/fastjson"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/flowdev/fdialog/ui"
)

var jsonArena = &fastjson.ArenaPool{}

// UIDescription runs a whole UI description and returns any error encountered.
func UIDescription(uiDescr ui.CommandsDescr) {
	mainWin, ok := uiDescr.Get(ui.WinMain)
	if !ok {
		log.Printf("FATAL: unable to find main window in UI description")
		return
	}
	if mainWin[ui.AttrKeyword] != ui.KeywordWindow {
		log.Printf(`command with name 'main' is not a window but a:  %q`, mainWin[ui.AttrKeyword])
	}
	appID := "org.flowdev.fdialog"
	if aid, ok := mainWin["appId"]; ok {
		appID = aid.(string)
	}
	log.Printf("INFO: Creating app with ID %q", appID)
	ui.NewApp(appID)

	win, ok := ui.RunFuncForKeyword(ui.KeywordWindow)
	if !ok {
		log.Printf(`unable to get run function for keyword 'window'`)
	}

	win(mainWin, ui.WinMain, nil, uiDescr)
	ui.RunApp()
}

// ---------------------------------------------------------------------------
// Keywords
//

// Window runs a Window description including all of its children.
// In the case of the main window it will run the whole UI.
// The fyne.Window parameter isn't currently used but might be used in the future for a parent window.
func Window(winDescr ui.AttributesDescr, fullName string, _ fyne.Window, uiDescr ui.CommandsDescr) {
	title := ""
	if atitle, ok := winDescr["title"]; ok {
		title = atitle.(string)
	}
	win := ui.NewWindow(title)

	width, height := GetSize(winDescr)
	var winSize fyne.Size
	if width > 0 && height > 0 {
		// TODO: after the real fix with Fyne 2.6 we can use the real values
		winSize = fyne.NewSize(width+1.0, height+1.0)
		win.Resize(winSize)
		win.SetFixedSize(true)
	}

	if fullName == ui.WinMain {
		if code, ok := winDescr["exitCode"]; ok { // set the correct exit code
			ui.StoreExitCode(int32(code.(int64)))
		}

		// Exit the app nicely with the correct exit code ...
		interceptor := func() {
			win.Close()
			ui.ExitApp(-1)
		}
		win.SetOnClosed(interceptor) // ... when the main window is closed or

		// ... when a terminating signal is received.
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			// Block until a signal is received.
			_ = <-signalChan
			interceptor()
		}()
	}

	if children, ok := winDescr[ui.AttrChildren]; ok {
		Children(children, fullName, win, uiDescr)
	}

	win.SetTitle(title)
	win.Show()
	if width > 0 && height > 0 { // TODO: after the real fix with Fyne 2.6 we can remove this workaround
		winSize = fyne.NewSize(width, height)
		win.Resize(winSize)
	}
}

func Children(achildren any, parent string, win fyne.Window, uiDescr ui.CommandsDescr) {
	childDescr := achildren.(ui.CommandsDescr) // type validation has happened already :)

	for name, keywordDescr := range childDescr.All() {
		Keyword(keywordDescr, ui.FullNameFor(parent, name), win, uiDescr)
	}
}

func Keyword(keywordDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	keyword := keywordDescr[ui.AttrKeyword]
	keywordFunc, ok := ui.RunFuncForKeyword(keyword.(string))
	if !ok {
		log.Printf(`ERROR: for %q: unknown keyword %q`, fullName, keyword)
		return
	}
	keywordFunc(keywordDescr, fullName, win, uiDescr)
}

func Link(linkDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	dest := linkDescr["destination"].(string) // has been validated already :)
	dnames := ui.SplitName(dest)
	ok := true

	n := len(dnames)
	tree := uiDescr // start at the top
	var attrs ui.AttributesDescr
	for i := 0; i < n-1; i++ {
		attrs, ok = tree.Get(dnames[i])
		if !ok {
			log.Printf("ERROR: for %q: link destination %q not found",
				fullName, strings.Join(dnames[:i+1], "."))
			return
		}
		dchildren, ok := attrs[ui.AttrChildren]
		if !ok || dchildren == nil {
			log.Printf("ERROR: for %q: no children found for link destination %q",
				fullName, strings.Join(dnames[:i+1], "."))
			return
		}
		tree = dchildren.(ui.CommandsDescr)
	}
	attrs, ok = tree.Get(dnames[n-1]) // the last name always exists or the link wouldn't be valid
	if !ok {
		log.Printf("ERROR: for %q: link destination %q not found", fullName, dest)
		return
	}

	Keyword(attrs, dest, win, uiDescr)
}

func Action(actionDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	action := actionDescr[ui.AttrType]
	runFunc, ok := ui.ActionRunFunc(action.(string))
	if !ok {
		log.Printf(`ERROR: for %q: unknown action %q`, fullName, action)
		return
	}
	runFunc(actionDescr, fullName, win, uiDescr)
}

// ---------------------------------------------------------------------------
// Actions
//

func Exit(exitDescr ui.AttributesDescr, _ string, _ fyne.Window, _ ui.CommandsDescr) {
	code := -1 // intentional default
	if exitDescr["code"] != nil {
		code = int(exitDescr["code"].(int64))
	}
	ui.ExitApp(code)
}

func Close(_ ui.AttributesDescr, _ string, win fyne.Window, _ ui.CommandsDescr) {
	win.Close()
}

func Group(groupDescr ui.AttributesDescr, parent string, win fyne.Window, uiDescr ui.CommandsDescr) {
	childrenDescr := groupDescr[ui.AttrChildren].(ui.CommandsDescr)
	for name, attrs := range childrenDescr.All() {
		if keyword := attrs[ui.AttrKeyword]; keyword.(string) != ui.KeywordAction {
			log.Printf(`ERROR: for %q: only actions allowed, got: %q`, ui.FullNameFor(parent, name), keyword)
			continue
		}
		Keyword(attrs, name, win, uiDescr)
	}
}

func Write(writeDescr ui.AttributesDescr, fullName string, _ fyne.Window, _ ui.CommandsDescr) {
	group, gok := writeDescr[ui.AttrGroup].(string)
	id, iok := writeDescr[ui.AttrID].(string)
	name, nok := writeDescr["fullName"].(string)
	okey, kok := writeDescr[ui.AttrOutputKey].(string)

	switch {
	case nok && kok: // read fullName and write outputKey
		v, ok := ui.GetValueByFullName(name, group)
		if !ok {
			log.Printf(`WARNING: for %q: no value found in group %q with full name %q for writing`,
				fullName, group, name)
		}
		writeMap(map[string]any{okey: v}, fullName)
	case iok && kok: // read ID and write outputKey
		v, ok := ui.GetValueByID(id, group)
		if !ok {
			log.Printf(`WARNING: for %q: no value found in group %q with ID %q for writing`,
				fullName, group, id)
		}
		writeMap(map[string]any{okey: v}, fullName)
	case gok: // write whole group
		m, ok := ui.GetValueGroup(group)
		if !ok {
			log.Printf(`WARNING: for %q: no value found in group %q for writing`, fullName, group)
		}
		writeMap(m, fullName)
	default:
		log.Printf(`WARNING: for %q: no value found for writing`, fullName)
	}
}
func writeMap(m map[string]any, fullName string) {
	m = normalizeMap(m, fullName)
	arena := jsonArena.Get()
	defer func() {
		arena.Reset()
		jsonArena.Put(arena)
	}()
	val := writeJSONMap(m, arena, fullName)
	_, err := os.Stdout.Write(val.MarshalTo(nil))
	if err != nil {
		log.Printf(`ERROR: for %q: unable to write to STDOUT: %v`, fullName, err)
	}
	_, _ = os.Stdout.WriteString("\n")
}
func normalizeMap(m map[string]any, fullName string) map[string]any {
	m2 := make(map[string]any, len(m))
	for k, v := range m {
		if m3, ok := v.(map[string]any); ok {
			v = normalizeMap(m3, fullName)
		}
		keys := ui.SplitName(k)
		if len(keys) <= 1 {
			m2[k] = v
			continue
		}
		m4 := m2
		key := ""
		ok2 := false
		for i := 0; i < len(keys)-1; i++ {
			key = keys[i]
			if v2, ok := m4[key]; ok {
				if m4, ok2 = v2.(map[string]any); !ok2 {
					log.Printf(`ERROR: for %q: key %q is clashing with key %q in map %#v`, fullName, k, key, m)
					m4 = m2
					key = k
					break
				}
			} else {
				tmp := make(map[string]any)
				m4[key] = tmp
				m4 = tmp
			}
		}
		key = keys[len(keys)-1]
		m4[key] = v
	}
	return m2
}
func writeJSONMap(m map[string]any, arena *fastjson.Arena, fullName string) *fastjson.Value {
	obj := arena.NewObject()
	for k, v := range m {
		obj.Set(k, writeJSONValue(v, arena, fullName))
	}
	return obj
}
func writeJSONValue(value any, arena *fastjson.Arena, fullName string) *fastjson.Value {
	switch v := value.(type) {
	case string:
		return arena.NewString(v)
	case int8:
		return arena.NewNumberInt(int(v))
	case int16:
		return arena.NewNumberInt(int(v))
	case int32:
		return arena.NewNumberInt(int(v))
	case int:
		return arena.NewNumberInt(v)
	case int64:
		return arena.NewNumberFloat64(float64(v))
	case uint8:
		return arena.NewNumberInt(int(v))
	case uint16:
		return arena.NewNumberInt(int(v))
	case uint32:
		return arena.NewNumberFloat64(float64(v))
	case uint:
		return arena.NewNumberFloat64(float64(v))
	case uint64:
		return arena.NewNumberFloat64(float64(v))
	case float32:
		return arena.NewNumberFloat64(float64(v))
	case float64:
		return arena.NewNumberFloat64(v)
	case bool:
		if v {
			return arena.NewTrue()
		}
		return arena.NewFalse()
	case []any:
		return writeJSONArray(v, arena, fullName)
	case map[string]any:
		return writeJSONMap(v, arena, fullName)
	default:
		log.Printf(`ERROR: for %q: unable to write unknown data type %T`, fullName, v)
		return arena.NewNull()
	}
}
func writeJSONArray(value []any, arena *fastjson.Arena, fullName string) *fastjson.Value {
	arr := arena.NewArray()
	for i, v := range value {
		arr.SetArrayItem(i, writeJSONValue(v, arena, fullName))
	}
	return arr
}

// ---------------------------------------------------------------------------
// Callbacks

func BooleanCallback(
	childrenDescr ui.CommandsDescr,
	submitName, cancelName string,
	fullName string,
	win fyne.Window,
	uiDescr ui.CommandsDescr,
) func(bool) {
	defaultCallback := func(_ bool) {
		return
	}

	keySubmit, _ := childrenDescr.Get(submitName)
	if keySubmit == nil {
		log.Printf("ERROR: for %q: %q keyword is missing", fullName, submitName)
		return defaultCallback
	}
	keyCancel, _ := childrenDescr.Get(cancelName)
	if keyCancel == nil {
		log.Printf("ERROR: for %q: %q keyword is missing", fullName, cancelName)
		return defaultCallback
	}

	return func(submitted bool) {
		if submitted {
			Keyword(keySubmit, ui.FullNameFor(fullName, submitName), win, uiDescr)
		} else {
			Keyword(keyCancel, ui.FullNameFor(fullName, cancelName), win, uiDescr)
		}
	}
}

func CloseCallback(childDescr ui.CommandsDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) func() {
	defaultCallback := func() {
		return
	}
	keyClose, _ := childDescr.Get("close")
	if keyClose == nil { // action is optional
		return defaultCallback
	}

	return func() {
		Keyword(keyClose, ui.FullNameFor(fullName, "close"), win, uiDescr)
	}
}

// ---------------------------------------------------------------------------
// Helpers
//

func GetSize(descr ui.AttributesDescr) (width, height float32) {
	width = float32(0)
	height = float32(0)
	if v, ok := descr["width"]; ok {
		width = float32(v.(float64))
	}
	if v, ok := descr["height"]; ok {
		height = float32(v.(float64))
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // default is wider than high
	}
	if width <= 0 && height > 0 {
		width = height * 2 // default is wider than high
	}
	return width, height
}

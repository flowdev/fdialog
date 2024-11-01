package run

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"

	"github.com/flowdev/fdialog/ui"
)

// UIDescription runs a whole UI description and returns any error encountered.
func UIDescription(uiDescr ui.CommandsDescr) {
	mainWin := uiDescr[ui.WinMain]
	if mainWin == nil {
		log.Printf("unable to find main window in UI description")
	}
	if mainWin[ui.KeyKeyword] != ui.KeywordWindow {
		log.Printf(`command with name 'main' is not a window but a:  %q`, mainWin[ui.KeyKeyword])
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

	if children, ok := winDescr[ui.KeyChildren]; ok {
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

	for name, keywordDescr := range childDescr {
		Keyword(keywordDescr, ui.FullNameFor(parent, name), win, uiDescr)
	}
}

func Keyword(keywordDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	keyword := keywordDescr[ui.KeyKeyword]
	keywordFunc, ok := ui.RunFuncForKeyword(keyword.(string))
	if !ok {
		log.Printf(`ERROR: for %q: unknown keyword %q`, fullName, keyword)
		return
	}
	keywordFunc(keywordDescr, fullName, win, uiDescr)
}

func Link(linkDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	dest := linkDescr["destination"].(string) // has been validated already :)
	dnames := strings.Split(dest, ".")

	n := len(dnames)
	tree := uiDescr // start at the top
	for i := 0; i < n-1; i++ {
		dchildren := tree[dnames[i]][ui.KeyChildren]
		if dchildren == nil {
			log.Printf("ERROR: for %q: no children found for link destination %q",
				fullName, strings.Join(dnames[:i+1], "."))
			return
		}
		tree = dchildren.(ui.CommandsDescr)
	}
	dkwMap := tree[dnames[n-1]] // the last name always exists or the link wouldn't be valid

	Keyword(dkwMap, dest, win, uiDescr)
}

func Action(actionDescr ui.AttributesDescr, fullName string, win fyne.Window, uiDescr ui.CommandsDescr) {
	action := actionDescr[ui.KeyType]
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

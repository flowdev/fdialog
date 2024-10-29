package run

import (
	"fmt"
	"fyne.io/fyne/v2"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/flowdev/fdialog/ui"
)

// UIDescription runs a whole UI description and returns any error encountered.
func UIDescription(uiDescr map[string]map[string]any) error {
	mainWin := uiDescr[ui.WinMain]
	if mainWin == nil {
		return fmt.Errorf("unable to find main window in UI description")
	}
	if mainWin[ui.KeyKeyword] != ui.KeywordWindow {
		return fmt.Errorf(`command with name 'main' is not a window but a:  %q`, mainWin[ui.KeyKeyword])
	}
	if err := ui.NewApp("org.flowdev.fdialog"); err != nil {
		return err
	}

	win, ok := ui.KeywordRunFunc(ui.KeywordWindow)
	if !ok {
		return fmt.Errorf(`unable to get run function for keyword 'window'`)
	}

	err := win(mainWin, []string{ui.WinMain}, nil, uiDescr)
	if err != nil {
		return err
	}
	ui.RunApp()
	return nil
}

// Window runs a Window description including all of its children.
// In the case of the main window it will run the whole UI.
// The fyne.Window parameter isn't currently used but might be used in the future for a parent window.
func Window(winDescr map[string]any, fullName []string, _ fyne.Window, uiDescr map[string]map[string]any) error {
	title := ""
	if _, ok := winDescr["title"]; ok {
		title = winDescr["title"].(string)
	}
	win := ui.NewWindow(title)

	width := float64(0)
	height := float64(0)
	if _, ok := winDescr["width"]; ok {
		width = winDescr["width"].(float64)
	}
	if _, ok := winDescr["height"]; ok {
		height = winDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide windows look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide windows look good
	}
	var winSize fyne.Size
	if width > 0 && height > 0 {
		// TODO: after the real fix with Fyne 2.6 we can use the real values
		winSize = fyne.NewSize(float32(width+1.0), float32(height+1.0))
		win.Resize(winSize)
		win.SetFixedSize(true)
	}

	if _, ok := winDescr[ui.KeyChildren]; ok {
		err := Children(winDescr[ui.KeyChildren], fullName, win, uiDescr)
		if err != nil {
			return err
		}
	}

	if ui.SameFullName(fullName, "main") {
		// Exit the app nicely with the correct exit code ...
		interceptor := func() {
			ui.ExitApp(-1)
		}
		win.SetCloseIntercept(interceptor) // ... when the main window is closed or

		// ... when a terminating signal is received.
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			// Block until a signal is received.
			_ = <-signalChan
			interceptor()
		}()
	}

	win.SetTitle(title)
	win.Show()
	if width > 0 && height > 0 { // TODO: after the real fix with Fyne 2.6 we can remove this workaround
		winSize = fyne.NewSize(float32(width), float32(height))
		win.Resize(winSize)
	}
	return nil
}

func Children(achildren any, parent []string, win fyne.Window, uiDescr map[string]map[string]any) error {
	childDescr := achildren.(map[string]map[string]any) // type validation has happened already :)

	for name, keywordDescr := range childDescr {
		fullName := append(parent, name)
		err := Keyword(keywordDescr, fullName, win, uiDescr)
		if err != nil {
			return err
		}
	}
	return nil
}

func Keyword(keywordDescr map[string]any, fullName []string, win fyne.Window, uiDescr map[string]map[string]any) error {
	keyword := keywordDescr[ui.KeyKeyword]
	keywordFunc, ok := ui.KeywordRunFunc(keyword.(string))
	if !ok {
		return fmt.Errorf(`for %q: unknown keyword %q`, ui.DisplayName(fullName), keyword)
	}
	return keywordFunc(keywordDescr, fullName, win, uiDescr)
}

func Link(linkDescr map[string]any, fullName []string, win fyne.Window, uiDescr map[string]map[string]any) error {
	dest := linkDescr["destination"].(string) // has been validated already :)
	dnames := strings.Split(dest, ".")

	n := len(dnames)
	tree := uiDescr // start at the top
	for i := 0; i < n-1; i++ {
		dchildren := tree[dnames[i]][ui.KeyChildren]
		if dchildren == nil {
			return fmt.Errorf("for %q: no children found for link destination %q",
				fullName, strings.Join(dnames[:i+1], "."))
		}
		tree = dchildren.(map[string]map[string]any)
	}
	dkwMap := tree[dnames[n-1]] // the last name always exists or the link wouldn't be valid

	return Keyword(dkwMap, ui.FullName(dest), win, uiDescr)
}

func Action(actionDescr map[string]any, fullName []string, win fyne.Window, uiDescr map[string]map[string]any) error {
	_ = uiDescr // currently not used but might change with more actions
	action := actionDescr[ui.KeyType]
	runFunc, ok := ui.ActionRunFunc(action.(string))
	if !ok {
		return fmt.Errorf(`for %q: unknown action %q`, ui.DisplayName(fullName), action)
	}
	return runFunc(actionDescr, fullName, win, uiDescr)
}

func Exit(exitDescr map[string]any, fullName []string, _ fyne.Window, _ map[string]map[string]any) error {
	code := 0 // intentional default
	if exitDescr["code"] != nil {
		code = int(exitDescr["code"].(int64))
	}
	log.Printf("INFO: exiting app as requested at position %q with code: %d", ui.DisplayName(fullName), code)
	os.Exit(code)
	return nil // just for the compiler :)
}

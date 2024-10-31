package run

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"

	"github.com/flowdev/fdialog/ui"
)

// UIDescription runs a whole UI description and returns any error encountered.
func UIDescription(uiDescr ui.CommandsDescr) error {
	mainWin := uiDescr[ui.WinMain]
	if mainWin == nil {
		return fmt.Errorf("unable to find main window in UI description")
	}
	if mainWin[ui.KeyKeyword] != ui.KeywordWindow {
		return fmt.Errorf(`command with name 'main' is not a window but a:  %q`, mainWin[ui.KeyKeyword])
	}
	appID := "org.flowdev.fdialog"
	if aid, ok := mainWin["appId"]; ok {
		appID = aid.(string)
	}
	log.Printf("INFO: Creating app with ID %q", appID)
	if err := ui.NewApp(appID); err != nil {
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
func Window(winDescr ui.AttributesDescr, fullName []string, _ fyne.Window, uiDescr ui.CommandsDescr) error {
	title := ""
	if _, ok := winDescr["title"]; ok {
		title = winDescr["title"].(string)
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
	win.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			win.Close()
		}
	})

	if ui.FullNameIs(fullName, "main") {
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
		if err := Children(children, fullName, win, uiDescr); err != nil {
			return err
		}
	}

	win.SetTitle(title)
	win.Show()
	if width > 0 && height > 0 { // TODO: after the real fix with Fyne 2.6 we can remove this workaround
		winSize = fyne.NewSize(width, height)
		win.Resize(winSize)
	}
	return nil
}

func Children(achildren any, parent []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	childDescr := achildren.(ui.CommandsDescr) // type validation has happened already :)

	for name, keywordDescr := range childDescr {
		fullName := append(parent, name)
		err := Keyword(keywordDescr, fullName, win, uiDescr)
		if err != nil {
			return err
		}
	}
	return nil
}

func Keyword(keywordDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	keyword := keywordDescr[ui.KeyKeyword]
	keywordFunc, ok := ui.KeywordRunFunc(keyword.(string))
	if !ok {
		return fmt.Errorf(`for %q: unknown keyword %q`, ui.DisplayName(fullName), keyword)
	}
	return keywordFunc(keywordDescr, fullName, win, uiDescr)
}

func Link(linkDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
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
		tree = dchildren.(ui.CommandsDescr)
	}
	dkwMap := tree[dnames[n-1]] // the last name always exists or the link wouldn't be valid

	return Keyword(dkwMap, ui.FullName(dest), win, uiDescr)
}

func Action(actionDescr ui.AttributesDescr, fullName []string, win fyne.Window, uiDescr ui.CommandsDescr) error {
	_ = uiDescr // currently not used but might change with more actions
	action := actionDescr[ui.KeyType]
	runFunc, ok := ui.ActionRunFunc(action.(string))
	if !ok {
		return fmt.Errorf(`for %q: unknown action %q`, ui.DisplayName(fullName), action)
	}
	return runFunc(actionDescr, fullName, win, uiDescr)
}

func Exit(exitDescr ui.AttributesDescr, fullName []string, _ fyne.Window, _ ui.CommandsDescr) error {
	code := -1 // intentional default
	if exitDescr["code"] != nil {
		code = int(exitDescr["code"].(int64))
	}
	ui.ExitApp(code)
	return nil // just for the compiler :)
}

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

package run

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"

	"github.com/flowdev/fdialog/parse"
)

const winMain = "main"

var fapp fyne.App // needed for exiting cleanly in actions
var exitCode = new(atomic.Int32)

func UIDescription(uiDescr map[string]map[string]any) error {
	mainWin := uiDescr[winMain]
	if mainWin == nil {
		return fmt.Errorf("unable to find main window in UI description")
	}
	if mainWin[parse.KeyKeyword] != parse.KeywordWindow {
		return fmt.Errorf(`keyword map with name "main" is not a window but a:  %q`, mainWin[parse.KeyKeyword])
	}
	fapp = app.NewWithID("github.com/flowdev/fdialog")

	err := runWindow(mainWin, winMain, nil, uiDescr)
	if err != nil {
		return err
	}
	fapp.Run()
	return nil
}

// runWindow runs a window description including all of its children.
// In the case of the main window it will run the whole UI.
// The fyne.Window parameter isn't currently used but might be used in the future for a parent window.
func runWindow(winDescr map[string]any, fullName string, _ fyne.Window, uiDescr map[string]map[string]any) error {
	title := ""
	if _, ok := winDescr["title"]; ok {
		title = winDescr["title"].(string)
	}
	win := fapp.NewWindow(title)

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

	if _, ok := winDescr[parse.KeyChildren]; ok {
		err := runChildren(winDescr[parse.KeyChildren], fullName, win, uiDescr)
		if err != nil {
			return err
		}
	}

	if fullName == "main" {
		// Exit the app nicely with the correct exit code ...
		interceptor := func() {
			fapp.Quit()
			code := int(exitCode.Load())
			log.Printf("INFO: exiting app as requested from main window with code: %d", code)
			os.Exit(code)
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

func runChildren(achildren any, parent string, win fyne.Window, uiDescr map[string]map[string]any) error {
	childDescr := achildren.(map[string]map[string]any) // type validation has happened already :)

	for name, keywordDescr := range childDescr {
		fullName := parse.JoinParentName(parent, name)
		err := runKeyword(keywordDescr, fullName, win, uiDescr)
		if err != nil {
			return err
		}
	}
	return nil
}

func runKeyword(keywordDescr map[string]any, fullName string, win fyne.Window, uiDescr map[string]map[string]any) error {
	var err error
	keyword := keywordDescr[parse.KeyKeyword]

	switch keyword {
	case parse.KeywordWindow:
		err = runWindow(keywordDescr, fullName, win, uiDescr)
	case parse.KeywordDialog:
		err = runDialog(keywordDescr, fullName, win, uiDescr)
	case parse.KeywordAction:
		err = runAction(keywordDescr, fullName, uiDescr)
	case parse.KeywordLink:
		err = runLink(keywordDescr, fullName, win, uiDescr)
	default:
		err = fmt.Errorf(`for %q: unknown keyword %q`, fullName, keyword)
	}
	return err
}

func runLink(linkDescr map[string]any, fullName string, win fyne.Window, uiDescr map[string]map[string]any) error {
	dest := linkDescr["destination"].(string) // has been validated already :)
	dnames := strings.Split(dest, ".")

	n := len(dnames)
	tree := uiDescr // start at the top
	for i := 0; i < n-1; i++ {
		dchildren := tree[dnames[i]][parse.KeyChildren]
		if dchildren == nil {
			return fmt.Errorf("for %q: no children found for link destination %q",
				fullName, strings.Join(dnames[:i+1], "."))
		}
		tree = tree[dnames[i]][parse.KeyChildren].(map[string]map[string]any)
	}
	dkwMap := tree[dnames[n-1]] // the last name always exists of the link wouldn't be valid

	return runKeyword(dkwMap, dest, win, uiDescr)
}

func runAction(actionDescr map[string]any, fullName string, uiDescr map[string]map[string]any) error {
	_ = uiDescr // currently not used but might change with more actions
	var err error
	action := actionDescr[parse.KeyType]

	switch action {
	case "exit":
		err = runExit(actionDescr, fullName)
	default:
		err = fmt.Errorf(`for %q: unknown action type %q`, fullName, action)
	}
	return err
}

func runExit(exitDescr map[string]any, fullName string) error {
	code := 0 // intentional default
	if exitDescr["code"] != nil {
		code = int(exitDescr["code"].(int64))
	}
	log.Printf("INFO: exiting app as requested at position %q with code: %d", fullName, code)
	os.Exit(code)
	return nil // just for the compiler :)
}

func runDialog(dialogDescr map[string]any, fullName string, win fyne.Window, uiDescr map[string]map[string]any) error {
	var err error
	dlg := dialogDescr[parse.KeyType]

	switch dlg {
	case "info":
		err = runInfo(dialogDescr, fullName, win)
	case "error":
		err = runError(dialogDescr, fullName, win)
	case "confirmation":
		err = runConfirmation(dialogDescr, fullName, win, uiDescr)
	case "openFile":
		err = runOpenFile(dialogDescr, fullName, win, uiDescr)
	default:
		err = fmt.Errorf(`for %q: unknown dialog type %q`, fullName, dlg)
	}
	return err
}

func runOpenFile(
	ofDescr map[string]any,
	fullName string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
) error {
	_, _ = fullName, uiDescr
	ofDialog := dialog.NewFileOpen(func(frd fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if frd == nil {
			fmt.Println("<CLOSED>")
			fapp.Quit()
			os.Exit(1)
		}
		fmt.Println("file to open:", strings.TrimPrefix(frd.URI().String(), "file://"))

		fapp.Quit()
		os.Exit(0)
	}, win)

	ofDialog.SetOnClosed(func() {
		fmt.Println("dialog closed, exiting?")
	})

	extAttr := ofDescr["extensions"]
	if extAttr != nil {
		extSlice := strings.Split(extAttr.(string), ",")
		for i := 0; i < len(extSlice); i++ {
			extSlice[i] = strings.TrimSpace(extSlice[i])
		}
		ofDialog.SetFilter(storage.NewExtensionFileFilter(extSlice))
	}

	value := ofDescr["confirmText"]
	if value != nil {
		ofDialog.SetConfirmText(value.(string))
	}
	value = ofDescr["dismissText"]
	if value != nil {
		ofDialog.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := ofDescr["width"]; ok {
		width = ofDescr["width"].(float64)
	}
	if _, ok := ofDescr["height"]; ok {
		height = ofDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		ofSize := fyne.NewSize(float32(width), float32(height))
		ofDialog.Resize(ofSize)
	}

	ofDialog.Show()
	return nil
}

func runConfirmation(
	cnfDescr map[string]any,
	fullName string,
	win fyne.Window,
	uiDescr map[string]map[string]any,
) error {

	callback, err := confirmCallback(cnfDescr[parse.KeyChildren].(map[string]map[string]any), fullName, uiDescr)
	if err != nil {
		return err
	}

	value := cnfDescr["title"]
	title := ""
	if value != nil {
		title = value.(string)
	}
	message := cnfDescr["message"].(string) // message is required
	cnf := dialog.NewConfirm(title, message, callback, win)

	value = cnfDescr["confirmText"]
	if value != nil {
		cnf.SetConfirmText(value.(string))
	}
	value = cnfDescr["dismissText"]
	if value != nil {
		cnf.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := cnfDescr["width"]; ok {
		width = cnfDescr["width"].(float64)
	}
	if _, ok := cnfDescr["height"]; ok {
		height = cnfDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		cnfSize := fyne.NewSize(float32(width), float32(height))
		cnf.Resize(cnfSize)
	}
	escapeKey := &desktop.CustomShortcut{KeyName: fyne.KeyEscape}
	win.Canvas().AddShortcut(escapeKey, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Escape")
	})

	exitCode.Store(1) // closing the window => dismissed
	cnf.Show()
	return nil
}

func confirmCallback(
	childrenDescr map[string]map[string]any,
	fullName string,
	uiDescr map[string]map[string]any,
) (func(bool), error) {

	actConfirm := childrenDescr["confirm"]
	if actConfirm == nil {
		return nil, fmt.Errorf("for %q: confirm action is missing", fullName)
	}
	keyword := actConfirm[parse.KeyKeyword].(string)
	if keyword != parse.KeywordAction {
		return nil, fmt.Errorf("for %q: confirm action is not an action but a %q", fullName, keyword)
	}

	actDismiss := childrenDescr["dismiss"]
	if actDismiss == nil {
		return nil, fmt.Errorf("for %q: dismiss action is missing", fullName)
	}
	keyword = actDismiss[parse.KeyKeyword].(string)
	if keyword != parse.KeywordAction {
		return nil, fmt.Errorf("for %q: dismiss action is not an action but a %q", fullName, keyword)
	}
	return func(confirmed bool) {
		if confirmed {
			err := runAction(actConfirm, parse.JoinParentName(fullName, "confirm"), uiDescr)
			if err != nil {
				log.Printf("ERROR: Can't run confirm action: %v", err)
			}
		} else {
			err := runAction(actDismiss, parse.JoinParentName(fullName, "dismiss"), uiDescr)
			if err != nil {
				log.Printf("ERROR: Can't run dismiss action: %v", err)
			}
		}
	}, nil
}

func runError(errorDescr map[string]any, fullName string, win fyne.Window) error {
	_ = fullName                              // currently not used but might change
	message := errorDescr["message"].(string) // message is required
	errorDialog := dialog.NewError(errors.New(message), win)
	errorDialog.SetOnClosed(func() {
		os.Exit(0) // error has been noted
	})

	value := errorDescr["buttonText"]
	if value != nil {
		errorDialog.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := errorDescr["width"]; ok {
		width = errorDescr["width"].(float64)
	}
	if _, ok := errorDescr["height"]; ok {
		height = errorDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		infoSize := fyne.NewSize(float32(width), float32(height))
		errorDialog.Resize(infoSize)
	}

	exitCode.Store(0) // error has been noted; so all is OK
	errorDialog.Show()
	return nil
}

func runInfo(infoDescr map[string]any, fullName string, win fyne.Window) error {
	_ = fullName // currently not used but might change
	value := infoDescr["title"]
	title := ""
	if value != nil {
		title = value.(string)
	}
	message := infoDescr["message"].(string) // message is required
	info := dialog.NewInformation(title, message, win)
	info.SetOnClosed(func() {
		os.Exit(0) // info has been noted
	})

	value = infoDescr["buttonText"]
	if value != nil {
		info.SetDismissText(value.(string))
	}

	width := float64(0)
	height := float64(0)
	if _, ok := infoDescr["width"]; ok {
		width = infoDescr["width"].(float64)
	}
	if _, ok := infoDescr["height"]; ok {
		height = infoDescr["height"].(float64)
	}
	if width > 0 && height <= 0 {
		height = width * 0.5 // wide dialogs look good
	}
	if width <= 0 && height > 0 {
		width = height * 2 // wide dialogs look good
	}
	if width > 0 && height > 0 {
		infoSize := fyne.NewSize(float32(width), float32(height))
		info.Resize(infoSize)
	}

	exitCode.Store(0) // info has been noted; so all is OK
	info.Show()
	return nil
}

package ui

import (
	"log"
	"os"
	"sync/atomic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var fapp fyne.App // needed for exiting cleanly in actions
var exitCode = new(atomic.Int32)

// ---------------------------------------------------------------------------
// Helpers

func NewApp(appid string) {
	if fapp != nil {
		log.Printf("ERROR: app with ID %q is already running", fapp.UniqueID())
		return
	}
	fapp = app.NewWithID(appid)
}

func RunApp() {
	fapp.Run()
}

func ExitApp(code int) {
	fapp.Quit()
	fapp = nil
	if code < 0 {
		code = int(exitCode.Load())
	}
	log.Printf("INFO: exiting app with code: %d", code)
	os.Exit(code)
}

func NewWindow(title string) fyne.Window {
	return fapp.NewWindow(title)
}

// StoreExitCode stores the given code as exit code for ending the app.
func StoreExitCode(code int32) {
	exitCode.Store(code)
}

func AnysToStrings(a any) []string {
	al, ok := a.([]any)
	if !ok {
		return nil
	}
	sl := make([]string, len(al))
	for i := 0; i < len(sl); i++ {
		sl[i] = al[i].(string) // we validated this
	}
	return sl
}

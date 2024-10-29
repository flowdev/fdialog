package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
	"sync/atomic"
)

var fapp fyne.App // needed for exiting cleanly in actions
var exitCode = new(atomic.Int32)

// ---------------------------------------------------------------------------
// Helpers

func NewApp(appid string) error {
	if fapp != nil {
		return fmt.Errorf("app with ID %q is already running", fapp.UniqueID())
	}
	fapp = app.NewWithID(appid)
	return nil
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
	log.Printf("INFO: exiting app as requested with code: %d", code)
	os.Exit(code)
}

func NewWindow(title string) fyne.Window {
	return fapp.NewWindow(title)
}

// StoreExitCode stores the given code as exit code for ending the app.
func StoreExitCode(code int32) {
	exitCode.Store(code)
}

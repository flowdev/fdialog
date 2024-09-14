package cmd

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/flowdev/fdialog/window"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/spf13/cobra"
)

var runCmdData = struct {
	format   string
	fileName string
	url      string
}{}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "display and run a description of a GUI",
	Long: `Display And Run a Description Of a GUI

If no file or URL is given, the UI description is read from standard input.`,
	Args: cobra.NoArgs,
	Run:  doRun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runCmdData.format, "format", "t", "uidl",
		"format of the GUI description (valid values are: json, go or uidl)")
	runCmd.Flags().StringVarP(&runCmdData.fileName, "file", "f", "",
		"name of file with GUI description")
	runCmd.Flags().StringVarP(&runCmdData.url, "url", "u", "",
		"URL where the GUI description can be fetched with HTTP GET")
	runCmd.MarkFlagsMutuallyExclusive("file", "url")
}

func doRun(cmd *cobra.Command, args []string) {
	fmt.Printf("run called with file=%q, url=%q, format=%q and args=%q\n",
		runCmdData.fileName, runCmdData.url, runCmdData.format, args)

	fda := app.New()

	// For info:
	//win := window.NewInformation("Info!", "This is the info for you.", fda)
	//win.SetDismissText("Got it.")
	//win.SetOnClosed(func() {
	//	fda.Quit()
	//	os.Exit(0)
	//})

	// For error:
	//win := window.NewError(errors.New("An error happened!"), fda)
	//win.SetOnClosed(func() {
	//	fda.Quit()
	//	os.Exit(0)
	//})

	// For Confirmation:
	win := window.NewConfirm("Confirmation", "Do you really want to do XXX?", confirmCallback(fda), fda)
	win.SetDismissText("Nah")
	win.SetConfirmText("Oh Yes!")

	win.Resize(fyne.NewSize(200, 100))
	win.SetFixedSize(true)
	win.Show()
	fda.Run()
}

func confirmCallback(fda fyne.App) func(bool) {
	return func(response bool) {
		fda.Quit()
		fmt.Println("Fyne confirmCallback responded with:", response)
		if response {
			os.Exit(0)
		}
		os.Exit(1)
	}
}

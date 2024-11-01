package cobracmd

import (
	"github.com/flowdev/fdialog/valid"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/flowdev/fdialog/parse"
	"github.com/flowdev/fdialog/run"
)

var runCmdData = struct {
	format   string
	fileName string
	url      string
	lenient  bool
}{}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a description for a UI",
	Long: `Run a Description For a User Interface And Display the Resulting GUI

If no file or URL is given, the UI description is read from standard input.`,
	Args: cobra.NoArgs,
	Run:  doRun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runCmdData.fileName, "file", "f", "",
		"name of file with UI description")
	runCmd.Flags().StringVarP(&runCmdData.url, "url", "u", "",
		"URL where the GUI description can be fetched with HTTP GET")
	runCmd.MarkFlagsMutuallyExclusive("file", "url")
	runCmd.Flags().StringVarP(&runCmdData.format, "format", "t", "uidl",
		"format of the UI description (valid values are: 'json' or 'uidl')")
	runCmd.Flags().BoolVarP(&runCmdData.lenient, "lenient", "l", false,
		"if flag is given, additional attributes in the UI description are only warned about")
}

func doRun(_ *cobra.Command, _ []string) {
	var rd io.Reader
	var err error

	if runCmdData.fileName != "" {
		rd, err = os.Open(runCmdData.fileName)
		if err != nil {
			log.Printf("ERROR: Could not open UI description file: %v", err)
			os.Exit(11)
		}
	} else {
		rd = os.Stdin
	}

	uiDescr, err := parse.UIDescription(rd, runCmdData.fileName, runCmdData.format)
	if err != nil {
		log.Printf("ERROR: Unable to parse UI description:\n%v", err)
		os.Exit(12)
	}
	if ok := valid.UIDescription(uiDescr, !runCmdData.lenient); !ok {
		os.Exit(13)
	}
	run.UIDescription(uiDescr)
}

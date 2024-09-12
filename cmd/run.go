package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runFormat string
var fileName string
var url string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "display and run a description of a GUI",
	Long:  `display and run a description of a GUI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runFormat, "format", "t", "uidl",
		"format of the GUI description (valid values are: json, go or uidl)")
	runCmd.Flags().StringVarP(&fileName, "file", "f", "",
		"name of file with GUI description")
	runCmd.Flags().StringVarP(&url, "url", "u", "",
		"URL where the GUI description can be fetched with HTTP GET")
	runCmd.MarkFlagsMutuallyExclusive("file", "url")
}

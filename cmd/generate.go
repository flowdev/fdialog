package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dest string
var genFormat string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go files that create the described GUI",
	Long:  `Generate Go files that create the described GUI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&genFormat, "format", "t", "uidl",
		"format of the GUI description (valid values are: json, go or uidl)")
	generateCmd.Flags().StringVarP(&dest, "dest", "d", ".",
		"destination directory for the generated result files")
}

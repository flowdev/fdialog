package cobracmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmdData = struct {
	dest   string
	format string
}{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go files that create the described GUI",
	Long:  `Generate Go files that create the described GUI`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("generate called with dest=%q, format=%q and args=%q\n",
			generateCmdData.dest, generateCmdData.format, args)
	},
}

func init() {
	//rootCmd.AddCommand(generateCmd)
	//
	//generateCmd.Flags().StringVarP(&generateCmdData.format, "format", "t", "uidl",
	//	"format of the GUI description (valid values are: 'json' and 'uidl')")
	//generateCmd.Flags().StringVarP(&generateCmdData.dest, "dest", "d", ".",
	//	"destination directory for the generated result files")
}

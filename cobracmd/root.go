package cobracmd

import (
	"os"

	cmd "github.com/can3p/kleiner/shared/cmd/cobra"
	"github.com/can3p/kleiner/shared/published"
	"github.com/flowdev/fdialog/generated/buildinfo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fdialog",
	Short: "Create Native GUIs With Ease",
	Long:  "Create Native GUIs With Ease",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	info := buildinfo.Info()

	cmd.Setup(info, rootCmd)
	published.MaybeNotifyAboutNewVersion(info)
}

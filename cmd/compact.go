/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// compactCmd represents the compact command
var compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "Compact (minimise) a UIDL file to a .min.uidl file",
	Long:  `Compact (minimise) a UIDL file to a .min.uidl file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compact called")
	},
}

func init() {
	rootCmd.AddCommand(compactCmd)
}

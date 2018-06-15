package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
	Short: "Returns version in semver format",
	Use:   "version",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {

		// Implement Methods

		fmt.Println("diff called")
	},
	Short: "Display differences between local copy of records and current state in Cloudflare",
	Use:   "diff",
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Define flags
}

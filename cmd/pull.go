package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Long: `Pulls current state of records from Cloudflare and overwrites local copy.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Implement methods here

		fmt.Println("pull called")
	},
	Short: "Pulls current state of records from Cloudflare",
	Use:   "pull",
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Define flags here
}

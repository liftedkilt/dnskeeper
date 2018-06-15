package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {

		// Implement methods here

		fmt.Println("push called")
	},
	Short: "Pushes local copy of records to Cloudflare, updating if needed",
	Use:   "push",
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Define flags here
}

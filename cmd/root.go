package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnskeeper",
	Short: "dnskeeper is a tool for managing DNS records in cloudflare and storing them locally",
	Long: `A tool for managing DNS records in cloudflare and storing them locally.
Complete documentation is available at http://github.com/liftedkilt/dnskeeper`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/liftedkilt/dnskeeper/dnskeeper"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Please specify zone file to validate")
		}
		return nil
	},
	Long: `Usage: dnskeeper validate /path/to/example.org

dnskeeper validate will verify that a zone file is pasable by dnskeeper. Returns quietly with an exit code of 0 on success, 1 on failure.`,
	Run: func(cmd *cobra.Command, args []string) {

		zoneFile := args[0]
		file, err := os.Open(zoneFile)
		if err != nil {
			log.Fatalln(err)
		}

		zoneBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln(err)
		}

		// Discard valid output - we only care if there is an error
		if _, err := dnskeeper.ZoneParse(zoneBytes); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	},
	Short: "checks that a given zone file is parsble by dnskeeper",
	Use:   "validate",
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

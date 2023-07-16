package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const SILEO_VERSION = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print sileo version",
	Long: "Display the current sileo version on this machine",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sileo -- %s\n", SILEO_VERSION)
	},
}

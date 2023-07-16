package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/theprojectafrica/sileo/pkg"
)

// Command line Flags
var entryPath string

func init() {
	rootCmd.Flags().StringVarP(&entryPath, "file", "f", "", "Path to golang file to run")
}

var rootCmd = &cobra.Command{
	Use: "sileo",
	Short: "Sileo - Automatic Golang Project Rebuilder",
	Long: "Sileo is a commandline tool mainly for monitoring and rebuilding golang projects during development",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Test()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

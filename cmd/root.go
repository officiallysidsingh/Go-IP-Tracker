package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iptracker",
	Short: "IP Tracker CLI App",
	Long: `Go-IP-Tracker is a CLI app that tracks the Location of the given IP Address.`,
}

func Execute() error {
	return rootCmd.Execute()
}
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Version string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of aws-tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("awstools-cli", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

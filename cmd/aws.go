package cmd

import (
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS tools for managing AWS resources",
	// Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(lambdaCmd)
}

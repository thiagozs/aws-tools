package cmd

import (
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS tools for managing AWS resources",
}

func init() {
	rootCmd.AddCommand(awsCmd)
	awsCmd.AddCommand(lambdaCmd)
}

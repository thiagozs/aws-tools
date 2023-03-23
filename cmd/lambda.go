package cmd

import (
	"github.com/spf13/cobra"
)

var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "AWS Lambda tools and commands",
}

func init() {
	lambdaCmd.AddCommand(lambdaCalcCmd)
}

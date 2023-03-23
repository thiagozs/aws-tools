package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	usecase "github.com/thiagozs/aws-tools/internal/usecase/lambda"
)

var (
	region   string
	memory   int
	requests int
	exectime int
)

var lambdaCalcCmd = &cobra.Command{
	Use:   "calc",
	Short: "AWS lambda calculator tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lambdaCalc called")

		lambda := usecase.NewUseCaseCalcLambda()
		cost, err := lambda.CalculateCost(region, memory, requests, exectime)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Cost: $%.2f\n", cost)

	},
}

func init() {
	rootCmd.AddCommand(lambdaCalcCmd)

	lambdaCalcCmd.Flags().StringVarP(&region, "region", "r", "us-east-1", "AWS region")
	lambdaCalcCmd.Flags().IntVarP(&memory, "memory", "m", 128, "Memory in MB")
	lambdaCalcCmd.Flags().IntVarP(&requests, "requests", "q", 1000, "Requests per month")
	lambdaCalcCmd.Flags().IntVarP(&exectime, "exectime", "t", 1000, "Execution time in ms")

}

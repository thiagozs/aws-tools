package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	usecase "github.com/thiagozs/aws-tools/internal/usecase/lambda"
	"github.com/thiagozs/aws-tools/pkg/utils"
)

var (
	region   string
	memory   int
	requests int
	exectime int
	format   string
)

var lambdaCalcCmd = &cobra.Command{
	Use:   "calc",
	Short: "AWS lambda calculator tool",
	Run: func(cmd *cobra.Command, args []string) {

		lambda := usecase.NewUseCaseCalcLambda()
		cost, err := lambda.CalculateCost(region, memory, requests, exectime)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := map[string]interface{}{
			"region":     region,
			"memory":     memory,
			"requests":   requests,
			"exectime":   exectime,
			"cost":       cost,
			"cost_human": fmt.Sprintf("%.2f", cost),
		}

		switch format {
		case "json":
			output, err := utils.WriteJSON(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(output)

		case "csv":
			output, err := utils.WriteCSV(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(output)

		default:
			fmt.Println("AWS Lambda Calculator")
			fmt.Println("Region:", region)
			fmt.Println("Memory:", memory)
			fmt.Println("Requests:", requests)
			fmt.Println("Execution time:", exectime)
			fmt.Printf("Cost: $%.2f\n", cost)
		}

	},
}

func init() {
	lambdaCalcCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "AWS region")
	lambdaCalcCmd.PersistentFlags().IntVarP(&memory, "memory", "m", 128, "Memory in MB")
	lambdaCalcCmd.PersistentFlags().IntVarP(&requests, "requests", "q", 1000, "Requests per month")
	lambdaCalcCmd.PersistentFlags().IntVarP(&exectime, "exectime", "t", 1000, "Execution time in ms")
	lambdaCalcCmd.PersistentFlags().StringVarP(&format, "format", "f", "stdout", "Output format (stdout, json, csv)")

	lambdaCalcCmd.PersistentFlags().Lookup("region").NoOptDefVal = "us-east-1"
	lambdaCalcCmd.PersistentFlags().Lookup("memory").NoOptDefVal = "128"
	lambdaCalcCmd.PersistentFlags().Lookup("requests").NoOptDefVal = "10000"
	lambdaCalcCmd.PersistentFlags().Lookup("exectime").NoOptDefVal = "5000"

}

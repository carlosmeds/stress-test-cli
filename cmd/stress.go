package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress test a given URL",
	Long:  `Stress test a given URL with a number of requests and concurrency level.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stress called for url " + cmd.Flag("url").Value.String() + " with " + cmd.Flag("requests").Value.String() + " requests and " + cmd.Flag("concurrency").Value.String() + " concurrency")
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)

	stressCmd.Flags().StringP("url", "u", "", "URL to stress test")
	stressCmd.Flags().IntP("requests", "r", 100, "Number of requests to perform")
	stressCmd.Flags().IntP("concurrency", "c", 1, "Number of multiple requests to make at a time")
	stressCmd.MarkFlagRequired("url")
}

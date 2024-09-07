package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress test a given URL",
	Long: `Stress test a given URL with a number of requests and concurrency level.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stress called")
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

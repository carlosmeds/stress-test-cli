package cmd

import (
	"github.com/carlosmeds/stress-test-cli/internal/usecase"
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stress",
		Short: "Stress test a given URL ",
		Long:  `Stress test a given URL with a number of requests and concurrency level.`,
		RunE:  runCreate(),
	}
}

func runCreate() RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		requests, err := cmd.Flags().GetInt("requests")
		if err != nil {
			return err
		}
		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			return err
		}

		stress := usecase.NewStressUseCase()
		_, err = stress.Execute(usecase.StressInputDTO{
			Url:         cmd.Flag("url").Value.String(),
			Requests:    requests,
			Concurrency: concurrency,
		})
		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	stressCmd := newCreateCmd()
	rootCmd.AddCommand(stressCmd)

	stressCmd.Flags().StringP("url", "u", "", "URL to stress test")
	stressCmd.Flags().IntP("requests", "r", 100, "Number of requests to perform")
	stressCmd.Flags().IntP("concurrency", "c", 1, "Number of multiple requests to make at a time")
	stressCmd.MarkFlagRequired("url")
}

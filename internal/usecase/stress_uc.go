package usecase

import (
	"fmt"
	"net/url"

	"github.com/carlosmeds/stress-test-cli/internal/infra/api"
)

type StressInputDTO struct {
	Url         string
	Requests    int
	Concurrency int
}

type StressUseCase struct{}

func NewStressUseCase() *StressUseCase {
	return &StressUseCase{}
}

func (u *StressUseCase) Execute(i StressInputDTO) (o string, err error) {
	fmt.Printf("Stress called for url %s with %d requests and %d concurrency\n", i.Url, i.Requests, i.Concurrency)

	if err := i.Validate(); err != nil {
		return "Error on validation", err
	}

	statusCode := api.RequestApi(i.Url)
	fmt.Printf("Status code: %d\n", statusCode)

	return "Use Case done!", nil
}

func (dto *StressInputDTO) Validate() error {
	if !isValidURL(dto.Url) {
		return fmt.Errorf("invalid URL: %s", dto.Url)
	}
	if dto.Requests <= 0 {
		return fmt.Errorf("requests must be greater than 0")
	}
	if dto.Concurrency <= 0 {
		return fmt.Errorf("concurrency must be greater than 0")
	}
	return nil
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

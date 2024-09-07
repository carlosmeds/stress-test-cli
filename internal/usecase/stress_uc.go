package usecase

import "fmt"

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

	return "Use Case done!", nil
}

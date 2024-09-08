package usecase

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/carlosmeds/stress-test-cli/internal/infra/api"
)

var (
	wg sync.WaitGroup
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

	fmt.Println("Starting stress test...")
	reqControl := make(chan struct{}, i.Concurrency)
	var statusCount sync.Map
	start := time.Now()
	for j := 0; j < i.Requests; j++ {
		wg.Add(1)
		reqControl <- struct{}{}

		go callApi(i.Url, reqControl, &statusCount)
	}
	wg.Wait()
	duration := time.Since(start)

	fmt.Println("\n\n---------------REPORT---------------")
	fmt.Printf("[URL]:      %s\n", i.Url)
	fmt.Printf("[REQUESTS]: %d\n", i.Requests)
	fmt.Printf("[DURATION]: %s\n\n", duration)

	found200 := false
	statusCount.Range(func(key, value interface{}) bool {
		fmt.Printf("%d requests returned status code %d\n", atomic.LoadInt64(value.(*int64)), key.(int))

		if key.(int) == 200 {
            found200 = true
        }
        return true
	})
	if !found200 {
        fmt.Println("No requests returned status code 200")
    }

	return "Use Case done!", nil
}

func callApi(url string, reqControl chan struct{}, statusCount *sync.Map) int {
	defer wg.Done()

	status := api.RequestApi(url)
	incrementStatusCount(statusCount, status)
	<-reqControl

	return status
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

func incrementStatusCount(statusCount *sync.Map, status int) {
	val, _ := statusCount.LoadOrStore(status, new(int64))
	atomic.AddInt64(val.(*int64), 1)
}

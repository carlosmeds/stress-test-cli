package api

import (
	"fmt"
	"net/http"
)

func RequestApi(url string) int {
	fmt.Println("Requesting", url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	return res.StatusCode
}

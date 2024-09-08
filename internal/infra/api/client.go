package api

import (
	"net/http"
)

func RequestApi(url string) int {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	return res.StatusCode
}

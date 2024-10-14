package cli

import "http-client/http"

func measureResponseTime(url string, rounds uint) (string, error) {
	// TODO

	return http.Get(url)
}

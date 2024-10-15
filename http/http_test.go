package http

import "testing"

func TestHttpGet(t *testing.T) {
	tests := []string{
		"gobyexample.com",
		"https://gobyexample.com/",
		"http://gobyexample.com/goroutines",
		"duckduckgo.com",
		"https://duckduckgo.com/",
	}

	for _, inputUrl := range tests {
		t.Run("GET "+inputUrl, func(t *testing.T) {
			response, err := Get(inputUrl)

			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			if len(response) == 0 {
				t.Error("expected a response")
			}
		})
	}
}

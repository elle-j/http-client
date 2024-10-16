package http

import (
	"http-client/utils"
	"testing"
)

func TestHttpGetSuccess(t *testing.T) {
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
				t.Errorf("expected no error, got: %s", err.Error())
			}
			if len(response) == 0 {
				t.Error("expected a response")
			}
		})
	}
}

func TestHttpGetError(t *testing.T) {
	type testCaseError struct {
		input                 string
		expectedErrorContains string
	}

	tests := []testCaseError{
		{
			input:                 "invalidHost",
			expectedErrorContains: "dial tcp: lookup invalidHost: no such host",
		},
	}

	for _, test := range tests {
		t.Run("GET "+test.input, func(t *testing.T) {
			response, err := Get(test.input)

			if err == nil {
				t.Errorf("expected error message to contain '%s', got no error", test.expectedErrorContains)
			} else {
				utils.ExpectContains(t, "error message", err.Error(), test.expectedErrorContains)
			}

			if len(response) != 0 {
				t.Errorf("expected no response, got: %s", response)
			}
		})
	}
}

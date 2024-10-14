package http

import "testing"

func expectMatch(t *testing.T, name string, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected %s to be %s, got %s", name, expected, actual)
	}
}

func TestParseRawUrlSuccess(t *testing.T) {
	type TestExpected struct {
		url    string
		scheme string
		host   string
		port   string
		path   string
	}

	type Test struct {
		input    string
		expected TestExpected
	}

	tests := [...]Test{
		{
			input: "gobyexample.com",
			expected: TestExpected{
				url:    "http://gobyexample.com:80/",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/",
			},
		},
		{
			input: "gobyexample.com:80",
			expected: TestExpected{
				url:    "http://gobyexample.com:80/",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/",
			},
		},
		{
			input: "gobyexample.com/goroutines",
			expected: TestExpected{
				url:    "http://gobyexample.com:80/goroutines",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/goroutines",
			},
		},
		{
			input: "http://gobyexample.com/goroutines",
			expected: TestExpected{
				url:    "http://gobyexample.com:80/goroutines",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/goroutines",
			},
		},
		{
			input: "https://gobyexample.com/goroutines",
			expected: TestExpected{
				url:    "https://gobyexample.com:443/goroutines",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/goroutines",
			},
		},
		{
			input: "https://gobyexample.com/#",
			expected: TestExpected{
				url:    "https://gobyexample.com:443/",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
		{
			input: "https://gobyexample.com#",
			expected: TestExpected{
				url:    "https://gobyexample.com:443/",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
		{
			input: "https://gobyexample.com/#heading",
			expected: TestExpected{
				url:    "https://gobyexample.com:443/#heading",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
	}

	for _, test := range tests {
		t.Run("Parse "+test.input, func(t *testing.T) {
			actual, err := parseRawUrl(test.input)

			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			expectMatch(t, "url", actual.String(), test.expected.url)
			expectMatch(t, "scheme", actual.Scheme, test.expected.scheme)
			expectMatch(t, "host", actual.Host, test.expected.host)
			expectMatch(t, "port", actual.Port(), test.expected.port)
			expectMatch(t, "path", actual.Path, test.expected.path)
		})
	}
}

func TestParseRawUrlError(t *testing.T) {
	type TestError struct {
		input        string
		errorMessage string
	}

	tests := [...]TestError{
		{
			input:        "https://",
			errorMessage: "provide a host in the URL",
		},
		{
			input:        "http://",
			errorMessage: "provide a host in the URL",
		},
		{
			input:        "file://localhost/",
			errorMessage: "only 'http' and 'https' schemes are currently supported",
		},
	}

	for _, test := range tests {
		t.Run("Parse invalid "+test.input, func(t *testing.T) {
			actualUrl, actualErr := parseRawUrl(test.input)

			if actualErr == nil {
				t.Errorf("expected error message '%s', got no error", test.errorMessage)
			}
			expectMatch(t, "error message", actualErr.Error(), test.errorMessage)

			if actualUrl != nil {
				t.Errorf("unexpected url, got %s", actualUrl.String())
			}
		})
	}
}

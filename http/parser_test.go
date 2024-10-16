package http

import (
	"http-client/utils"
	"testing"
)

func TestParseRawUrlSuccess(t *testing.T) {
	type testExpected struct {
		url    string
		scheme string
		host   string
		port   string
		path   string
	}

	type testCaseSuccess struct {
		input    string
		expected testExpected
	}

	tests := [...]testCaseSuccess{
		{
			input: "gobyexample.com",
			expected: testExpected{
				url:    "http://gobyexample.com:80/",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/",
			},
		},
		{
			input: "gobyexample.com:80",
			expected: testExpected{
				url:    "http://gobyexample.com:80/",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/",
			},
		},
		{
			input: "gobyexample.com/goroutines",
			expected: testExpected{
				url:    "http://gobyexample.com:80/goroutines",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/goroutines",
			},
		},
		{
			input: "http://gobyexample.com/goroutines",
			expected: testExpected{
				url:    "http://gobyexample.com:80/goroutines",
				scheme: "http",
				host:   "gobyexample.com:80",
				port:   "80",
				path:   "/goroutines",
			},
		},
		{
			input: "https://gobyexample.com/goroutines",
			expected: testExpected{
				url:    "https://gobyexample.com:443/goroutines",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/goroutines",
			},
		},
		{
			input: "https://gobyexample.com/#",
			expected: testExpected{
				url:    "https://gobyexample.com:443/",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
		{
			input: "https://gobyexample.com#",
			expected: testExpected{
				url:    "https://gobyexample.com:443/",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
		{
			input: "https://gobyexample.com/#heading",
			expected: testExpected{
				url:    "https://gobyexample.com:443/#heading",
				scheme: "https",
				host:   "gobyexample.com:443",
				port:   "443",
				path:   "/",
			},
		},
		{
			input: "localhost:8080",
			expected: testExpected{
				url:    "http://localhost:8080/",
				scheme: "http",
				host:   "localhost:8080",
				port:   "8080",
				path:   "/",
			},
		},
		{
			input: "127.0.0.1:8080",
			expected: testExpected{
				url:    "http://127.0.0.1:8080/",
				scheme: "http",
				host:   "127.0.0.1:8080",
				port:   "8080",
				path:   "/",
			},
		},
		{
			input: "postgresql://user:password@localhost/my_database",
			expected: testExpected{
				url:    "postgresql://user:password@localhost/my_database",
				scheme: "postgresql",
				host:   "localhost",
				port:   "",
				path:   "/my_database",
			},
		},
		{
			input: "postgresql://user:password@localhost:5432/my_database",
			expected: testExpected{
				url:    "postgresql://user:password@localhost:5432/my_database",
				scheme: "postgresql",
				host:   "localhost:5432",
				port:   "5432",
				path:   "/my_database",
			},
		},
		{
			input: "mongodb://username:password@localhost:27017/my_database",
			expected: testExpected{
				url:    "mongodb://username:password@localhost:27017/my_database",
				scheme: "mongodb",
				host:   "localhost:27017",
				port:   "27017",
				path:   "/my_database",
			},
		},
		{
			input: " localhost:8080    ",
			expected: testExpected{
				url:    "http://localhost:8080/",
				scheme: "http",
				host:   "localhost:8080",
				port:   "8080",
				path:   "/",
			},
		},
		{
			input: "\tlocalhost:8080\t\t",
			expected: testExpected{
				url:    "http://localhost:8080/",
				scheme: "http",
				host:   "localhost:8080",
				port:   "8080",
				path:   "/",
			},
		},
	}

	for _, test := range tests {
		t.Run("Parse "+test.input, func(t *testing.T) {
			actual, err := parseRawUrl(test.input)

			if err != nil {
				t.Errorf("expected no error, got: %s", err.Error())
			}
			utils.ExpectMatch(t, "url", actual.String(), test.expected.url)
			utils.ExpectMatch(t, "scheme", actual.Scheme, test.expected.scheme)
			utils.ExpectMatch(t, "host", actual.Host, test.expected.host)
			utils.ExpectMatch(t, "port", actual.Port(), test.expected.port)
			utils.ExpectMatch(t, "path", actual.Path, test.expected.path)
		})
	}
}

func TestParseRawUrlError(t *testing.T) {
	type testCaseError struct {
		input                 string
		expectedErrorContains string
	}

	tests := [...]testCaseError{
		{
			input:                 "https://",
			expectedErrorContains: "provide a host in the URL",
		},
		{
			input:                 "http://",
			expectedErrorContains: "provide a host in the URL",
		},
		{
			input:                 "http://localhost:invalidPort",
			expectedErrorContains: "invalid port \":invalidPort\" after host",
		},
	}

	for _, test := range tests {
		t.Run("Parse invalid "+test.input, func(t *testing.T) {
			actualUrl, actualErr := parseRawUrl(test.input)

			if actualErr == nil {
				t.Errorf("expected error message to contain '%s', got no error", test.expectedErrorContains)
			} else {
				utils.ExpectContains(t, "error message", actualErr.Error(), test.expectedErrorContains)
			}

			if actualUrl != nil {
				t.Errorf("expected no url, got: %s", actualUrl.String())
			}
		})
	}
}

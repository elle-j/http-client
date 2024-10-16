package cli

import (
	"http-client/utils"
	"strings"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	tests := []string{
		"get -url https://gobyexample.com/",
		"get -url https://gobyexample.com/ -rounds 1",
		"get -url https://gobyexample.com/ -rounds 2",
		"get -rounds 1 -url https://gobyexample.com/",
	}

	for _, command := range tests {
		t.Run("Command "+command, func(t *testing.T) {
			err := Run(strings.Split(command, " "))

			if err != nil {
				t.Errorf("expected no error, got: %s", err.Error())
			}
		})
	}
}

func TestRunError(t *testing.T) {
	type testCaseError struct {
		command               string
		expectedErrorContains string
	}

	tests := []testCaseError{
		{
			command:               "",
			expectedErrorContains: "Usage",
		},
		{
			command:               "invalidCommand -invalidFlag",
			expectedErrorContains: "Usage",
		},
		{
			command:               "get",
			expectedErrorContains: "Usage",
		},
		{
			command:               "get -invalidFlag",
			expectedErrorContains: "flag provided but not defined: -invalidFlag",
		},
		{
			command:               "get -url",
			expectedErrorContains: "flag needs an argument: -url",
		},
		{
			command:               "get -url ",
			expectedErrorContains: "you must provide a -url flag with a specific URL",
		},
		{
			command:               "get -url https://gobyexample.com/ -rounds",
			expectedErrorContains: "flag needs an argument: -rounds",
		},
		{
			command:               "get -url https://gobyexample.com/ -rounds -1",
			expectedErrorContains: "if providing -rounds, it must be a positive integer",
		},
		{
			command:               "get -url https://gobyexample.com/ -rounds 2.5",
			expectedErrorContains: "invalid value \"2.5\" for flag -rounds: parse error",
		},
		{
			command:               "get -url https://gobyexample.com/ -rounds not_a_number",
			expectedErrorContains: "invalid value \"not_a_number\" for flag -rounds: parse error",
		},
	}

	for _, test := range tests {
		t.Run("Invalid command "+test.command, func(t *testing.T) {
			err := Run(strings.Split(test.command, " "))

			if err == nil {
				t.Errorf("expected error message to contain '%s', got no error", test.expectedErrorContains)
			} else {
				utils.ExpectContains(t, "error message", err.Error(), test.expectedErrorContains)
			}
		})
	}
}

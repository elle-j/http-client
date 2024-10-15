package cli

import (
	"strings"
	"testing"
)

func expectContains(t *testing.T, name string, actual string, expected string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %s to contain '%s', got: %s", name, expected, actual)
	}
}

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
				t.Errorf("unexpected error: %s", err.Error())
			}
		})
	}
}

func TestRunError(t *testing.T) {
	type TestError struct {
		command              string
		errorMessageContains string
	}

	tests := []TestError{
		{
			command:              "",
			errorMessageContains: "Usage",
		},
		{
			command:              "invalidCommand -invalidFlag",
			errorMessageContains: "Usage",
		},
		{
			command:              "get",
			errorMessageContains: "Usage",
		},
		{
			command:              "get -invalidFlag",
			errorMessageContains: "flag provided but not defined: -invalidFlag",
		},
		{
			command:              "get -url",
			errorMessageContains: "flag needs an argument: -url",
		},
		{
			command:              "get -url ",
			errorMessageContains: "you must provide a -url flag with a specific URL",
		},
		{
			command:              "get -url https://gobyexample.com/ -rounds",
			errorMessageContains: "flag needs an argument: -rounds",
		},
		{
			command:              "get -url https://gobyexample.com/ -rounds -1",
			errorMessageContains: "if providing -rounds, it must be a positive integer",
		},
		{
			command:              "get -url https://gobyexample.com/ -rounds 2.5",
			errorMessageContains: "invalid value \"2.5\" for flag -rounds: parse error",
		},
		{
			command:              "get -url https://gobyexample.com/ -rounds not_a_number",
			errorMessageContains: "invalid value \"not_a_number\" for flag -rounds: parse error",
		},
	}

	for _, test := range tests {
		t.Run("Invalid command "+test.command, func(t *testing.T) {
			err := Run(strings.Split(test.command, " "))

			if err == nil {
				t.Errorf("expected error message to contain '%s', got no error", test.errorMessageContains)
			}
			expectContains(t, "error message", err.Error(), test.errorMessageContains)
		})
	}
}

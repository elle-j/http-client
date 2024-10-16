package utils

import (
	"strings"
	"testing"
)

func ExpectMatch(t *testing.T, name string, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected %s to be '%s', got: %s", name, expected, actual)
	}
}

func ExpectContains(t *testing.T, name string, actual string, expected string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %s to contain '%s', got: %s", name, expected, actual)
	}
}

package utils

import (
	"strings"
	"testing"
)

func ExpectMatch[T comparable](t *testing.T, name string, actual T, expected T) {
	if actual != expected {
		t.Errorf("expected %s to be '%v', got: %v", name, expected, actual)
	}
}

func ExpectContains(t *testing.T, name string, actual string, expected string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %s to contain '%s', got: %s", name, expected, actual)
	}
}

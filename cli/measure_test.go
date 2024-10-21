package cli

import (
	"fmt"
	"http-client/utils"
	"slices"
	"testing"
	"time"
)

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}

	return duration
}

var durationsUnsortedEven = []time.Duration{parseDuration("60ms"), parseDuration("55ms"), parseDuration("50ms"), parseDuration("65ms")}
var durationsUnsortedOdd = []time.Duration{parseDuration("60ms"), parseDuration("55ms"), parseDuration("100ms"), parseDuration("50ms"), parseDuration("65ms")}

type testCase struct {
	input    []time.Duration
	expected time.Duration
}

func TestCalculateSum(t *testing.T) {
	tests := []testCase{
		{
			input:    []time.Duration{},
			expected: time.Duration(0),
		},
		{
			input:    durationsUnsortedEven,
			expected: parseDuration("230ms"),
		},
		{
			input:    durationsUnsortedOdd,
			expected: parseDuration("330ms"),
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("Case #%d", index+1), func(t *testing.T) {
			actual := calculateSum(test.input)

			utils.ExpectMatch(t, "sum", actual, test.expected)
		})
	}
}

func TestCalculateMean(t *testing.T) {
	tests := []testCase{
		{
			input:    []time.Duration{},
			expected: time.Duration(0),
		},
		{
			input:    durationsUnsortedEven,
			expected: parseDuration("57500us"),
		},
		{
			input:    durationsUnsortedOdd,
			expected: parseDuration("66ms"),
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("Case #%d", index+1), func(t *testing.T) {
			actual := calculateMean(test.input)

			utils.ExpectMatch(t, "mean", actual, test.expected)
		})
	}
}

func TestCalculateMedian(t *testing.T) {
	durationsSortedEven := durationsUnsortedEven
	durationsSortedOdd := durationsUnsortedOdd
	slices.Sort(durationsSortedEven)
	slices.Sort(durationsSortedOdd)

	tests := []testCase{
		{
			input:    []time.Duration{},
			expected: time.Duration(0),
		},
		{
			input:    durationsSortedEven,
			expected: parseDuration("57500us"),
		},
		{
			input:    durationsSortedOdd,
			expected: parseDuration("60ms"),
		},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("Case #%d", index+1), func(t *testing.T) {
			actual := calculateMedian(test.input)

			utils.ExpectMatch(t, "median", actual, test.expected)
		})
	}
}

func TestRequestAndMeasure(t *testing.T) {
	const url string = "https://gobyexample.com/"
	const rounds int = 4

	measures, err := requestAndMeasure(url, rounds)

	if err != nil {
		t.Errorf("expected no error, got: %s", err.Error())
	} else {
		summary := measures.stringify()
		utils.ExpectContains(t, "summary", summary, "URL: https://gobyexample.com")
		utils.ExpectContains(t, "summary", summary, fmt.Sprintf("Number of requests: %d", rounds))
		utils.ExpectContains(t, "summary", summary, "Size")
		utils.ExpectContains(t, "summary", summary, "Fastest time")
		utils.ExpectContains(t, "summary", summary, "Slowest time")
		utils.ExpectContains(t, "summary", summary, "Mean time")
		utils.ExpectContains(t, "summary", summary, "Median time")
	}
}

package cli

import (
	"fmt"
	"http-client/http"
	"slices"
	"strings"
	"time"
)

type responseData struct {
	response  string
	size      uint32
	durations []time.Duration
}

type responseSummary struct {
	url      string
	rounds   int
	response string
	size     uint32
	fastest  time.Duration
	slowest  time.Duration
	mean     time.Duration
	median   time.Duration
}

func (summary *responseSummary) stringify() string {
	var result strings.Builder
	result.WriteString("\n================\n")
	result.WriteString("RESPONSE \n")
	result.WriteString("================\n\n")
	result.WriteString(summary.response)
	result.WriteString("\n\n================\n")
	result.WriteString("RESPONSE SUMMARY\n")
	result.WriteString("================\n\n")
	result.WriteString(fmt.Sprintf("URL: %v\n", summary.url))
	result.WriteString(fmt.Sprintf("Number of requests: %v\n", summary.rounds))
	result.WriteString(fmt.Sprintf("Size (bytes): %v\n", summary.size))
	result.WriteString(fmt.Sprintf("Fastest time: %v\n", summary.fastest))
	result.WriteString(fmt.Sprintf("Slowest time: %v\n", summary.slowest))
	result.WriteString(fmt.Sprintf("Mean time: %v\n", summary.mean))
	result.WriteString(fmt.Sprintf("Median time: %v\n", summary.median))

	return result.String()
}

func calculateSum(durations []time.Duration) time.Duration {
	var sum time.Duration = 0
	for _, value := range durations {
		sum += value
	}

	return sum
}

func calculateMean(durations []time.Duration) time.Duration {
	return calculateSum(durations) / time.Duration(len(durations))
}

func calculateMedian(sortedDurations []time.Duration) time.Duration {
	count := len(sortedDurations)
	midIndex := count / 2
	median := sortedDurations[midIndex]
	if count%2 == 0 {
		median = (sortedDurations[midIndex-1] + median) / 2
	}

	return median
}

func summarize(url string, data *responseData) *responseSummary {
	slices.Sort(data.durations)

	return &responseSummary{
		url:      url,
		rounds:   len(data.durations),
		response: data.response,
		size:     data.size,
		fastest:  data.durations[0],
		slowest:  data.durations[len(data.durations)-1],
		mean:     calculateMean(data.durations),
		median:   calculateMedian(data.durations),
	}
}

func timeRequest(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	duration := time.Since(start)

	return duration, err
}

func repeatRequest(url string, rounds int, data *responseData) error {
	for i := 0; i < rounds; i++ {
		duration, err := timeRequest(url)
		if err != nil {
			return err
		}

		data.durations[i] = duration
	}

	return nil
}

func getInitialResponseData(url string, rounds int) (*responseData, error) {
	// Perform an initial request that is not measured.
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data := &responseData{
		response:  response,
		size:      uint32(len(response)),
		durations: make([]time.Duration, rounds),
	}

	return data, nil
}

// Make GET requests to the specific URL and measure the response times.
func requestAndMeasure(url string, rounds int) (*responseSummary, error) {
	data, err := getInitialResponseData(url, rounds)
	if err != nil {
		return nil, err
	}

	err = repeatRequest(url, rounds, data)
	if err != nil {
		return nil, err
	}

	return summarize(url, data), nil
}

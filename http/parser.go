package http

import (
	"errors"
	"net/url"
	"strings"
)

func addDefaultPorts(parsedUrl *url.URL) {
	if parsedUrl.Scheme == "http" {
		parsedUrl.Host += ":80"
	} else if parsedUrl.Scheme == "https" {
		parsedUrl.Host += ":443"
	}
}

func parseScheme(rawUrl string) string {
	substrings := strings.Split(rawUrl, "://")
	hasScheme := len(substrings) > 1
	if hasScheme {
		return substrings[0]
	}
	return ""
}

func parseRawUrl(rawUrl string) (*url.URL, error) {
	explicitScheme := parseScheme(rawUrl)
	if explicitScheme == "" {
		// http is assumed if a scheme is not explicitly provided.
		rawUrl = "http://" + rawUrl
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	if parsedUrl.Host == "" {
		return nil, errors.New("provide a host in the URL")
	}

	if parsedUrl.Path == "" {
		parsedUrl.Path = "/"
	}

	if parsedUrl.Port() == "" {
		addDefaultPorts(parsedUrl)
	}

	return parsedUrl, nil
}

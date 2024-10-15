package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/url"
	"time"
)

// Establish a network connection to the address.
func connect(parsedUrl *url.URL) (net.Conn, error) {
	var connection net.Conn
	var err error

	if parsedUrl.Scheme == "https" {
		connection, err = tls.Dial("tcp", parsedUrl.Host, &tls.Config{})
	} else {
		timeout, _ := time.ParseDuration("10s")
		connection, err = net.DialTimeout("tcp", parsedUrl.Host, timeout)
	}

	if err != nil {
		return nil, err
	}

	return connection, nil
}

func buildRequestString(parsedUrl *url.URL) []byte {
	var request bytes.Buffer
	request.WriteString(fmt.Sprintf("GET %v HTTP/1.1\r\n", parsedUrl.Path))
	request.WriteString(fmt.Sprintf("Host: %v\r\n", parsedUrl.Host))
	request.WriteString("Accept: */*\r\n")
	request.WriteString("Connection: close\r\n")
	request.WriteString("\r\n")

	return request.Bytes()
}

// Write data to the connection.
func sendRequest(connection net.Conn, parsedUrl *url.URL) error {
	_, err := connection.Write(buildRequestString(parsedUrl))

	return err
}

// Read data from the connection.
func readResponse(connection net.Conn) (string, error) {
	response, err := io.ReadAll(connection)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(response), nil
}

// Send an HTTP GET request to the provided URL.
func Get(rawUrl string) (string, error) {
	parsedUrl, parseErr := parseRawUrl(rawUrl)
	if parseErr != nil {
		return "", parseErr
	}

	connection, connectionErr := connect(parsedUrl)
	if connectionErr != nil {
		return "", connectionErr
	}
	defer connection.Close()

	requestErr := sendRequest(connection, parsedUrl)
	if requestErr != nil {
		return "", requestErr
	}

	return readResponse(connection)
}

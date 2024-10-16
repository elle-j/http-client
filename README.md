# An HTTP Client CLI for Measuring Response Times

This is an HTTP client CLI for measuring the response times of requests to a given URL.

It uses a custom light-weight implementation for the TCP/IP application layer for communicating with the lower transport layer.

## Purpose

The main use case of the CLI is for testing the response times of retrieving your own resources (e.g. personal websites) or those of others (e.g. public resources on the web). Additionally, it can be used for inspecting the response.

### Why a Custom TCP/IP Application Layer?

Since Golang (Go) provides the `net/http` package, why not just use that? Well, that would be the most convenient way of going about it, but another purpose of creating this HTTP client is to learn a little bit of Go and familiarize myself with the lower level TCP/IP stack interfaces.

Therefore, the application layer is more light-weight and limited than Go's, but handles the communication with the transport layer manually.

## Usage

```
$ ./http-client --help

Usage
  ./http-client [command] [flags]

Commands
  get        Make HTTP GET requests and see the response times

Flags
  -url       The URL to send the request to
  -rounds    The number of requests to make (default 1)

Example
  ./http-client -url https://gobyexample.com -rounds 10
```

### Example Input and Output

The below command will make 10 GET requests to `https://gobyexample.com/` and then output:

* The GET response
* The full URL used
* The number of requests
* The response size
* The fastest response time
* The slowest response time
* The mean response time
* The median response time

```
$ ./http-client get -url https://gobyexample.com/ -rounds 10

================
RESPONSE
================

HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 6824
Connection: close

<rest of response>

================
RESPONSE SUMMARY
================

URL: https://gobyexample.com:443/
Number of requests: 10
Size (bytes): 7308
Fastest time: 48.0016ms
Slowest time: 51.6554ms
Mean time: 49.75598ms
Median time: 49.7726ms
```

### Considerations

#### Response Time

Note that the response time for each request includes:
* Parsing the raw URL
* Connecting to the network address
  * For TLS connections, this includes a TLS handshake
* Writing to the connection (sending your request)
* Reading from the connection (retrieving the response)

#### Raw URL

If the provided URL does not explicitly contain a [scheme](https://developer.mozilla.org/en-US/docs/Web/URI/Schemes) (e.g. prefixed with `https://`), then `http` is assumed.

## Getting Started

### Prerequisites

* [Golang](https://go.dev/dl/)
  * See the file [go.mod](./go.mod) for the version used in this project.
  * You can verify your own version and installation by running `go version`.

### Build the CLI

Run the below command from the root of the project to create a binary executable called `http-client`:

```
go build
```

### Run the CLI

To run the CLI using the executable created in the [previous step](#build-the-cli), run the following command from the same directory as the executable, and replace `[command]` and `[flags]` with the supported options explained in the section [Usage](#usage):

**On Linux and Mac:**
```
./http-client [command] [flags]
```

**On Windows:**
```
./http-client.exe [command] [flags]
```

### Run Tests

To run the tests, run the following command from the root of the project:

```
go test ./...
```

To run the tests in verbose mode, add the `-v` flag:

```
go test -v ./...
```

## License

This software is licensed under the terms of the [MIT license](LICENSE).

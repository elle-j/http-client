package main

import (
	"http-client/cli"
	"os"
)

func main() {
	cli.Run(os.Args[1:])
}

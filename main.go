package main

import (
	"http-client/cli"
	"log"
	"os"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

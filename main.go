package main

import (
	"http-client/cli"
	"log"
	"os"
)

func main() {
	err := cli.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

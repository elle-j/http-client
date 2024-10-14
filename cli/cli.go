package cli

import (
	"flag"
	"fmt"
	"log"
)

type getCommand struct {
	name  string
	flags *flag.FlagSet
}

func newGetCommand() (*getCommand, *string, *uint) {
	command := &getCommand{
		name:  "get",
		flags: flag.NewFlagSet("get", flag.ExitOnError),
	}
	url := command.flags.String("url", "", "The URL to request (e.g. https://example.com)")
	rounds := command.flags.Uint("rounds", 1, "The number of rounds/iterations to perform the request")

	return command, url, rounds
}

func getUsage() string {
	return "TODO USAGE"
}

func printUsageAndExit() {
	log.Fatal(getUsage())
}

func Run(args []string) {
	if len(args) < 2 {
		printUsageAndExit()
	}

	command, url, rounds := newGetCommand()

	if args[0] != command.name {
		printUsageAndExit()
	}

	command.flags.Parse(args[1:])

	if *url == "" {
		fmt.Println("You must provide a -url flag with a specific URL.")
		printUsageAndExit()
	}

	measures := measureResponseTime(*url, *rounds)
	fmt.Println(measures)
}

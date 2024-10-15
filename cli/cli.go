package cli

import (
	"errors"
	"flag"
	"fmt"
)

type getCommand struct {
	name  string
	flags *flag.FlagSet
}

const usage = `
Usage
  http_client [command] [flags]

Commands
  get        Make an HTTP GET request and see the response time

Flags
  -url       The URL to request (e.g. https://example.com)
  -rounds    The number of requests to make when measuring response times (default 1)`

func newGetCommand() (*getCommand, *string, *int) {
	command := &getCommand{
		name:  "get",
		flags: flag.NewFlagSet("get", flag.ContinueOnError),
	}
	url := command.flags.String("url", "", "The URL to request (e.g. https://example.com)")
	rounds := command.flags.Int("rounds", 1, "The number of requests to make when measuring response times")

	return command, url, rounds
}

// Start the CLI.
func Run(args []string) error {
	if len(args) < 2 {
		return errors.New(usage)
	}

	command, url, rounds := newGetCommand()

	if args[0] != command.name {
		return errors.New(usage)
	}

	err := command.flags.Parse(args[1:])
	if err != nil {
		return err
	}

	if *url == "" {
		return errors.New("you must provide a -url flag with a specific URL\n" + usage)
	}

	if *rounds < 1 {
		return errors.New("if providing -rounds, it must be a positive integer\n" + usage)
	}

	measures, err := requestAndMeasure(*url, *rounds)
	if err != nil {
		return err
	}

	fmt.Println(measures.stringify())

	return nil
}

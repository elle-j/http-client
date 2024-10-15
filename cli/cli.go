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

func newGetCommand() (*getCommand, *string, *int) {
	command := &getCommand{
		name:  "get",
		flags: flag.NewFlagSet("get", flag.ExitOnError),
	}
	url := command.flags.String("url", "", "The URL to request (e.g. https://example.com)")
	rounds := command.flags.Int("rounds", 1, "The number of rounds/iterations to perform the request")

	return command, url, rounds
}

func getUsage() string {
	return "TODO USAGE"
}

// Start the CLI.
func Run(args []string) error {
	if len(args) < 2 {
		return errors.New(getUsage())
	}

	command, url, rounds := newGetCommand()

	if args[0] != command.name {
		return errors.New(getUsage())
	}

	command.flags.Parse(args[1:])

	if *url == "" {
		return errors.New("You must provide a -url flag with a specific URL.\n\n" + getUsage())
	}

	if *rounds < 1 {
		return errors.New("If providing -rounds, it must be a positive number.\n\n" + getUsage())
	}

	measures, err := requestAndMeasure(*url, *rounds)
	if err != nil {
		return err
	}

	fmt.Println(measures.stringify())

	return nil
}

package cmd

import "github.com/urfave/cli/v2"

func DomainCommand() *cli.Command {
	return &cli.Command{
		Name:        "domain",
		Usage:       "Add core domain in project",
		UsageText:   "gt domain [--name|-n=][DomainName]",
		Description: "The domain command create a new core domain with go struct，this command will generate some necessary files or dir in core directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "example",
				Usage:   "[[--name|-n=]ServiceName]",
			},
		},
		Action: domainCommandFunc,
	}
}

func domainCommandFunc(ctx *cli.Context) error {

	return nil
}

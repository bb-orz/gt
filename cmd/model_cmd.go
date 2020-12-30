package cmd

import "github.com/urfave/cli/v2"

func ModelCommand() *cli.Command {
	return &cli.Command{
		Name:        "model",
		Usage:       "Add core model",
		UsageText:   "gt model [--name|-n=][ModelName]",
		Description: "The model command create a new core model with go struct，this command will generate some necessary files or dir in core directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "example",
				Usage:   "[[--name|-n=]ServiceName]",
			},
		},
		Action: modelCommandFunc,
	}
}

func modelCommandFunc(ctx *cli.Context) error {

	return nil
}

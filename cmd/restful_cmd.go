package cmd

import "github.com/urfave/cli/v2"

func RestfulCommand() *cli.Command {
	return &cli.Command{
		Name:        "restful",
		Usage:       "Add Restful API",
		UsageText:   "gt restful [--name|-n=][RestfulName]",
		Description: "The restful command create a new restful api with go struct，this command will generate some necessary files or dir in restful directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "example",
				Usage:   "[[--name|-n=]RestfulName]",
			},
		},
		Action: restfulCommandFunc,
	}
}

func restfulCommandFunc(ctx *cli.Context) error {

	return nil
}

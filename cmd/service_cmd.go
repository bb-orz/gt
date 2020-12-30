package cmd

import "github.com/urfave/cli/v2"

func ServiceCommand() *cli.Command {
	return &cli.Command{
		Name:        "service",
		Usage:       "Add Application Service",
		UsageText:   "gt service [--name|-n=][ServiceName]",
		Description: "The service command create a new service go interface，this command will generate some necessary files in service directory.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "goapp",
				Usage:   "[[--name|-n=]ServiceName]",
			},
		},
		Action: serviceCommandFunc,
	}
}

func serviceCommandFunc(ctx *cli.Context) error {

	return nil
}

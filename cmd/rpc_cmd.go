package cmd

import "github.com/urfave/cli/v2"

func RPCCommand() *cli.Command {
	return &cli.Command{
		Name:        "rpc",
		Usage:       "Add RPC Service",
		UsageText:   "gt rpc [--name|-n=][RPCName]",
		Description: "The rpc command create a new rpc service with go struct，this command will generate some necessary files or dir in rpc directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "example",
				Usage:   "[[--name|-n=]RPCName]",
			},
		},
		Action: RPCCommandAction,
	}
}

func RPCCommandAction(ctx *cli.Context) error {

	return nil
}

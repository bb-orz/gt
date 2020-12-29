package cmd

import "github.com/urfave/cli/v2"

func InitCommand() *cli.Command  {
	return &cli.Command{
		Name:        "init",
		Usage:       "Go Web Application Initialization",
		UsageText:   "gt init [--name|-n] [project_name] [--git|-g]",
		Description: "The init command create a new go web application in current directory，this command will generate some necessary folders and files, create a project.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "name",
				Aliases:[]string{"n"},
				Usage: "project name",
			},
			&cli.BoolFlag{
				Name: "git",
				Aliases:[]string{"g"},
				Usage: "git init",
			},
		},
		Action: initCommandFunc,
	}
}

func initCommandFunc(ctx *cli.Context) error {
	var name string
	if ctx.NArg() == 1 {
		name = ctx.Args().Get(0)
	}else{
		name = ctx.String("name")
	}

	if name == "" {
		name = "goapp"
	}


	return nil
}
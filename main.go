package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/cmd"
	"gt/utils"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "gt"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Usage = "A generation tool of go app scaffold which base on bb-orz/goinfras."
	app.UsageText = "gt [option] [command] [args]"
	app.ArgsUsage = "[args and such]"
	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("gt (goinfras tool) is a generation tool of go app scaffold which base on bb-orz/goinfras.")
		return nil
	}

	app.Commands = []*cli.Command{
		cmd.InitCommand(), // 初始化命令

	}

	app.Action = func(ctx *cli.Context) error {
		utils.CommandLogger.Debug(utils.AppCmd, "debug log ...")
		utils.CommandLogger.Info(utils.AppCmd, "info log ...")
		utils.CommandLogger.Warning(utils.AppCmd, "warning log ...")
		utils.CommandLogger.Error(utils.AppCmd, errors.New("error log ... "))
		utils.CommandLogger.Fail(utils.AppCmd, errors.New("fail log ... "))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		utils.CommandLogger.Fail(utils.AppCmd, err)
	}

}

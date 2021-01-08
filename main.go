package main

import (
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
		cmd.InitCommand(),    // 初始化命令
		cmd.ModelCommand(),   // 创建数据库表模型命令
		cmd.DomainCommand(),  // 创建领域模块命令
		cmd.ServiceCommand(), // 服务创建命令
		cmd.RestfulCommand(), // Restful API创建命令
		cmd.RPCCommand(),     // RPC Service 创建命令
	}

	err := app.Run(os.Args)
	if err != nil {
		utils.CommandLogger.Fail(utils.AppCmd, err)
	}

}

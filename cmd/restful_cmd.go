package cmd

import (
	"fmt"
	"github.com/bb-orz/gt/libs/libRestful"
	"github.com/bb-orz/gt/utils"
	"github.com/urfave/cli/v2"
	"io"
)

func RestfulCommand() *cli.Command {
	return &cli.Command{
		Name:        "restful",
		Usage:       "Add Restful API",
		UsageText:   "gt restful [--name|-n=][RestfulName]",
		Description: "The restful command create a new restful api with go struct，this command will generate some necessary files or dir in restful directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "example"},
			&cli.StringFlag{Name: "engine", Aliases: []string{"e"}, Value: "gin"},
			&cli.StringFlag{Name: "output_path", Aliases: []string{"o"}, Value: "./restful"},
		},
		Action: RestfulCommandAction,
	}
}

func RestfulCommandAction(ctx *cli.Context) error {
	var cmdParams = libRestful.CmdParams{
		Name:       ctx.String("name"),
		Engine:     ctx.String("engine"),
		OutputPath: ctx.String("output_path"),
	}

	var err error
	var fileWriter io.Writer

	// 检查服务接口是否存在，不存在则创建
	fileName := cmdParams.OutputPath + "/" + cmdParams.Name + "_restful.go"
	if !IsServiceFileExist(fileName) {
		if fileWriter, err = utils.CreateFile(fileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameRestful, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameRestful, fmt.Sprintf("Create %s Restful API File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), fileName))

			switch cmdParams.Engine {
			case "gin":
				// 格式化写入
				if err = libRestful.NewFormatterGinEngine().Format(cmdParams.Name).WriteOut(fileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRestful, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRestful, fmt.Sprintf("Write %s Restful API (Gin Engine) File Successful!", cmdParams.Name))
				}

			case "echo":
				// 格式化写入
				if err = libRestful.NewFormatterEchoEngine().Format(cmdParams.Name).WriteOut(fileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRestful, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRestful, fmt.Sprintf("Write %s Restful API (Echo Engine) File Successful!", cmdParams.Name))
				}
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameRestful, fmt.Sprintf("%s Restful API File Is Exist!", cmdParams.Name))
	}

	return nil
}

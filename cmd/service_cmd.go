package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/libs/libService"
	"gt/utils"
	"io"
	"os"
)

func ServiceCommand() *cli.Command {
	return &cli.Command{
		Name:        "service",
		Usage:       "Add Application Service",
		UsageText:   "gt service [--name|-n=][ServiceName]",
		Description: "The service command create a new service go interface，this command will generate some necessary files in service directory.",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "example"},
			&cli.StringFlag{Name: "version", Aliases: []string{"v"}, Value: "V1"},
			&cli.StringFlag{Name: "interface_output_path", Aliases: []string{"o"}, Value: "./services"},
			&cli.StringFlag{Name: "implement_output_path", Aliases: []string{"c"}, Value: "./services"},
			&cli.StringFlag{Name: "dto_output_path", Aliases: []string{"d"}, Value: "./dtos"},
		},
		Action: ServiceCommandAction,
	}
}

func ServiceCommandAction(ctx *cli.Context) error {
	var err error
	var cmdParams = libService.CmdParams{
		Name:        ctx.String("name"),
		Version:     ctx.String("version"),
		IOutputPath: ctx.String("interface_output_path"),
		MOutputPath: ctx.String("implement_output_path"),
		DOutputPath: ctx.String("dto_output_path"),
	}

	var interfaceFile, implementFile io.Writer

	// 检查服务接口是否存在，不存在则创建
	interfaceFileName := cmdParams.IOutputPath + "/" + cmdParams.Name + "_service.go"
	if !IsServiceFileExist(interfaceFileName) {
		if interfaceFile, err = utils.CreateFile(interfaceFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameService, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Create %s Service Interface File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), interfaceFileName))

			// 格式化写入
			if err = libService.NewFormatterServiceInterface().Format(cmdParams.Name, cmdParams.Version).WriteOut(interfaceFile); err != nil {
				utils.CommandLogger.Error(utils.CommandNameService, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Write %s Service Interface File Successful!", cmdParams.Name))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameService, fmt.Sprintf("%s Service Interface File Is Exist!", cmdParams.Name))
	}

	// 创建实现服务指定版本的文件
	implementFileName := cmdParams.MOutputPath + "/" + cmdParams.Name + "_service_" + cmdParams.Version + ".go"
	if !IsServiceFileExist(implementFileName) {
		if implementFile, err = utils.CreateFile(implementFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameService, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Create %s Service Interface File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), implementFileName))

			// 格式化写入
			if err = libService.NewFormatterServiceImplement().Format(cmdParams.Name, cmdParams.Version).WriteOut(implementFile); err != nil {
				utils.CommandLogger.Error(utils.CommandNameService, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Write %s Service Implement File Successful!", cmdParams.Name))
			}

		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameService, fmt.Sprintf("%s Service Implement File Is Exist!", cmdParams.Name))
	}

	// 生成服务数据传输对象
	dtoFileName := cmdParams.DOutputPath + "/service_" + cmdParams.Name + "_dto.go"
	if !IsServiceFileExist(dtoFileName) {
		if implementFile, err = utils.CreateFile(dtoFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameService, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Create %s Service DTO File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), dtoFileName))

			// 格式化写入
			if err = libService.NewFormatterServiceDto().Format(cmdParams.Name, cmdParams.Version).WriteOut(implementFile); err != nil {
				utils.CommandLogger.Error(utils.CommandNameService, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameService, fmt.Sprintf("Write %s Service DTO File Successful!", cmdParams.Name))
			}

		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameService, fmt.Sprintf("%s Service DTO File Is Exist!", cmdParams.Name))
	}

	return nil
}

func IsServiceFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

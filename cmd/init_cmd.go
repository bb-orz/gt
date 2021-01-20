package cmd

import (
	"errors"
	"fmt"
	"github.com/bb-orz/gt/libs/libInit"
	"github.com/bb-orz/gt/utils"
	"github.com/urfave/cli/v2"
	"os"
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Usage:       "Go Web Application Initialization",
		UsageText:   "gt init [--name|-n=][project_name] [--git|-g=true|false] [--mod|-m=true|false] ",
		Description: "The init command create a new go web application in current directory，this command will generate some necessary folders and files, create a project.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Value:   "goapp",
				Usage:   "[[--name|-n=]ProjectName]",
			},
			&cli.StringFlag{
				Name:    "sample",
				Aliases: []string{"s"},
				Value:   "sample",
				Usage:   "[--sample|-s=[sample|account|grpc|micro]]",
			},
			&cli.BoolFlag{
				Name:    "mod",
				Aliases: []string{"m"},
				Value:   true,
				Usage:   "[--mod|-m=true|false]",
			},
			&cli.BoolFlag{
				Name:    "git",
				Aliases: []string{"g"},
				Value:   true,
				Usage:   "[--git|-g=true|false]",
			},
		},
		Action: InitCommandAction,
	}
}

func InitCommandAction(ctx *cli.Context) error {
	var err error
	var fileInfo os.FileInfo
	var nameFlag string
	var sampleFlag string
	var modFlag bool
	var gitFlag bool

	if ctx.NArg() == 1 {
		nameFlag = ctx.Args().Get(0)
	} else {
		nameFlag = ctx.String("name")
	}

	utils.CommandLogger.Info(utils.CommandNameInit, "Directory Checking ... ")
	// 检查当前目录是否有权限读写
	if err = utils.CheckDirMode(); err != nil {
		utils.CommandLogger.Error(utils.CommandNameInit, err)
		return nil
	}
	utils.CommandLogger.OK(utils.CommandNameInit, "Current directory is readable and writable. ")

	// 检查项目目录是否已存在
	fileInfo, err = os.Stat(nameFlag)
	if err == nil && fileInfo.IsDir() {
		utils.CommandLogger.Error(utils.CommandNameInit, errors.New("Directory is existed, please try again! "))
		return nil
	}

	// 拉取脚手架模板
	sampleFlag = ctx.String("sample")
	utils.CommandLogger.Info(utils.CommandNameInit, "Pull Scaffold Sample ... ")
	switch sampleFlag {
	case "sample":
		if err := libInit.GitCloneSample(nameFlag); err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameInit, "git clone sample scaffold successful!")
	case "account":
		if err := libInit.GitCloneAccount(nameFlag); err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameInit, "git clone account scaffold successful!")
	case "grpc":
		if err := libInit.GitCloneGrpc(nameFlag); err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameInit, "git clone grpc scaffold successful!")
	case "micro":
		if err := libInit.GitCloneMicro(nameFlag); err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameInit, "git clone go-micro scaffold successful!")
	}

	// go mod
	modFlag = ctx.Bool("mod")
	utils.CommandLogger.Info(utils.CommandNameInit, fmt.Sprintf("Mod:%v", modFlag))
	if modFlag {
		if version, b := libInit.CheckGoMod(); b {
			utils.CommandLogger.OK(utils.CommandNameInit, fmt.Sprintf("Current Version:%s, enable to go mod! ", version))
		} else {
			utils.CommandLogger.Error(utils.CommandNameInit, errors.New(fmt.Sprintf("Current Version:%s, not enable to go mod! ", version)))
		}

		// 获取完整目录
		pwd, err := os.Getwd()
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		}

		if nameFlag != "goapp" {
			err = utils.ReplaceMainPackageNAme(pwd, nameFlag)
			if err != nil {
				utils.CommandLogger.Error(utils.CommandNameInit, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameInit, fmt.Sprintf("Replace main package as %s successful", nameFlag))
			}
		}

		tidyCmd := `cd ` + nameFlag + ` && go mod tidy`
		err = utils.ExecShellCommand(tidyCmd)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameInit, "Go mod tidy successful")
		}
	}

	// 本地 git 初始化
	gitFlag = ctx.Bool("git")
	utils.CommandLogger.Info(utils.CommandNameInit, fmt.Sprintf("Git:%v", gitFlag))
	if gitFlag {
		utils.CommandLogger.Info(utils.CommandNameInit, "Git init ... ")
		InitGitCmd := "cd " + nameFlag + " && git init && git checkout -b dev && git remote rm origin"
		err := utils.ExecShellCommand(InitGitCmd)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameInit, err)
		}
		utils.CommandLogger.OK(utils.CommandNameInit, "Reinitialized git repository successful!")
	}

	return nil
}

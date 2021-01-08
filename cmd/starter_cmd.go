package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/libs/libStarter"
	"gt/utils"
	"io"
)

func StarterCommand() *cli.Command {
	return &cli.Command{
		Name:        "starter",
		Usage:       "Add Goinfras Starter",
		UsageText:   "gt starter [--name|-n=][StarterName]",
		Description: "The starter command create a new starter base on goinfras ，this command will generate some necessary files or dir in starter directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "example"},
			&cli.StringFlag{Name: "output_path", Aliases: []string{"o"}, Value: "./starter"},
		},
		Action: StarterCommandAction,
	}
}

func StarterCommandAction(ctx *cli.Context) error {
	var err error
	var cmdParams = libStarter.CmdParams{
		Name:       ctx.String("name"),
		OutputPath: ctx.String("output_path"),
	}

	// 创建starter/config/test/readme/x
	// create service,server,client file
	var starterFileWriter, configFileWriter, testingFileWriter, readmeFileWriter, xFileWriter io.Writer
	var starterFileName, configFileName, testingFileName, readmeFileName, xFileName string

	// 创建 starter 文件
	starterFileName = cmdParams.OutputPath + "/" + cmdParams.Name + "/starter.go"
	if !IsServiceFileExist(starterFileName) {
		if starterFileWriter, err = utils.CreateFile(starterFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameStarter, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Create %s Starter File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), starterFileName))

			// 格式化写入
			if err = libStarter.NewFormatterStarter().Format(&cmdParams).WriteOut(starterFileWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameStarter, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Write %s Starter File Successful!", utils.CamelString(cmdParams.Name)))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameStarter, fmt.Sprintf("%s Starter File Is Exist!", utils.CamelString(cmdParams.Name)))
		return nil
	}

	// 创建 starter config 文件
	configFileName = cmdParams.OutputPath + "/" + cmdParams.Name + "/config.go"
	if !IsServiceFileExist(configFileName) {
		if configFileWriter, err = utils.CreateFile(configFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameStarter, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Create %s Starter Config File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), configFileName))

			// 格式化写入
			if err = libStarter.NewFormatterStarterConfig().Format(&cmdParams).WriteOut(configFileWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameStarter, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Write %s Starter Config File Successful!", utils.CamelString(cmdParams.Name)))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameStarter, fmt.Sprintf("%s Starter Config File Is Exist!", utils.CamelString(cmdParams.Name)))
		return nil
	}

	// 创建 starter testing 文件
	testingFileName = cmdParams.OutputPath + "/" + cmdParams.Name + "/run_test.go"
	if !IsServiceFileExist(testingFileName) {
		if testingFileWriter, err = utils.CreateFile(testingFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameStarter, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Create %s Starter Testing File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), testingFileName))

			// 格式化写入
			if err = libStarter.NewFormatterStarterTesting().Format(&cmdParams).WriteOut(testingFileWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameStarter, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Write %s Starter Testing File Successful!", utils.CamelString(cmdParams.Name)))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameStarter, fmt.Sprintf("%s Starter Testing File Is Exist!", utils.CamelString(cmdParams.Name)))
		return nil
	}

	// 创建 starter x 文件
	xFileName = cmdParams.OutputPath + "/" + cmdParams.Name + "/x.go"
	if !IsServiceFileExist(xFileName) {
		if xFileWriter, err = utils.CreateFile(xFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameStarter, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Create %s Starter X Instence File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), xFileName))

			// 格式化写入
			if err = libStarter.NewFormatterStarterX().Format(&cmdParams).WriteOut(xFileWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameStarter, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Write %s Starter X Instence File Successful!", utils.CamelString(cmdParams.Name)))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameStarter, fmt.Sprintf("%s Starter X Instence File Is Exist!", utils.CamelString(cmdParams.Name)))
		return nil
	}

	// 创建 starter readme 文件
	readmeFileName = cmdParams.OutputPath + "/" + cmdParams.Name + "/README.md"
	if !IsServiceFileExist(readmeFileName) {
		if readmeFileWriter, err = utils.CreateFile(readmeFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameStarter, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Create %s Starter README.md File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), readmeFileName))

			// 格式化写入
			if err = libStarter.NewFormatterStarterReadme().Format(&cmdParams).WriteOut(readmeFileWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameStarter, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameStarter, fmt.Sprintf("Write %s Starter README.md File Successful!", utils.CamelString(cmdParams.Name)))
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameStarter, fmt.Sprintf("%s Starter README.md File Is Exist!", utils.CamelString(cmdParams.Name)))
		return nil
	}

	return nil
}

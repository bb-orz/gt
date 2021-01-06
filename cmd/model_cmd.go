package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/libs/libModel"
	"gt/utils"
)

func ModelCommand() *cli.Command {
	return &cli.Command{
		Name:        "model",
		Usage:       "Add core model",
		Description: "The model command create a new core model with go struct，this command will generate some necessary files or dir in core directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "driver", Aliases: []string{"D"}, Value: "mysql"},
			&cli.StringFlag{Name: "host", Aliases: []string{"H"}, Value: "localhost"},
			&cli.IntFlag{Name: "port", Aliases: []string{"P"}, Value: 3306},
			&cli.StringFlag{Name: "database", Aliases: []string{"d"}, Required: true},
			&cli.StringFlag{Name: "table", Aliases: []string{"t"}, Required: true},
			&cli.StringFlag{Name: "user", Aliases: []string{"u"}, Value: "dev"},
			&cli.StringFlag{Name: "password", Aliases: []string{"p"}, Value: "123456"},
			&cli.StringFlag{Name: "output_path", Aliases: []string{"o"}, Value: "./core/"},
			&cli.StringFlag{Name: "formatter", Aliases: []string{"f"}, Value: "gorm"},
		},
		Action: ModelCommandAction,
	}
}

func ModelCommandAction(ctx *cli.Context) error {
	cmdParams := &libModel.CmdParams{
		Driver:     ctx.String("driver"),
		Host:       ctx.String("host"),
		Port:       ctx.Int("port"),
		DbName:     ctx.String("database"),
		Table:      ctx.String("table"),
		User:       ctx.String("user"),
		Password:   ctx.String("password"),
		OutputPath: ctx.String("output_path"),
		Formatter:  ctx.String("formatter"),
	}

	// 获取db连接实例
	db, err := libModel.GetDBInstance(cmdParams)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}

	// 获取表结构
	columns, err := libModel.GetTableSchema(db, cmdParams.DbName, cmdParams.Table)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}

	// 使用表结构生成go代码
	switch cmdParams.Formatter {
	case "gorm":
		// 创建输出文件
		fileName := cmdParams.OutputPath + cmdParams.Table + "/" + cmdParams.Table + "_model.go"
		writer, err := utils.CreateFile(fileName)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Create Gorm %s Model Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), fileName))
		}

		// 格式化输出
		if err = libModel.NewFormatterGorm().Format(cmdParams.Table, columns).WriteOut(writer); err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameModel, "Write Gorm %s Model Successful!")
		}
	case "sqlbuilder":
		// 创建输出文件
		fileName := cmdParams.OutputPath + cmdParams.Table + "/" + cmdParams.Table + "_model.go"
		writer, err := utils.CreateFile(fileName)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Create SqlBuilder %s Model Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), fileName))
		}

		// 格式化输出
		if err = libModel.NewFormatterSqlBuilder().Format(cmdParams.Table, columns).WriteOut(writer); err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameModel, "Write SqlBuilder %s Model Successful!")
		}
	}

	// 生成model相关DTO
	// 创建输出文件
	dtoFileName := "./dtos/" + cmdParams.Table + "_dto.go"
	dtoWriter, err := utils.CreateFile(dtoFileName)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	} else {
		utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Create %s DTO Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), dtoFileName))
	}

	// 格式化输出
	if err = libModel.NewFormatterDTOStruct().Format(cmdParams.Table, columns).WriteOut(dtoWriter); err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	} else {
		utils.CommandLogger.OK(utils.CommandNameModel, "Write %s DTO Successful!")
	}

	return nil
}

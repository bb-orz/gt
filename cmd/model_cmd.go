package cmd

import (
	"fmt"
	"github.com/bb-orz/gt/libs/libModel"
	"github.com/bb-orz/gt/utils"
	"github.com/urfave/cli/v2"
)

func ModelCommand() *cli.Command {
	return &cli.Command{
		Name:        "model",
		Usage:       "Add core model",
		Description: "The model command create a new core model with go struct，this command will generate some necessary files or dir in core directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
			&cli.StringFlag{Name: "driver", Aliases: []string{"D"}, Value: "mysql"},
			&cli.StringFlag{Name: "host", Aliases: []string{"H"}, Value: "localhost"},
			&cli.IntFlag{Name: "port", Aliases: []string{"P"}, Value: 3306},
			&cli.StringFlag{Name: "database", Aliases: []string{"d"}, Required: true},
			&cli.StringFlag{Name: "table", Aliases: []string{"t"}, Required: true},
			&cli.StringFlag{Name: "user", Aliases: []string{"u"}, Value: "dev"},
			&cli.StringFlag{Name: "password", Aliases: []string{"p"}, Value: "123456"},
			&cli.StringFlag{Name: "output_path", Aliases: []string{"o"}, Value: "./core"},
			&cli.StringFlag{Name: "dto_output_path", Aliases: []string{"O"}, Value: "./dtos"},
			&cli.StringFlag{Name: "formatter", Aliases: []string{"f"}, Value: "gorm"},
		},
		Action: ModelCommandAction,
	}
}

func ModelCommandAction(ctx *cli.Context) error {
	cmdParams := &libModel.CmdParams{
		Name:        ctx.String("name"),
		Driver:      ctx.String("driver"),
		Host:        ctx.String("host"),
		Port:        ctx.Int("port"),
		DbName:      ctx.String("database"),
		Table:       ctx.String("table"),
		User:        ctx.String("user"),
		Password:    ctx.String("password"),
		OutputPath:  ctx.String("output_path"),
		DOutputPath: ctx.String("dto_output_path"),
		Formatter:   ctx.String("formatter"),
	}

	if cmdParams.Name == "" {
		cmdParams.Name = cmdParams.Table
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
		fileName := cmdParams.OutputPath + "/" + cmdParams.Name + "/" + cmdParams.Table + "_model.go"
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
		fileName := cmdParams.OutputPath + "/" + cmdParams.Name + "/" + cmdParams.Table + "_model.go"
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

	// 生成Model相关DAO
	daoFileName := cmdParams.OutputPath + "/" + cmdParams.Name + "/" + cmdParams.Table + "_dao.go"
	if daoWriter, err := utils.CreateFile(daoFileName); err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	} else {
		utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Create %s Model Dao File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Table), daoFileName))

		switch cmdParams.Formatter {
		case "gorm":
			if err = libModel.NewFormatterGormDao().Format(cmdParams.Table, nil).WriteOut(daoWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameModel, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Write %s GORM Dao File Successful!", utils.CamelString(cmdParams.Table)))
			}
		case "sqlbuilder":
			if err = libModel.NewFormatterSqlBuilderDao().Format(cmdParams.Table, nil).WriteOut(daoWriter); err != nil {
				utils.CommandLogger.Error(utils.CommandNameModel, err)
				return nil
			} else {
				utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Write %s SqlBuilder Dao File Successful!", utils.CamelString(cmdParams.Table)))
			}
		}
	}

	// 生成model相关DTO
	// 创建输出文件
	dtoFileName := cmdParams.DOutputPath + "/" + cmdParams.Table + "_dto.go"
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

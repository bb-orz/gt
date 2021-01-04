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
		Action: modelCommandFunc,
	}
}

func modelCommandFunc(ctx *cli.Context) error {
	cmdCfg := &libModel.CmdCfg{
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
	db, err := libModel.GetDBInstance(cmdCfg)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}

	// 把数据库连接参数写进setting配置

	// 获取表结构
	columns, err := libModel.GetTableSchema(db, cmdCfg.DbName, cmdCfg.Table)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}

	// 使用表结构生成go代码
	switch cmdCfg.Formatter {
	case "gorm":
		// 创建输出文件
		fileName := cmdCfg.OutputPath + cmdCfg.Table + "/" + cmdCfg.Table + "_model.go"
		writer, err := utils.CreateFile(fileName)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		}

		// 格式化输出
		err = libModel.NewFormatterGormStruct().Format(cmdCfg.Table, columns).WriteOut(writer)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Generate Gorm %s Model Successful! >>> FilePath：%s", utils.CamelString(cmdCfg.Table), fileName))
	case "sqlbuilder":
		// 创建输出文件
		fileName := cmdCfg.OutputPath + cmdCfg.Table + "/" + cmdCfg.Table + "_model.go"
		writer, err := utils.CreateFile(fileName)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		}

		// 格式化输出
		err = libModel.NewFormatterSqlBuilderStruct().Format(cmdCfg.Table, columns).WriteOut(writer)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameModel, err)
			return nil
		}
		utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Generate SqlBuilder %s Model Successful! >>> FilePath：%s", utils.CamelString(cmdCfg.Table), fileName))
	}

	// 生成model相关DTO
	// 创建输出文件
	dtoFileName := "./dtos/" + cmdCfg.Table + "_dto.go"
	dtoWriter, err := utils.CreateFile(dtoFileName)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}

	// 格式化输出
	err = libModel.NewFormatterDTOStruct().Format(cmdCfg.Table, columns).WriteOut(dtoWriter)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
		return nil
	}
	utils.CommandLogger.OK(utils.CommandNameModel, fmt.Sprintf("Generate %s DTO Successful! >>> FilePath：%s", utils.CamelString(cmdCfg.Table), dtoFileName))

	return nil
}

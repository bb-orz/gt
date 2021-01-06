package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/libs/libDomain"
	"gt/utils"
)

func DomainCommand() *cli.Command {
	return &cli.Command{
		Name:        "domain",
		Usage:       "Add core domain in project",
		UsageText:   "gt domain [--name|-n=][DomainName]",
		Description: "The domain command create a new core domain with go struct，this command will generate some necessary files or dir in core directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "example", Required: true},
			&cli.StringFlag{Name: "driver", Aliases: []string{"D"}, Value: "mysql"},
			&cli.StringFlag{Name: "host", Aliases: []string{"H"}, Value: "localhost"},
			&cli.IntFlag{Name: "port", Aliases: []string{"P"}, Value: 3306},
			&cli.StringFlag{Name: "database", Aliases: []string{"d"}},
			&cli.StringFlag{Name: "table", Aliases: []string{"t"}},
			&cli.StringFlag{Name: "user", Aliases: []string{"u"}, Value: "dev"},
			&cli.StringFlag{Name: "password", Aliases: []string{"p"}, Value: "123456"},
			&cli.StringFlag{Name: "output_path", Aliases: []string{"o"}, Value: "./core"},
			&cli.StringFlag{Name: "dto_output_path", Aliases: []string{"O"}, Value: "./dtos"},
			&cli.StringFlag{Name: "formatter", Aliases: []string{"f"}, Value: "gorm"},
		},
		Action: DomainCommandAction,
	}
}

func DomainCommandAction(ctx *cli.Context) error {
	// 接收领域名参数，创建领域代码文件
	cmdParams := &libDomain.CmdParams{
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

	domainFile, daoFile, testFile, err := libDomain.CreateDomainFile(cmdParams)

	// 写入domain/dao/test格式化代码模板
	if err = libDomain.NewFormatterDomain().Format(cmdParams.Name).WriteOut(domainFile); err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
		return nil
	} else {
		utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Write %s Domain File Successful!", cmdParams.Name))
	}

	if err = libDomain.NewFormatterDomainTesting().Format(cmdParams.Name).WriteOut(testFile); err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
		return nil
	} else {
		utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Write %s Domain Testing File Successful!", cmdParams.Name))
	}

	switch cmdParams.Formatter {
	case "gorm":
		if err = libDomain.NewFormatterDomainGormDao().Format(cmdParams.Name).WriteOut(daoFile); err != nil {
			utils.CommandLogger.Error(utils.CommandNameDomain, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Write %s Domain Gorm Dao File Successful!", cmdParams.Name))
		}
	case "sqlbuilder":
		if err = libDomain.NewFormatterDomainSqlBuilderDao().Format(cmdParams.Name).WriteOut(daoFile); err != nil {
			utils.CommandLogger.Error(utils.CommandNameDomain, err)
			return nil
		} else {
			utils.CommandLogger.OK(utils.CommandNameDomain, fmt.Sprintf("Write %s Domain SqlBuilder Dao File Successful!", cmdParams.Name))
		}
	}

	// 如有传递数据库连接参数，生成相应的model文件及常用的curd dao代码
	if cmdParams.DbName != "" && cmdParams.Table != "" {
		// 调用model命令
		err := ModelCommandAction(ctx)
		if err != nil {
			utils.CommandLogger.Error(utils.CommandNameDomain, err)
			return nil
		}
	}

	return nil
}

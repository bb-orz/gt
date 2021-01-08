package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gt/libs/libRpc"
	"gt/utils"
	"io"
)

func RPCCommand() *cli.Command {
	return &cli.Command{
		Name:        "rpc",
		Usage:       "Add RPC Service",
		UsageText:   "gt rpc [--name|-n=][RPCName]",
		Description: "The rpc command create a new rpc service with go struct，this command will generate some necessary files or dir in rpc directory .",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "example"},
			&cli.StringFlag{Name: "rpc_type", Aliases: []string{"t"}, Value: "grpc", Usage: "[-t grpc|micro]"},
			&cli.StringFlag{Name: "protobuf_path", Aliases: []string{"p"}, Value: "./protobuf"},
			&cli.StringFlag{Name: "server_output_path", Aliases: []string{"s"}, Value: "./starter"},
			&cli.StringFlag{Name: "client_output_path", Aliases: []string{"c"}, Value: "./clients"},
			&cli.StringFlag{Name: "proto_gen_output_path", Aliases: []string{"P"}, Value: "./protobuf"},
			&cli.StringFlag{Name: "service_output_path", Aliases: []string{"S"}, Value: "./services"},
		},
		Action: RPCCommandAction,
	}
}

func RPCCommandAction(ctx *cli.Context) error {
	var err error
	var cmdParams = libRpc.CmdParams{
		Name:               ctx.String("name"),
		Type:               ctx.String("rpc_type"),
		ProtoBufPath:       ctx.String("protobuf_path"),
		ClientOutputPath:   ctx.String("client_output_path"),
		ProtoGenOutputPath: ctx.String("proto_gen_output_path"),
		ServiceOutputPath:  ctx.String("service_output_path"),
		ServerOutputPath:   ctx.String("server_output_path"),
	}

	// Gen protobuf go file
	var shellCmd string
	switch cmdParams.Type {
	case "grpc":
		shellCmd = "cd " + cmdParams.ProtoBufPath + " && protoc --proto_path= --go_out=plugins=grpc:. ./" + cmdParams.Name + "_" + cmdParams.Type + ".proto"
	case "micro":
		shellCmd = "cd " + cmdParams.ProtoBufPath + " && protoc --proto_path= --micro_out=. --go_out=. ./" + cmdParams.Name + "_" + cmdParams.Type + ".proto"
	}

	if err = utils.ExecShellCommand(shellCmd); err != nil {
		utils.CommandLogger.Error(utils.CommandNameRpc, err)
		return nil
	}

	// create service server client file
	var serviceFileWriter, serverFileWriter, clientFileWriter io.Writer

	// 创建 rpc service 文件
	serviceFileName := cmdParams.ServiceOutputPath + "/" + cmdParams.Name + "_" + cmdParams.Type + "_service.go"
	if !IsServiceFileExist(serviceFileName) {
		if serviceFileWriter, err = utils.CreateFile(serviceFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameRpc, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Create %s %s Service File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type), serviceFileName))

			switch cmdParams.Type {
			case "grpc":
				// 格式化写入
				if err = libRpc.NewFormatterGrpcService().Format(&cmdParams).WriteOut(serviceFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Service File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}
			case "micro":
				// 格式化写入
				if err = libRpc.NewFormatterMicroService().Format(&cmdParams).WriteOut(serviceFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Service File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}

			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameRpc, fmt.Sprintf("%s %s Service File Is Exist!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
	}

	// 创建 rpc client 文件
	clientFileName := cmdParams.ClientOutputPath + "/" + cmdParams.Name + "_" + cmdParams.Type + "_client.go"
	if !IsServiceFileExist(clientFileName) {
		if clientFileWriter, err = utils.CreateFile(clientFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameRpc, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Create %s %s Client File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type), clientFileName))

			switch cmdParams.Type {
			case "grpc":
				// 格式化写入
				if err = libRpc.NewFormatterGrpcClient().Format(&cmdParams).WriteOut(clientFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Client File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}
			case "micro":
				if err = libRpc.NewFormatterMicroClient().Format(&cmdParams).WriteOut(clientFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Client File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameRpc, fmt.Sprintf("%s %s Client File Is Exist!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
	}

	// 创建 rpc server starter
	serverFileName := cmdParams.ServerOutputPath + "/" + cmdParams.Name + utils.CamelString(cmdParams.Type) + "/starter.go"
	if !IsServiceFileExist(serverFileName) {
		if serverFileWriter, err = utils.CreateFile(serverFileName); err != nil {
			utils.CommandLogger.Error(utils.CommandNameRpc, err)
		} else {
			utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Create %s %s Server Starter File Successful! >>> FilePath：%s", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type), serverFileName))

			switch cmdParams.Type {
			case "grpc":
				// 格式化写入
				if err = libRpc.NewFormatterGrpcServer().Format(&cmdParams).WriteOut(serverFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Server Starter File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}
			case "micro":
				// 格式化写入
				if err = libRpc.NewFormatterMicroServer().Format(&cmdParams).WriteOut(serverFileWriter); err != nil {
					utils.CommandLogger.Error(utils.CommandNameRpc, err)
					return nil
				} else {
					utils.CommandLogger.OK(utils.CommandNameRpc, fmt.Sprintf("Write %s %s Server Starter File Successful!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
				}
			}
		}
	} else {
		utils.CommandLogger.Warning(utils.CommandNameRpc, fmt.Sprintf("%s %s Server Starter File Is Exist!", utils.CamelString(cmdParams.Name), utils.CamelString(cmdParams.Type)))
	}

	utils.CommandLogger.Info(utils.CommandNameRpc, fmt.Sprintf("Please implement protobuf service interface in %s", serverFileName))
	utils.CommandLogger.Info(utils.CommandNameRpc, "Please register starter rpc server in app/register.go")

	return nil
}

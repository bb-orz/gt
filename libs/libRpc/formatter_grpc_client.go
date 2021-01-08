package libRpc

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterGppcClient() *FormatterGrpcClient {
	return new(FormatterGrpcClient)
}

type FormatterGrpcClient struct {
	FormatterStruct
}

func (f *FormatterGrpcClient) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = utils.GetLastPath(cmdParams.ClientOutputPath)
	f.StructName = utils.CamelString(cmdParams.Name)
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["grpc"] = ImportItem{Alias: "", Package: "google.golang.org/grpc"}

	return f
}

func (f *FormatterGrpcClient) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GrpcClientTemplate").Parse(GrpcClientCodeTemplate)).Execute(writer, *f)
}

const GrpcClientCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
	"goapp/protobuf/generate/{{ .Name }}Pb"
	"goapp/starter/{{ .Name }}{{ .TypeName }}"
)


func Get{{ .StructName }}{{ .TypeName }}Client() ({{ .Name }}Pb.{{ .StructName }}{{ .TypeName }}ServiceClient, error) {
	// 检查服务
	var err error
	var conn *grpc.ClientConn
	// 获得一个grpc服务端的连接
	if conn, err = grpc.Dial(fmt.Sprintf("%s:%d", {{ .Name }}{{ .TypeName }}.Cfg.ServerHost, {{ .Name }}{{ .TypeName }}.Cfg.ServerPort), grpc.WithInsecure()); err != nil {
		return nil, err
	}

	// 创建客户端
	client := {{ .Name }}Pb.New{{ .StructName }}{{ .TypeName }}ServiceClient(conn)
	return client, nil
}

`

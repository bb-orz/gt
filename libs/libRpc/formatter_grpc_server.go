package libRpc

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterGrpcServer() *FormatterGrpcServer {
	return new(FormatterGrpcServer)
}

type FormatterGrpcServer struct {
	FormatterStruct
}

func (f *FormatterGrpcServer) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name + "Grpc"
	f.StructName = utils.CamelString(cmdParams.Name)
	f.ImportList = make(map[string]ImportItem)
	f.ImportList[""] = ImportItem{Alias: "", Package: ""}
	return f
}

func (f *FormatterGrpcServer) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GrpcServerTemplate").Parse(GrpcServerCodeTemplate)).Execute(writer, *f)
}

const GrpcServerCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

`

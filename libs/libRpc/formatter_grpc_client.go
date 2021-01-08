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
	f.ImportList = make(map[string]ImportItem)
	f.ImportList[""] = ImportItem{Alias: "", Package: ""}

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
)

`

package libRpc

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterGrpcService() *FormatterGrpcService {
	return new(FormatterGrpcService)
}

type FormatterGrpcService struct {
	FormatterStruct
}

func (f *FormatterGrpcService) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = utils.GetLastPath(cmdParams.ServiceOutputPath)
	f.StructName = utils.CamelString(cmdParams.Name)
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.ImportList = make(map[string]ImportItem)
	f.ImportList[""] = ImportItem{Alias: "", Package: ""}
	return f
}

func (f *FormatterGrpcService) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GrpcServiceTemplate").Parse(GrpcServiceCodeTemplate)).Execute(writer, *f)
}

const GrpcServiceCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

// TODO 请实现 {{ .PackageName }}Pb.{{ .StructName }}{{ .TypeName }}ServiceServer 接口
type {{ .StructName }}{{ .TypeName }}Service struct{}

`

package libRpc

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterMicroService() *FormatterMicroService {
	return new(FormatterMicroService)
}

type FormatterMicroService struct {
	FormatterStruct
}

func (f *FormatterMicroService) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = utils.GetLastPath(cmdParams.ServiceOutputPath)
	f.StructName = utils.CamelString(cmdParams.Name)
	f.Name = cmdParams.Name
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.ImportList = make(map[string]ImportItem)
	return f
}

func (f *FormatterMicroService) WriteOut(writer io.Writer) error {
	return template.Must(template.New("MicroServiceTemplate").Parse(MicroServiceCodeTemplate)).Execute(writer, *f)
}

const MicroServiceCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

// TODO 请实现 protobuf 协议生成的 {{ .Name }}Pb.{{ .StructName }}{{ .TypeName }}ServiceHandler 接口
type {{ .StructName }}{{ .TypeName }}Service struct{}

`

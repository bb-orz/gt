package libRpc

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterMicroClient() *FormatterMicroClient {
	return new(FormatterMicroClient)
}

type FormatterMicroClient struct {
	FormatterStruct
}

func (f *FormatterMicroClient) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = utils.GetLastPath(cmdParams.ClientOutputPath)
	f.StructName = utils.CamelString(cmdParams.Name)
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["micro"] = ImportItem{Alias: "", Package: "github.com/micro/go-micro"}

	return f
}

func (f *FormatterMicroClient) WriteOut(writer io.Writer) error {
	return template.Must(template.New("MicroClientTemplate").Parse(MicroClientCodeTemplate)).Execute(writer, *f)
}

const MicroClientCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
	"goapp/protobuf/generate/{{ .Name }}Pb"
	"goapp/starter/{{ .Name }}{{ .TypeName }}"
)

func Get{{ .StructName }}{{ .TypeName }}Client() ({{ .Name }}Pb.{{ .StructName }}{{ .TypeName }}Service) {
	service := micro.NewService (
		micro.Name({{ .Name }}{{ .TypeName }}.Cfg.Name),
		micro.Address(fmt.Sprintf("%s:%d", {{ .Name }}{{ .TypeName }}.Cfg.ServerHost, {{ .Name }}{{ .TypeName }}.Cfg.ServerPort)),
	)

	service.Init()

	return {{ .Name }}Pb.New{{ .StructName }}{{ .TypeName }}Service({{ .Name }}{{ .TypeName }}.Cfg.Name, service.Client())
}


`

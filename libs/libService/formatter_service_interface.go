package libService

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterServiceInterface() *FormatterServiceInterface {
	return new(FormatterServiceInterface)
}

type FormatterServiceInterface struct {
	FormatterStruct
}

func (f *FormatterServiceInterface) Format(name, version string) IFormatter {
	f.PackageName = "services"
	f.Name = name
	f.StructName = utils.CamelString(name)
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}

	return f
}

func (f *FormatterServiceInterface) WriteOut(writer io.Writer) error {
	return template.Must(template.New("ServiceInterfaceTemplate").Parse(ServiceInterfaceCodeTemplate)).Execute(writer, *f)
}

const ServiceInterfaceCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

var {{ .Name }}Service I{{ .StructName }}Service

func Set{{ .StructName }}Service(sv I{{ .StructName }}Service) {
	{{ .Name }}Service = sv
}

func Get{{ .StructName }}Service() I{{ .StructName }}Service{
	return {{ .Name }}Service
}

type I{{ .StructName }}Service interface {
	Foo(i dtos.FooDTO) error
	Bar(i dtos.BarDTO) error
}


`

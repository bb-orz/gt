package libService

import (
	"gt/utils"
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

var service{{ .StructName }} IService{{ .StructName }}

func SetService{{ .StructName }}(sv IService{{ .StructName }}) {
	service{{ .StructName }} = sv
}

func GetService{{ .StructName }}() IService{{ .StructName }} {
	return service{{ .StructName }}
}

type IService{{ .StructName }} interface {
	Foo(i dtos.FooDTO) error
	Bar(i dtos.BarDTO) error
}


`

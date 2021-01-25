package libService

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterServiceImplement() *FormatterServiceImplement {
	return new(FormatterServiceImplement)
}

type FormatterServiceImplement struct {
	FormatterStruct
}

func (f *FormatterServiceImplement) Format(name, version string) IFormatter {
	f.PackageName = "services"
	f.Name = name
	f.StructName = utils.CamelString(name)
	f.Version = version
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["sync"] = ImportItem{Alias: "", Package: "sync"}
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}

	return f
}

func (f *FormatterServiceImplement) WriteOut(writer io.Writer) error {
	return template.Must(template.New("ServiceImplementTemplate").Parse(ServiceImplementCodeTemplate)).Execute(writer, *f)
}

const ServiceImplementCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

/*
实现服务逻辑 Version:{{ .Version }}
*/


var _ I{{ .StructName }}Service = new({{ .StructName }}Service{{ .Version }})

func init() {
	var once sync.Once
	once.Do(func() {
		Set{{ .StructName }}Service(new({{ .StructName }}Service{{ .Version }}))
	})
}

type {{ .StructName }}Service{{ .Version }} struct{}


func (s *{{ .StructName }}Service{{ .Version }}) Foo(i dtos.FooDTO) error {
	var err error
	// TODO Call Domain
	fmt.Println("Service {{ .StructName }}{{ .Version }}.Foo")

	return err
}


func (s *{{ .StructName }}Service{{ .Version }}) Bar(i dtos.BarDTO) error {
	var err error
	// TODO  Call Domain
	fmt.Println("Service {{ .StructName }}{{ .Version }}.Bar")
	return err
}

`

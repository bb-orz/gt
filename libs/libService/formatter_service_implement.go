package libService

import (
	"gt/utils"
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
	f.PackageName = "core"
	f.StructName = utils.CamelString(name)
	f.Version = version
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["services"] = ImportItem{Alias: "", Package: "goapp/services"}
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


var _ services.IService{{ .StructName }} = new(Service{{ .StructName }}{{ .Version }})

func init() {
	var once sync.Once
	once.Do(func() {
		services.SetService{{ .StructName }}(new(Service{{ .StructName }}{{ .Version }}))
	})
}

type Service{{ .StructName }}{{ .Version }} struct{}


func (s *Service{{ .StructName }}{{ .Version }}) Foo(i dtos.FooDTO) error {
	var err error
	// TODO Call Domain
	fmt.Println("Service {{ .StructName }}{{ .Version }}.Foo")

	return err
}


func (s *Service{{ .StructName }}{{ .Version }}) Bar(i dtos.BarDTO) error {
	var err error
	// TODO  Call Domain
	fmt.Println("Service {{ .StructName }}{{ .Version }}.Bar")
	return err
}

`

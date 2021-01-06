package libService

import (
	"fmt"
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterServiceDto() *FormatterServiceDto {
	return new(FormatterServiceDto)
}

type FormatterServiceDto struct {
	FormatterStruct
}

func (f *FormatterServiceDto) Format(name, version string) IFormatter {
	f.PackageName = "dtos"
	f.StructName = utils.CamelString(name)
	f.FieldList = make([]Field, 0)
	exampleFile := Field{Name: "Id", Type: "uint", Comment: "示例参数"}
	exampleFile.StructTag = fmt.Sprintf("`validate:\"required,numeric\" json:\"%s\"`", name)
	f.FieldList = append(f.FieldList, exampleFile)

	return f
}

func (f *FormatterServiceDto) WriteOut(writer io.Writer) error {
	return template.Must(template.New("ServiceDTOTemplate").Parse(ServiceDTOCodeTemplate)).Execute(writer, *f)
}

const ServiceDTOCodeTemplate = `package {{ .PackageName }}

// FooDTO is a example data transfer object for {{ .StructName }} Service
type FooDTO struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}


// BarDTO is a example data transfer object for {{ .StructName }} Service
type BarDTO struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}

`

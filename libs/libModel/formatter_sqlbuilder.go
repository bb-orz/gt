package libModel

import (
	"fmt"
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterSqlBuilder() *FormatterSqlBuilder {
	return new(FormatterSqlBuilder)
}

type FormatterSqlBuilder struct {
	FormatterStruct
}

func (f *FormatterSqlBuilder) Format(tableName string, cols []Column) IFormatter {
	f.PackageName = tableName
	f.ImportList = make(map[string]ImportItem)
	f.StructName = utils.CamelString(tableName)
	f.TableName = tableName
	f.FieldList = make([]Field, len(cols))

	for idx, col := range cols {
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		if colType == CTypeTime {
			f.ImportList["time"] = ImportItem{Alias: "", Package: "time"}
		}
		f.FieldList[idx] = Field{
			Name:      col.GetName(),
			Type:      colType,
			StructTag: fmt.Sprintf("`ddb:\"%s\" json:\"%s\"`", col.Name, col.Name),
			Comment:   col.GetComment(),
		}
	}
	return f
}

func (f *FormatterSqlBuilder) WriteOut(writer io.Writer) error {
	return template.Must(template.New("SqlBuilderTemplate").Parse(SqlBuilderStructCodeTemplate)).Execute(writer, *f)
}

const SqlBuilderStructCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

const {{ .StructName }}TableName = "{{ .TableName }}"

// {{ .StructName }} is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}} struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }}
{{- end}}
}

func New{{.StructName}}() *{{.StructName}} {
	return new({{.StructName}})
}

func (*{{ .StructName }}) TableName() string {
	return {{ .StructName }}TableName	 	// {{ .Comment }}
}


// To DTO
func (m *{{ .StructName }}) ToDTO() *dtos.{{ .StructName }}DTO {
	return &dtos.{{ .StructName }}DTO{
		{{- range .FieldList }}
			{{ .Name }} : m.{{ .Name }},
		{{- end}}
	}
}

// From DTO
func (m *{{ .StructName }}) FromDTO(dto *dtos.{{ .StructName }}DTO) {
	{{- range .FieldList }}
		m.{{ .Name }} = dto.{{ .Name }}
	{{- end}}
}


`

package libModel

import (
	"fmt"
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterSqlBuilderStruct() *FormatterSqlBuilderStruct {
	return new(FormatterSqlBuilderStruct)
}

type FormatterSqlBuilderStruct struct {
	FormatterStruct
}

func (f *FormatterSqlBuilderStruct) Format(tableName string, cols []Column) IFormatter {
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

func (f *FormatterSqlBuilderStruct) WriteOut(writer io.Writer) error {
	err := template.Must(template.New("SqlBuilderStruct").Parse(SqlBuilderStructCodeTemplate)).Execute(writer, *f)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameModel, err)
	}
	return nil
}

const SqlBuilderStructCodeTemplate = `
package {{ .PackageName }}

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

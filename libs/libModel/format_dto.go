package libModel

import (
	"fmt"
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterDTOStruct() *FormatterDTOStruct {
	return new(FormatterDTOStruct)
}

type FormatterDTOStruct struct {
	FormatterStruct
}

func (f *FormatterDTOStruct) Format(tableName string, cols []Column) IFormatter {
	f.PackageName = "dtos"
	f.ImportList = make(map[string]ImportItem)
	f.StructName = utils.CamelString(tableName) + "DTO"
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

		dtoType, err := col.GetDTOType()
		if nil != err {
			continue
		}
		field := Field{
			Name:    col.GetName(),
			Type:    colType,
			Comment: col.GetComment(),
		}

		field.StructTag = fmt.Sprintf("`validate:\"required,%s\" json:\"%s\"`", dtoType, col.Name)
		f.FieldList[idx] = field
	}
	return f
}

func (f *FormatterDTOStruct) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DTOTemplate").Parse(DTOStructCodeTemplate)).Execute(writer, *f)
}

const DTOStructCodeTemplate = `
package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

// {{ .StructName }} is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}} struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}

`

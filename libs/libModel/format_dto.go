package libModel

import (
	"fmt"
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterDTOStruct() *FormatterDTO {
	return new(FormatterDTO)
}

type FormatterDTO struct {
	FormatterStruct
}

func (f *FormatterDTO) Format(name, tableName string, cols []Column) IFormatter {
	f.PackageName = "dtos"
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
			f.ImportList["validate"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XValidate"}
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

		field.StructTag = fmt.Sprintf("`validate:\"%s\" json:\"%s\"`", dtoType, col.Name)
		f.FieldList[idx] = field
	}
	return f
}

func (f *FormatterDTO) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DTOTemplate").Parse(DTOCodeTemplate)).Execute(writer, *f)
}

const DTOCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

// {{ .StructName }}DTO is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}}DTO struct {
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}

func (dto *{{.StructName}}DTO) Validate() error {
	return XValidate.V(dto)
}

`

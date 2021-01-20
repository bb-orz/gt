package libModel

import (
	"fmt"
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterGorm() *FormatterGorm {
	return new(FormatterGorm)
}

type FormatterGorm struct {
	FormatterStruct
}

func (f *FormatterGorm) Format(name, table string, cols []Column) IFormatter {
	f.PackageName = name
	f.ImportList = make(map[string]ImportItem)
	f.StructName = utils.CamelString(table)
	f.TableName = table
	f.FieldList = make([]Field, len(cols))

	for idx, col := range cols {
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		if colType == CTypeTime {
			f.ImportList["time"] = ImportItem{Alias: "", Package: "time"}
			// f.ImportList["gorm"] = ImportItem{Alias: "", Package: "gorm.io/gorm"}
		}
		f.FieldList[idx] = Field{
			Name:      col.GetName(),
			Type:      colType,
			StructTag: fmt.Sprintf("`gorm:\"%s\" json:\"%s\"`", col.Name, col.Name),
			Comment:   col.GetComment(),
		}
	}
	return f
}

func (f *FormatterGorm) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GormTemplate").Parse(GormStructCodeTemplate)).Execute(writer, *f)
}

// TODO 完善Gorm Model模板，根据goapp_account的实践
const GormStructCodeTemplate = `package {{ .PackageName }}
 
import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
	"goapp/dtos"
	// "gorm.io/gorm"
)

const {{ .StructName }}TableName = "{{ .TableName }}"

// {{ .StructName }}Model is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}}Model struct {
	// gorm.Model 如需伪删除等操作，内嵌gorm.Model可自行打开注释
{{- range .FieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 	// {{ .Comment }}
{{- end}}
}

func New{{.StructName}}Model() *{{.StructName}}Model {
	return new({{.StructName}}Model)
}

func (*{{ .StructName }}Model) TableName() string {
	return {{ .StructName }}TableName
}


// To DTO
func (m *{{ .StructName }}Model) ToDTO() *dtos.{{ .StructName }}DTO {
	return &dtos.{{ .StructName }}DTO{
		{{- range .FieldList }}
			{{ .Name }} : m.{{ .Name }},
		{{- end}}
	}
}

// From DTO
func (m *{{ .StructName }}Model) FromDTO(dto *dtos.{{ .StructName }}DTO) {
	{{- range .FieldList }}
		m.{{ .Name }} = dto.{{ .Name }}
	{{- end}}
}


`

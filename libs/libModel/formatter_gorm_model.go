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
	f.ModelFieldList = make([]Field, 0)
	f.DTOFieldList = make([]Field, 0)
	f.CreateDTOFieldList = make([]Field, 0)
	f.UpdateDTOFieldList = make([]Field, 0)

	for _, col := range cols {
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		if colType == CTypeTime {
			// f.ImportList["time"] = ImportItem{Alias: "", Package: "time"}
			f.ImportList["gorm"] = ImportItem{Alias: "", Package: "gorm.io/gorm"}
		}

		if !utils.InStringSlice(col.GetName(), []string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt"}) {
			f.ModelFieldList = append(f.ModelFieldList, Field{
				Name:      col.GetName(),
				Type:      colType,
				StructTag: fmt.Sprintf("`gorm:\"%s\" json:\"%s\"`", col.Name, col.Name),
				Comment:   col.GetComment(),
			})
		}

		if !utils.InStringSlice(col.GetName(), []string{"Id", "DeletedAt"}) {
			f.DTOFieldList = append(f.DTOFieldList, Field{
				Name: col.GetName(),
			})
		}

		if !utils.InStringSlice(col.GetName(), []string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt"}) {
			f.CreateDTOFieldList = append(f.CreateDTOFieldList, Field{
				Name: col.GetName(),
			})
			f.UpdateDTOFieldList = append(f.UpdateDTOFieldList, Field{
				Name: col.GetName(),
			})
		}

	}
	return f
}

func (f *FormatterGorm) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GormTemplate").Parse(GormStructCodeTemplate)).Execute(writer, *f)
}

// 完善Gorm Model模板，根据goapp_account的实践
const GormStructCodeTemplate = `package {{ .PackageName }}
 
import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
	"goapp/dtos"
)

const {{ .StructName }}TableName = "{{ .TableName }}"

// {{ .StructName }}Model is a mapping object for {{ .TableName }} table in mysql
type {{.StructName}}Model struct {
	gorm.Model
{{- range .ModelFieldList }}
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
		Id : m.ID,
		{{- range .DTOFieldList }}
			{{ .Name }} : m.{{ .Name }},
		{{- end}}
		DeletedAt:      m.DeletedAt.Time,
	}
}

// From DTO
func (m *{{ .StructName }}Model) FromDTO(dto *dtos.{{ .StructName }}DTO) {
	m.ID = dto.Id
	{{- range .ModelFieldList }}
		m.{{ .Name }} = dto.{{ .Name }}
	{{- end}}
}

// From CreateDTO
func (m *{{ .StructName }}Model) FromCreateDTO(dto *dtos.Create{{ .StructName }}DTO) {
	{{- range .CreateDTOFieldList }}
		m.{{ .Name }} = dto.{{ .Name }}
	{{- end}}
}

// From UpdateDTO
func (m *{{ .StructName }}Model) FromUpdateDTO(dto *dtos.Update{{ .StructName }}DTO) {
	m.ID = dto.Id
	{{- range .UpdateDTOFieldList }}
		m.{{ .Name }} = dto.{{ .Name }}
	{{- end}}
}

`

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
	f.DTOFieldList = make([]Field, len(cols)) // 所有字段
	f.CreateDTOFieldList = make([]Field, 0)   // 除去id、created_at、updated_at、deleted_at 的所有字段，用于创新条目时的数据校验
	f.UpdateDTOFieldList = make([]Field, 0)   // 除去 created_at、updated_at、deleted_at 的所有字段，用于更新条目时的数据校验

	for idx, col := range cols {
		var err error
		var dtoType string
		var colType string
		if colType, err = col.GetType(); err != nil {
			continue
		}
		if dtoType, err = col.GetDTOType(); err != nil {
			continue
		}

		if colType == CTypeTime {
			f.ImportList["time"] = ImportItem{Alias: "", Package: "time"}
			f.ImportList["validate"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XValidate"}
		}

		// DTO field
		f.DTOFieldList[idx] = Field{
			Name:      col.GetName(),
			Type:      colType,
			Comment:   col.GetComment(),
			StructTag: fmt.Sprintf("`json:\"%s\"`", col.Name),
		}

		// CreateDTO
		if !utils.InStringSlice(col.GetName(), []string{"Id", "CreatedAt", "UpdatedAt", "DeletedAt"}) {
			f.CreateDTOFieldList = append(f.CreateDTOFieldList, Field{
				Name:      col.GetName(),
				Type:      colType,
				Comment:   col.GetComment(),
				StructTag: fmt.Sprintf("`json:\"%s\" validate:\"%s\"`", col.Name, dtoType),
			})
		}

		// UpdateDTO
		if !utils.InStringSlice(col.GetName(), []string{"CreatedAt", "UpdatedAt", "DeletedAt"}) {
			f.UpdateDTOFieldList = append(f.UpdateDTOFieldList, Field{
				Name:      col.GetName(),
				Type:      colType,
				Comment:   col.GetComment(),
				StructTag: fmt.Sprintf("`json:\"%s\" validate:\"%s\"`", col.Name, dtoType),
			})
		}

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
{{- range .DTOFieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}

// Create{{ .StructName }}DTO
type Create{{.StructName}}DTO struct {
{{- range .CreateDTOFieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}


func (dto *Create{{.StructName}}DTO) Validate() error {
	return XValidate.V(dto)
}

// Update{{ .StructName }}DTO
type Update{{.StructName}}DTO struct {
{{- range .UpdateDTOFieldList }}
	{{ .Name }} {{ .Type }} {{ .StructTag }} 		// {{ .Comment }}
{{- end}}
}


func (dto *Update{{.StructName}}DTO) Validate() error {
	return XValidate.V(dto)
}


`

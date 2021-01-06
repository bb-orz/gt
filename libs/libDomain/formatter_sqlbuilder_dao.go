package libDomain

import (
	"gt/utils"
	"io"
	"text/template"
)

type FormatterDomainSqlBuilderDao struct {
	FormatterStruct
}

func NewFormatterDomainSqlBuilderDao() *FormatterDomainSqlBuilderDao {
	return new(FormatterDomainSqlBuilderDao)
}

func (f *FormatterDomainSqlBuilderDao) Format(name string) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(name)
	f.TableName = name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["xgorm"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XStore/XGorm"}
	f.ImportList["gorm"] = ImportItem{Alias: "", Package: "gorm.io/gorm"}

	return f
}

func (f *FormatterDomainSqlBuilderDao) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainSqlBuilderTemplate").Parse(DomainSqlBuilderDaoCodeTemplate)).Execute(writer, *f)
}

const DomainSqlBuilderDaoCodeTemplate = `
package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type {{ .StructName }}DAO struct{}

func New{{ .StructName }}DAO() *{{ .StructName }}DAO {
	dao := new({{ .StructName }}DAO)
	return dao
}

func (d *{{ .StructName }}DAO) isExist(where *{{ .StructName }}) (bool, error) {
	
	return false, nil
}

// 查找id是否存在
func (d *{{ .StructName }}DAO) IsIdExist(id uint) (bool, error) {

	return false,nil
}

// 通过Id查找
func (d *{{ .StructName }}DAO) GetById(id uint) (*dtos.{{ .StructName }}DTO, error) {
	
	return dto, nil
}


// 创建
func (d *{{ .StructName }}DAO) Create(dto *dtos.{{ .StructName }}DTO) (*dtos.{{ .StructName }}DTO, error) {
	
	return nil, nil
}



// 设置单个信息字段
func (d *{{ .StructName }}DAO) Set{{ .StructName }}(id uint, field string, value interface{}) error {
	
	return nil
}

// 设置多个信息字段
func (d *{{ .StructName }}DAO) Update{{ .StructName }}(id uint, dto dtos.{{ .StructName }}DTO) error {
	
	return nil
}

// 真删除
func (d *{{ .StructName }}DAO) DeleteById(id uint) error {
	
	return nil
}

// 伪删除
func (d *{{ .StructName }}DAO) SetDeletedAtById(id uint) error {
	
	return nil
}


`

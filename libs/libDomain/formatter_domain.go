package libDomain

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

type FormatterDomain struct {
	FormatterStruct
}

func NewFormatterDomain() *FormatterDomain {
	return new(FormatterDomain)
}

func (f *FormatterDomain) Format(name string) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(name)
	f.TableName = name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["common"] = ImportItem{Alias: "", Package: "goapp/common"}

	return f
}

func (f *FormatterDomain) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainTemplate").Parse(DomainCodeTemplate)).Execute(writer, *f)
}

const DomainCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

/*
{{ .StructName }} 领域层：实现{{ .StructName }}相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type {{ .StructName }}Domain struct {
	dao   *{{ .StructName }}DAO
}

func New{{ .StructName }}Domain() *{{ .StructName }}Domain {
	domain := new({{ .StructName }}Domain)
	domain.dao = New{{ .StructName }}DAO()
	return domain
}

func (domain *{{ .StructName }}Domain) DomainName() string {
	return "{{ .StructName }}Domain"
}

// 查找指定id数据是否已存在
func (domain *{{ .StructName }}Domain) Is{{ .StructName }}Exist(id uint) (bool, error) {
	var err error
	var isExist bool
	if isExist, err = domain.dao.IsIdExist(id); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsExist")
	} else if isExist {
		return true, nil
	}
	return false, nil
}

// 创建
func (domain *{{ .StructName }}Domain) Create{{ .StructName }}(dto dtos.{{ .StructName }}DTO) (int64, error) {
	var err error
	var insertId int64

	if insertId, err = domain.dao.Create(&dto); err != nil {
		return -1, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}
	return insertId, nil
}

func (domain *{{ .StructName }}Domain) Get{{ .StructName }}ById(id uint) (*dtos.{{ .StructName }}DTO, error) {
	var err error
	var result *dtos.{{ .StructName }}DTO
	if result, err = domain.dao.GetById(id); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetById")
	}
	return result, nil
}


// 设置单个字段信息
func (domain *{{ .StructName }}Domain) Set{{ .StructName }}Field(uid uint, field string, value interface{}) error {
	var err error
	if err = domain.dao.Set{{ .StructName }}(uid, field, value); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUserInfo")
	}
	return nil
}

// 设置多个信息
func (domain *{{ .StructName }}Domain) Update{{ .StructName }}Infos(id uint, dto dtos.{{ .StructName }}DTO) error {
	var err error
	if err = domain.dao.Update{{ .StructName }}(id, dto); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUserInfos")
	}

	return nil
}


// 真删除
func (domain *{{ .StructName }}Domain) Delete{{ .StructName }}(id uint) error {
	var err error
	if err = domain.dao.DeleteById(id); err != nil {
		return common.DomainInnerErrorOnSqlDelete(err, "DeleteById")
	}
	return nil
}

// 伪删除
func (domain *{{ .StructName }}Domain) ShamDelete{{ .StructName }}(id uint) error {
	var err error
	if err = domain.dao.SetDeletedAtById(id); err != nil {
		return common.DomainInnerErrorOnSqlShamDelete(err, "SetDeletedAtById")
	}
	return nil
}

`

package libModel

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

type FormatterSqlBuilderDao struct {
	FormatterStruct
}

func NewFormatterSqlBuilderDao() *FormatterSqlBuilderDao {
	return new(FormatterSqlBuilderDao)
}

func (f *FormatterSqlBuilderDao) Format(name, table string, cols []Column) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(table)
	f.TableName = table
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["time"] = ImportItem{Alias: "", Package: "time"}
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["common"] = ImportItem{Alias: "", Package: "goapp/common"}
	f.ImportList["sqlbuilder"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XStore/XSQLBuilder"}

	return f
}

func (f *FormatterSqlBuilderDao) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainSqlBuilderDaoTemplate").Parse(DomainSqlBuilderDaoCodeTemplate)).Execute(writer, *f)
}

const DomainSqlBuilderDaoCodeTemplate = `package {{ .PackageName }}

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

func New{{ .StructName }}DAO() *ResumesDAO {
	dao := new({{ .StructName }}DAO)
	return dao
}

// 查找id是否存在
func (d *{{ .StructName }}DAO) IsIdExist(id uint) (bool, error) {
	var err error
	var count int64
	var where map[string]interface{}
	where = map[string]interface{}{
		"id": id,
	}
	count, err = XSQLBuilder.XCommon().GetCount({{ .StructName }}TableName, where)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 通过Id查找
func (d *{{ .StructName }}DAO) GetById(id uint) (*dtos.{{ .StructName }}DTO, error) {
	var err error
	var where map[string]interface{}
	var result {{ .StructName }}Model

	where = map[string]interface{}{
		"id": id,
	}

	if err = XSQLBuilder.XCommon().GetOne({{ .StructName }}TableName, where, nil, &result); err != nil {
		return nil, err
	}
	return result.ToDTO(), nil
}

func (d *{{ .StructName }}DAO) Find(where map[string]interface{}, selectField []string) ([]*dtos.{{ .StructName }}DTO, error) {
	var err error
	var results []{{ .StructName }}Model
	var resultsDTO = make([]*dtos.{{ .StructName }}DTO, 0)
	if err = XSQLBuilder.XCommon().GetMulti({{ .StructName }}TableName, where, selectField, &results); err != nil {
		return nil, err
	}
	for _, v := range results {
		resultsDTO = append(resultsDTO, v.ToDTO())
	}
	return resultsDTO, nil
}

// 创建
func (d *{{ .StructName }}DAO) Create(dto *dtos.{{ .StructName }}DTO) (int64, error) {
	var err error
	var data []map[string]interface{}
	var insertId int64
	data = append(data, common.Struct2Map(dto))

	if insertId, err = XSQLBuilder.XCommon().Insert({{ .StructName }}TableName, data); err != nil {
		return -1, err
	}

	return insertId, nil
}

// 设置单个信息字段
func (d *{{ .StructName }}DAO) Set{{ .StructName }}(id uint, field string, value interface{}) error {
	var err error
	var where = make(map[string]interface{})
	var updater = make(map[string]interface{})
	where["id"] = id
	updater[field] = value
	if _, err = XSQLBuilder.XCommon().Update({{ .StructName }}TableName, where, updater); err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段,注意DTO Struct的零值,如不能避免零值设置错误请使用Update{{ .StructName }}WithMap方法
func (d *{{ .StructName }}DAO) UpdateResumes(id uint, dto dtos.ResumesDTO) error {
	var err error
	var where = make(map[string]interface{})
	var updater = make(map[string]interface{})
	where["id"] = id
	updater = common.Struct2Map(dto)
	if _, err = XSQLBuilder.XCommon().Update({{ .StructName }}TableName, where, updater); err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段,注意DTO Struct的零值
func (d *{{ .StructName }}DAO) Update{{ .StructName }}WithMap(id uint, updater map[string]interface{}) error {
	var err error
	var where = make(map[string]interface{})
	where["id"] = id
	if _, err = XSQLBuilder.XCommon().Update({{ .StructName }}TableName, where, updater); err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *{{ .StructName }}DAO) DeleteById(id uint) error {
	var err error
	var where = make(map[string]interface{})
	where["id"] = id
	if _, err = XSQLBuilder.XCommon().Delete({{ .StructName }}TableName, where); err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *{{ .StructName }}DAO) SetDeletedAtById(id uint) error {
	return d.SetResumes(id, "deleted_at", time.Now())
}

`

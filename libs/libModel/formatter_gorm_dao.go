package libModel

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

type FormatterGormDao struct {
	FormatterStruct
}

func NewFormatterGormDao() *FormatterGormDao {
	return new(FormatterGormDao)
}

func (f *FormatterGormDao) Format(name, table string, cols []Column) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(table)
	f.TableName = table
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}
	f.ImportList["xgorm"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XStore/XGorm"}
	f.ImportList["gorm"] = ImportItem{Alias: "", Package: "gorm.io/gorm"}

	return f
}

func (f *FormatterGormDao) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainGormDaoTemplate").Parse(DomainGormDaoCodeTemplate)).Execute(writer, *f)
}

// TODO 完善Gorm Dao模板，根据goapp_account的实践
const DomainGormDaoCodeTemplate = `package {{ .PackageName }}

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

func (d *{{ .StructName }}DAO) isExist(where *{{ .StructName }}Model) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&{{ .StructName }}Model{}).Where(where).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return false, nil
		} else {
			// 除无记录外的错误返回
			return false, err
		}
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 查找id是否存在
func (d *{{ .StructName }}DAO) IsIdExist(id uint) (bool, error) {
	return d.isExist(&{{ .StructName }}Model{Id: id})
}


// 通过Id查找
func (d *{{ .StructName }}DAO) GetById(id uint) (*dtos.{{ .StructName }}DTO, error) {
	var err error
	var {{ .TableName }}Result {{ .StructName }}Model
	err = XGorm.XDB().Where(id).First(&{{ .TableName }}Result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := {{ .TableName }}Result.ToDTO()
	return dto, nil
}


// 创建
func (d *{{ .StructName }}DAO) Create(dto *dtos.{{ .StructName }}DTO) (*dtos.{{ .StructName }}DTO, error) {
	var err error
	var {{ .TableName }}DTO *dtos.{{ .StructName }}DTO
	var {{ .TableName }}Model {{ .StructName }}Model

	{{ .TableName }}Model.FromDTO(dto)
	if err = XGorm.XDB().Create(&{{ .TableName }}Model).Error; err != nil {
		return nil, err
	}
	{{ .TableName }}DTO = {{ .TableName }}Model.ToDTO()
	return {{ .TableName }}DTO, nil
}



// 设置单个信息字段
func (d *{{ .StructName }}DAO) Set{{ .StructName }}(id uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&{{ .StructName }}Model{}).Where("id", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段
func (d *{{ .StructName }}DAO) Update{{ .StructName }}(id uint, dto dtos.{{ .StructName }}DTO) error {
	var err error

	if err = XGorm.XDB().Model(&{{ .StructName }}Model{}).Where("id", id).Updates(&dto).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *{{ .StructName }}DAO) DeleteById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&{{ .StructName }}Model{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *{{ .StructName }}DAO) SetDeletedAtById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&{{ .StructName }}Model{}).Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(id).Error; err != nil {
		return err
	}
	return nil
}

`
package libModel

import (
	"fmt"
	"gt/utils"
)

func NewFormatterGormStruct() *FormatterGormStruct {
	return new(FormatterGormStruct)
}

type FormatterGormStruct struct {
	FormatterStruct
}

func (f *FormatterGormStruct) Format(tableName string, cols []Column) IFormatter {
	f.PackageName = tableName
	f.StructName = utils.CamelString(tableName)
	f.TableName = tableName
	f.FieldList = make([]Field, len(cols))

	for idx, col := range cols {
		colType, err := col.GetType()
		if nil != err {
			continue
		}
		f.FieldList[idx] = Field{
			Name:      col.GetName(),
			Type:      colType,
			StructTag: fmt.Sprintf("`gorm:\"%s\" json:\"%s\"`", col.Name, col.Name),
		}
	}
	return f
}

func (f *FormatterGormStruct) WriteOut() error {
	utils.CommandLogger.Info(utils.CommandNameModel, fmt.Sprintf("Formatter:%+v \n", f))
	return nil
}

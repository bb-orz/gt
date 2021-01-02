package libModel

import (
	"fmt"
	"gt/utils"
)

func NewFormatterSqlBuilderStruct() *FormatterGormStruct {
	return new(FormatterGormStruct)
}

type FormatterSqlBuilderStruct struct {
	FormatterStruct
}

func (f *FormatterSqlBuilderStruct) Format(tableName string, cols []Column) IFormatter {
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
			StructTag: fmt.Sprintf("`ddb:\"%s\" json:\"%s\"`", col.Name, col.Name),
		}
	}
	return f
}

func (f *FormatterSqlBuilderStruct) WriteOut() error {

	return nil
}

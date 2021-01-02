package libModel

type IFormatter interface {
	Format(tableName string, cols []Column) IFormatter
	WriteOut() error
}

// 格式化信息结构体
type FormatterStruct struct {
	PackageName string
	StructName  string
	TableName   string
	FieldList   []Field
}

// 表列信息结构体
type Field struct {
	Name      string
	Type      string
	StructTag string
}

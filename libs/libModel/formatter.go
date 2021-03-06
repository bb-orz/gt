package libModel

import "io"

type IFormatter interface {
	Format(name, table string, cols []Column) IFormatter
	WriteOut(writer io.Writer) error
}

// 格式化信息结构体
type FormatterStruct struct {
	PackageName        string
	ImportList         map[string]ImportItem
	StructName         string
	TableName          string
	FieldList          []Field
	ModelFieldList     []Field
	DTOFieldList       []Field
	CreateDTOFieldList []Field
	UpdateDTOFieldList []Field
}

// 表列信息结构体
type Field struct {
	Name      string
	Type      string
	StructTag string
	Comment   string
}

type ImportItem struct {
	Alias   string
	Package string
}

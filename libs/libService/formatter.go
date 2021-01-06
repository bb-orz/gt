package libService

import "io"

type IFormatter interface {
	Format(name, version string) IFormatter
	WriteOut(writer io.Writer) error
}

// 格式化信息结构体
type FormatterStruct struct {
	PackageName string
	ImportList  map[string]ImportItem
	StructName  string
	Version     string
	FieldList   []Field
}

// struct列信息结构体
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

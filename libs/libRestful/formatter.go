package libRestful

import "io"

type IFormatter interface {
	Format(name string) IFormatter
	WriteOut(writer io.Writer) error
}

// 格式化信息结构体
type FormatterStruct struct {
	PackageName string
	ImportList  map[string]ImportItem
	StructName  string
	RouteGroup  string
}

type ImportItem struct {
	Alias   string
	Package string
}

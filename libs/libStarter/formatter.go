package libStarter

import "io"

type IFormatter interface {
	Format(cmdParams *CmdParams) IFormatter
	WriteOut(writer io.Writer) error
}

// 格式化信息结构体
type FormatterStruct struct {
	PackageName string
	ImportList  map[string]ImportItem
	Name        string
	StructName  string
	TypeName    string
}

type ImportItem struct {
	Alias   string
	Package string
}

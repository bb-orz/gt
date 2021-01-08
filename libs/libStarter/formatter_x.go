package libStarter

import (
	"io"
	"text/template"
)

func NewFormatterStarterX() *FormatterStarterX {
	return new(FormatterStarterX)
}

type FormatterStarterX struct {
	FormatterStruct
}

func (f *FormatterStarterX) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name
	return f
}

func (f *FormatterStarterX) WriteOut(writer io.Writer) error {
	return template.Must(template.New("StarterXTemplate").Parse(StarterXCodeTemplate)).Execute(writer, *f)
}

const StarterXCodeTemplate = `package {{ .PackageName }}

// 可供外部调用的运行实例

func X() {

	return
}

`

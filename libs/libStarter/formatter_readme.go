package libStarter

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterStarterReadme() *FormatterStarterReadme {
	return new(FormatterStarterReadme)
}

type FormatterStarterReadme struct {
	FormatterStruct
}

func (f *FormatterStarterReadme) Format(cmdParams *CmdParams) IFormatter {
	f.StructName = utils.CamelString(cmdParams.Name)
	f.Name = cmdParams.Name
	return f
}

func (f *FormatterStarterReadme) WriteOut(writer io.Writer) error {
	return template.Must(template.New("StarterReadmeTemplate").Parse(StarterReadmeCodeTemplate)).Execute(writer, *f)
}

const StarterReadmeCodeTemplate = `
# {{ .StructName }} Starter

> 基于  包

### Documentation

> 


### Starter Usage


### {{ .StructName }} Config Setting


### {{ .StructName }} Usage


`

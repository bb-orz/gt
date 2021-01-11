package libDomain

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

type FormatterDomainTesting struct {
	FormatterStruct
}

func NewFormatterDomainTesting() *FormatterDomainTesting {
	return new(FormatterDomainTesting)
}

func (f *FormatterDomainTesting) Format(name string) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(name)
	f.TableName = name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["convey"] = ImportItem{Alias: ".", Package: "github.com/smartystreets/goconvey/convey"}
	f.ImportList["testing"] = ImportItem{Alias: "", Package: "testing"}

	return f
}

func (f *FormatterDomainTesting) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainTestingTemplate").Parse(DomainTestingCodeTemplate)).Execute(writer, *f)
}

const DomainTestingCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)


func Test{{ .StructName }}Domain(t *testing.T) {
	Convey("{{ .StructName }} Domain Testing:", t, func() {
		
	})
}

`

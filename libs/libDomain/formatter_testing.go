package libDomain

import (
	"gt/utils"
	"io"
	"text/template"
)

type FormatterDomainTestingStruct struct {
	FormatterStruct
}

func NewFormatterDomainTestingStruct() *FormatterDomainTestingStruct {
	return new(FormatterDomainTestingStruct)
}

func (f *FormatterDomainTestingStruct) Format(name string) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(name)
	f.TableName = name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["convey"] = ImportItem{Alias: ".", Package: "github.com/smartystreets/goconvey/convey"}
	f.ImportList["testing"] = ImportItem{Alias: "", Package: "testing"}

	return f
}

func (f *FormatterDomainTestingStruct) WriteOut(writer io.Writer) error {
	err := template.Must(template.New("DomainTestingTemplate").Parse(DomainTestingCodeTemplate)).Execute(writer, *f)
	if err != nil {
		utils.CommandLogger.Error(utils.CommandNameDomain, err)
	}
	return nil
}

const DomainTestingCodeTemplate = `
package {{ .PackageName }}

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

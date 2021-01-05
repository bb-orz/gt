package libDomain

import (
	"gt/utils"
	"io"
	"text/template"
)

type FormatterDomainStruct struct {
	FormatterStruct
}

func NewFormatterDomainStruct() *FormatterDomainStruct {
	return new(FormatterDomainStruct)
}

func (f *FormatterDomainStruct) Format(name string) IFormatter {
	f.PackageName = name
	f.StructName = utils.CamelString(name)
	f.TableName = name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["dtos"] = ImportItem{Alias: "", Package: "goapp/dtos"}

	return f
}

func (f *FormatterDomainStruct) WriteOut(writer io.Writer) error {
	return template.Must(template.New("DomainTemplate").Parse(DomainCodeTemplate)).Execute(writer, *f)
}

const DomainCodeTemplate = `
package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

/*
{{ .StructName }} 领域层：实现{{ .StructName }}相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type {{ .StructName }}Domain struct {
	dao   *{{ .StructName }}DAO
}

func New{{ .StructName }}Domain() *{{ .StructName }}Domain {
	domain := new({{ .StructName }}Domain)
	domain.dao = New{{ .StructName }}DAO()
	return domain
}

func (domain *{{ .StructName }}Domain) DomainName() string {
	return "{{ .StructName }}Domain"
}



`

package libStarter

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterStarterTesting() *FormatterStarterTesting {
	return new(FormatterStarterTesting)
}

type FormatterStarterTesting struct {
	FormatterStruct
}

func (f *FormatterStarterTesting) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name
	f.StructName = utils.CamelString(cmdParams.Name)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["convey"] = ImportItem{Alias: ".", Package: "github.com/smartystreets/goconvey/convey"}
	f.ImportList["goinfras"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras"}
	f.ImportList["testing"] = ImportItem{Alias: "", Package: "testing"}
	return f
}

func (f *FormatterStarterTesting) WriteOut(writer io.Writer) error {
	return template.Must(template.New("StarterTestingTemplate").Parse(StarterTestingCodeTemplate)).Execute(writer, *f)
}

const StarterTestingCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)


func Test{{ .StructName }}(t *testing.T) {
	Convey("{{ .StructName }} Domain Testing:", t, func() {
		
	})
}

func TestStarter(t *testing.T) {
	Convey("Test {{ .StructName }} Starter", t, func() {
		logger := goinfras.NewCommandLineStarterLogger()
		sctx := goinfras.CreateDefaultStarterContext(nil, logger)
		s := NewStarter()
		s.Init(sctx)
		s.Setup(sctx)
		s.Check(sctx)
		s.Start(sctx)
	})
}
`

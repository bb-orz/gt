package libStarter

import (
	"gt/utils"
	"io"
	"text/template"
)

func NewFormatterStarter() *FormatterStarter {
	return new(FormatterStarter)
}

type FormatterStarter struct {
	FormatterStruct
}

func (f *FormatterStarter) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name
	f.StructName = utils.CamelString(cmdParams.Name)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["goinfras"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras"}

	return f
}

func (f *FormatterStarter) WriteOut(writer io.Writer) error {
	return template.Must(template.New("StarterTemplate").Parse(StarterCodeTemplate)).Execute(writer, *f)
}

const StarterCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

func NewStarter() *starter {
	s := new(starter)
	s.cfg = &Config{}
	return s
}

type starter struct {
	goinfras.BaseStarter
	cfg        *Config
}

func (s *starter) Name() string {
	return "{{ .StructName }}"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("{{ .StructName }}", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	s.cfg = &define
	Cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))

}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	
	

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "{{ .StructName }} Starter Setuped!")
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "{{ .StructName }} Starter Setup Successful!")
	return false
}

func (s *starter) Start(sctx *goinfras.StarterContext) {

}

func (s *starter) StartBlocking() bool {
	return false
}

func (s *starter) Stop() {

}

func (s *starter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *starter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}


`

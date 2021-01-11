package libRpc

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterMicroServer() *FormatterMicroServer {
	return new(FormatterMicroServer)
}

type FormatterMicroServer struct {
	FormatterStruct
}

func (f *FormatterMicroServer) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name + "Micro"
	f.StructName = utils.CamelString(cmdParams.Name)
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["goinfras"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras"}
	f.ImportList["services"] = ImportItem{Alias: "", Package: "goapp/services"}
	f.ImportList["micro"] = ImportItem{Alias: "", Package: "github.com/micro/go-micro"}
	f.ImportList["server"] = ImportItem{Alias: "", Package: "github.com/micro/go-micro/server"}

	return f
}

func (f *FormatterMicroServer) WriteOut(writer io.Writer) error {
	return template.Must(template.New("MicroServerTemplate").Parse(MicroServerCodeTemplate)).Execute(writer, *f)
}

const MicroServerCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
	"goapp/protobuf/generate/{{ .Name }}Pb"

)




var Cfg *Config
type Config struct {
	Name       string
	ServerHost string
	ServerPort int64
}


func NewStarter() *starter {
	s := new(starter)
	s.cfg = &Config{}
	return s
}

type starter struct {
	goinfras.BaseStarter
	cfg        *Config
	microServer server.Server
}

func (s *starter) Name() string {
	return "{{ .StructName }}{{ .TypeName }}"
}

func (s *starter) Init(sctx *goinfras.StarterContext) {
	var err error
	var define Config
	viper := sctx.Configs()
	if viper != nil {
		err = viper.UnmarshalKey("{{ .StructName }}{{ .TypeName }}", &define)
		sctx.PassWarning(s.Name(), goinfras.StepInit, err)
	}
	s.cfg = &define
	Cfg = &define
	sctx.Logger().Debug(s.Name(), goinfras.StepInit, fmt.Sprintf("Config: %+v ", define))

}

func (s *starter) Setup(sctx *goinfras.StarterContext) {
	var err error
	service := micro.NewService(
		micro.Name(s.cfg.Name),
		micro.Address(fmt.Sprintf("%s:%d", s.cfg.ServerHost, s.cfg.ServerPort)),
		)

	service.Init()
	s.microServer = service.Server()

	// 注册服务端处理方法
	if err = {{ .Name }}Pb.Register{{ .StructName }}{{ .TypeName }}ServiceHandler(s.microServer, &services.{{ .StructName }}{{ .TypeName }}Service{});err != nil {
		sctx.Logger().Error(s.Name(),goinfras.StepSetup,err)
	}

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "{{ .StructName }}{{ .TypeName }} Service Server Setuped!")
}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	if s.microServer != nil {
		sctx.Logger().OK(s.Name(), goinfras.StepCheck, "{{ .StructName }}{{ .TypeName }} Service Server Steup Successful!")
		return true
	}
	return false
}

func (s *starter) Start(sctx *goinfras.StarterContext) {
	// 服务运行
	sctx.Logger().OK(s.Name(), goinfras.StepStart, fmt.Sprintf("{{ .StructName }}{{ .TypeName }} Service Server Startup ... Listening: %s:%d", s.cfg.ServerHost, s.cfg.ServerPort))
	if err := s.microServer.Start(); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	}
}

func (s *starter) StartBlocking() bool {
	return true
}

func (s *starter) Stop() {
	_ = s.microServer.Stop()

}

func (s *starter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *starter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}


`

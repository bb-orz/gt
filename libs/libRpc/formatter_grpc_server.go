package libRpc

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterGrpcServer() *FormatterGrpcServer {
	return new(FormatterGrpcServer)
}

type FormatterGrpcServer struct {
	FormatterStruct
}

func (f *FormatterGrpcServer) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name + "Grpc"
	f.StructName = utils.CamelString(cmdParams.Name)
	f.TypeName = utils.CamelString(cmdParams.Type)
	f.Name = cmdParams.Name
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["net"] = ImportItem{Alias: "", Package: "net"}
	f.ImportList["goinfras"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras"}
	f.ImportList["services"] = ImportItem{Alias: "", Package: "goapp/services"}
	f.ImportList["grpc"] = ImportItem{Alias: "", Package: "google.golang.org/grpc"}

	return f
}

func (f *FormatterGrpcServer) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GrpcServerTemplate").Parse(GrpcServerCodeTemplate)).Execute(writer, *f)
}

const GrpcServerCodeTemplate = `package {{ .PackageName }}

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
	listener   net.Listener
	grpcServer *grpc.Server
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
	// 注册rpc服务
	var err error
	// 先启动一个监听端口
	if s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.ServerHost, s.cfg.ServerPort)); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepSetup, err)
	}

	// 创建grpc服务并注册
	s.grpcServer = grpc.NewServer()
	// 注册服务
	{{ .Name }}Pb.Register{{ .StructName }}{{ .TypeName }}ServiceServer(s.grpcServer, &services.{{ .StructName }}{{ .TypeName }}Service{})

	sctx.Logger().OK(s.Name(), goinfras.StepSetup, "Register {{ .StructName }}{{ .TypeName }} Service Server Successful!")

}

func (s *starter) Check(sctx *goinfras.StarterContext) bool {
	if s.grpcServer != nil && s.listener != nil {
		return true
	}
	return false
}

func (s *starter) Start(sctx *goinfras.StarterContext) {
	// 服务运行
	sctx.Logger().OK(s.Name(), goinfras.StepStart, fmt.Sprintf("{{ .StructName }}{{ .TypeName }} Service Server Startup ... Listening: %s:%d", s.cfg.ServerHost, s.cfg.ServerPort))
	if err := s.grpcServer.Serve(s.listener); err != nil {
		sctx.Logger().Error(s.Name(), goinfras.StepStart, err)
	}
}

func (s *starter) StartBlocking() bool {
	return true
}

func (s *starter) Stop() {
	s.grpcServer.Stop()

}

func (s *starter) PriorityGroup() goinfras.PriorityGroup {
	return goinfras.AppGroup
}

func (s *starter) Priority() int {
	return goinfras.DEFAULT_PRIORITY
}


`

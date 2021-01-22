package libRestful

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterEchoEngine() *FormatterEchoEngine {
	return new(FormatterEchoEngine)
}

type FormatterEchoEngine struct {
	FormatterStruct
}

func (f *FormatterEchoEngine) Format(name string) IFormatter {

	f.PackageName = "restful"
	f.StructName = utils.CamelString(name)
	f.RouteGroup = utils.SnakeString(name)
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["sync"] = ImportItem{Alias: "", Package: "sync"}
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["echo"] = ImportItem{Alias: "", Package: "github.com/labstack/echo/v4"}
	f.ImportList["xecho"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XEcho"}

	return f
}

func (f *FormatterEchoEngine) WriteOut(writer io.Writer) error {
	return template.Must(template.New("EchoRestfulTemplate").Parse(EchoRestfulCodeTemplate)).Execute(writer, *f)
}

const EchoRestfulCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

func init() {
	var once sync.Once
	once.Do(func() {
		// 初始化时自动注册该API到Echo Engine
		XEcho.RegisterApi(new({{ .StructName }}Api))
	})

}

type {{ .StructName }}Api struct {}

// SetRouter由Echo Engine 启动时调用
func (s *{{ .StructName }}Api) SetRoutes() {
	engine := XEcho.XEngine()
	{{ .RouteGroup }}Group := engine.Group("/{{ .RouteGroup }}")
	{{ .RouteGroup }}Group.GET("/foo", s.Foo)
	{{ .RouteGroup }}Group.GET("/bar", s.Bar)
}

func (s *{{ .StructName }}Api) Foo(ctx echo.Context) {
	// TODO Call Service Or Domain Method
	fmt.Println("{{ .StructName }}.Foo Restful API")
}

func (s *{{ .StructName }}Api) Bar(ctx echo.Context) {
	// TODO Call Service Or Domain Method
	fmt.Println("{{ .StructName }}.Bar Restful API")
}


`

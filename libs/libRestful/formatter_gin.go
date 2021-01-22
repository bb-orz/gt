package libRestful

import (
	"github.com/bb-orz/gt/utils"
	"io"
	"text/template"
)

func NewFormatterGinEngine() *FormatterGinEngine {
	return new(FormatterGinEngine)
}

type FormatterGinEngine struct {
	FormatterStruct
}

func (f *FormatterGinEngine) Format(name string) IFormatter {
	f.PackageName = "restful"
	f.StructName = utils.CamelString(name)
	f.RouteGroup = utils.SnakeString(name)
	f.ImportList = make(map[string]ImportItem)
	f.ImportList["sync"] = ImportItem{Alias: "", Package: "sync"}
	f.ImportList["fmt"] = ImportItem{Alias: "", Package: "fmt"}
	f.ImportList["gin"] = ImportItem{Alias: "", Package: "github.com/gin-gonic/gin"}
	f.ImportList["xgin"] = ImportItem{Alias: "", Package: "github.com/bb-orz/goinfras/XGin"}

	return f
}

func (f *FormatterGinEngine) WriteOut(writer io.Writer) error {
	return template.Must(template.New("GinRestfulTemplate").Parse(GinRestfulCodeTemplate)).Execute(writer, *f)
}

const GinRestfulCodeTemplate = `package {{ .PackageName }}

import (
	{{- range .ImportList }}
	{{ .Alias }} "{{ .Package }}"
	{{- end}}
)

func init() {
	var once sync.Once
	once.Do(func() {
		// 初始化时自动注册该API到Gin Engine
		XGin.RegisterApi(new({{ .StructName }}Api))
	})

}

type {{ .StructName }}Api struct {}

// SetRouter由Gin Engine 启动时调用
func (s *{{ .StructName }}Api) SetRoutes() {
	engine := XGin.XEngine()
	{{ .RouteGroup }}Group := engine.Group("/{{ .RouteGroup }}")
	{{ .RouteGroup }}Group.GET("/foo", s.Foo)
	{{ .RouteGroup }}Group.GET("/bar", s.Bar)
}

func (s *{{ .StructName }}Api) Foo(ctx *gin.Context) {
	// TODO Call Service Or Domain Method
	fmt.Println("{{ .StructName }}.Foo Restful API")
}

func (s *{{ .StructName }}Api) Bar(ctx *gin.Context) {
	// TODO Call Service Or Domain Method
	fmt.Println("{{ .StructName }}.Bar Restful API")
}


`

package libStarter

import (
	"io"
	"text/template"
)

func NewFormatterStarterConfig() *FormatterStarterConfig {
	return new(FormatterStarterConfig)
}

type FormatterStarterConfig struct {
	FormatterStruct
}

func (f *FormatterStarterConfig) Format(cmdParams *CmdParams) IFormatter {
	f.PackageName = cmdParams.Name
	return f
}

func (f *FormatterStarterConfig) WriteOut(writer io.Writer) error {
	return template.Must(template.New("StarterConfigTemplate").Parse(StarterConfigCodeTemplate)).Execute(writer, *f)
}

const StarterConfigCodeTemplate = `package {{ .PackageName }}

type Config struct {
	
}

func DefaultConfig() *Config {
	return &Config{
		
	}
}

`

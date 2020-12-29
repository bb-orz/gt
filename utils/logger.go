package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)


// 启动步骤常量
const (
	CommandInit  = "Init"

)

// 日志等级命名常量
const (
	DebugLevel   = "Debug"
	InfoLevel    = "Info"
	WarningLevel = "Warning"
	ErrorLevel   = "Error"
	FailLevel    = "Fail"
)

const (
	blue = "\033[97;44m"
	cyan = "\033[97;46m"

	whitef   = "\033[37m"
	yellowf  = "\033[33m"
	bluef    = "\033[34m"
	greenf   = "\033[32m"
	redf     = "\033[31m"
	magentaf = "\033[35m"

	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[1;97;41m"
	magenta = "\033[1;97;45m"

	reset = "\033[0m"
)

// 可定义多个输出
type StarterLoggerOutput struct {
	Formatter LogFormatter // 格式转化器
	Writers   io.Writer    // 输出
}

// 格式转化签名函数
type LogFormatter func(params LogFormatterParams) string

// 格式化输出参数
type LogFormatterParams struct {
	CommandName string  // 命令名称
	LogLevel  string    // 记录日志级别
	TimeStamp time.Time // 记录时间戳
	Message   string    // 记录信息
}


// 每个命令设置不同的颜色标示
func (p *LogFormatterParams) CommandColor() string {
	switch p.CommandName {
	case CommandInit:
		return white

	default:
		return cyan
	}
}

// 每种错误级别输出不同的颜色
func (p *LogFormatterParams) LogLevelColor() string {
	switch p.LogLevel {
	case DebugLevel:
		return greenf
	case InfoLevel:
		return whitef
	case WarningLevel:
		return yellowf
	case ErrorLevel:
		return redf
	case FailLevel:
		return magentaf
	default:
		return white
	}
}

// 颜色重置
func (p *LogFormatterParams) ResetColor() string {
	return reset
}

// 启动日志默认终端颜色输出格式
var defaultLogFormatter = func(param LogFormatterParams) string {
	var  commandColor, logLevelColor, resetColor string

	commandColor = param.CommandColor()
	logLevelColor = param.LogLevelColor()
	resetColor = param.ResetColor()

	return fmt.Sprintf("[%s %s %s] | %v | %s [%s] >>>>>> %s %s \n",
		commandColor, param.CommandName , resetColor,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		logLevelColor, param.LogLevel, param.Message, resetColor,
	)
}

// 启动日志文件输出格式
var commonWriterFormatter = func(param LogFormatterParams) string {
	return fmt.Sprintf("[%s] | %v | [%s] >>>>>>  %s",
		param.CommandName,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.LogLevel, param.Message,
	)
}

type IStarterLogger interface {
	ADebug(msg string)
	AInfo(msg string)
	AWarning(msg string)
	AError(err error)
	AFail(err error)
	SDebug(name, step, msg string)
	SInfo(name, step, msg string)
	SWarning(name, step, msg string)
	SError(name, step string, err error)
	SFail(name, step string, err error)
}

// 启动器日志记录器
type StarterLogger struct {
	Outputs []*StarterLoggerOutput
}

func (l *StarterLogger) Debug(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			LogLevel:  DebugLevel,
			TimeStamp: time.Now(),
			Message:   msg,
			// 是否增加caller
		}))
	}
}

func (l *StarterLogger) Info(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			LogLevel:  InfoLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) Warning(msg string) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			LogLevel:  WarningLevel,
			TimeStamp: time.Now(),
			Message:   msg,
		}))
	}
}
func (l *StarterLogger) Error(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			LogLevel:  ErrorLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}
func (l *StarterLogger) Fail(err error) {
	for _, o := range l.Outputs {
		_, _ = fmt.Fprint(o.Writers, o.Formatter(LogFormatterParams{
			LogLevel:  FailLevel,
			TimeStamp: time.Now(),
			Message:   err.Error(),
		}))
	}
}



// 标准颜色输出日志记录器
type CommandLineStarterLogger struct {
	StarterLogger
}

// 针对终端输出的默认启动日志记录器
func NewCommandLineStarterLogger() *CommandLineStarterLogger {
	logger := new(CommandLineStarterLogger)
	output := new(StarterLoggerOutput)
	output.Formatter = defaultLogFormatter
	output.Writers = os.Stdout
	logger.Outputs = make([]*StarterLoggerOutput, 0)
	logger.Outputs = append(logger.Outputs, output)
	return logger
}

package utils

import (
	"fmt"
	"os"
)

// 启动步骤常量
const (
	AppCmd             = "APP" // application
	CommandNameInit    = "Init"
	CommandNameModel   = "Model"
	CommandNameDomain  = "Domain"
	CommandNameService = "Service"
	CommandNameRestful = "Restful"
	CommandNameRpc     = "RPC"
	CommandNameStarter = "Starter"
)

// 日志等级命名常量
const (
	DebugLevel      = "Debug"
	InfoLevel       = "Info"
	SuccessfulLevel = "OK"
	WarningLevel    = "Warning"
	ErrorLevel      = "Error"
	FailLevel       = "Fail"
)

const (
	// 不带背景的字符颜色设置，用于各种日志级别信息输出
	bluef    = "\033[34m"
	whitef   = "\033[37m"
	greenf   = "\033[32m"
	yellowf  = "\033[33m"
	redf     = "\033[31m"
	magentaf = "\033[35m"

	// 带背景设的颜色设置，用于特指各种命令
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	cyan    = "\033[97;46m"
	blue    = "\033[97;44m"
	green   = "\033[97;42m"
	red     = "\033[1;97;41m"
	magenta = "\033[1;97;45m"

	reset = "\033[0m"
)

// 格式转化签名函数
type LogFormatter func(params LogFormatterParams) string

// 格式化输出参数
type LogFormatterParams struct {
	CommandName string // 命令名称
	LogLevel    string // 记录日志级别
	// TimeStamp   time.Time // 记录时间戳
	Message string // 记录信息
}

// 每个命令设置不同的颜色标示
func (p *LogFormatterParams) CommandColor() string {
	switch p.CommandName {
	case AppCmd:
		return cyan
	case CommandNameInit:
		return yellow
	case CommandNameModel:
		return blue
	case CommandNameDomain:
		return cyan
	case CommandNameService:
		return green
	case CommandNameRestful:
		return red
	case CommandNameRpc:
		return magenta
	case CommandNameStarter:
		return white
	default:
		return white
	}
}

// 每种错误级别输出不同的颜色
func (p *LogFormatterParams) LogLevelColor() string {
	switch p.LogLevel {
	case DebugLevel:
		return bluef
	case InfoLevel:
		return whitef
	case SuccessfulLevel:
		return greenf
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
	var commandColor, logLevelColor, resetColor string
	commandColor = param.CommandColor()
	logLevelColor = param.LogLevelColor()
	resetColor = param.ResetColor()

	return fmt.Sprintf("[%s %s %s] %s [%5s] >>>\t%s %s \n",
		commandColor, param.CommandName, resetColor,
		// param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		logLevelColor, param.LogLevel, param.Message, resetColor,
	)
}

var CommandLogger = &commandLogger{defaultLogFormatter}

// 启动器日志记录器
type commandLogger struct {
	formatter LogFormatter
}

func (l *commandLogger) Debug(cmd, msg string) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    DebugLevel,
		// TimeStamp:   time.Now(),
		Message: msg,
		// 是否增加caller
	}))
}

func (l *commandLogger) Info(cmd, msg string) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    InfoLevel,
		// TimeStamp:   time.Now(),
		Message: msg,
	}))
}

func (l *commandLogger) OK(cmd, msg string) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    SuccessfulLevel,
		// TimeStamp:   time.Now(),
		Message: msg,
	}))
}

func (l *commandLogger) Warning(cmd, msg string) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    WarningLevel,
		// TimeStamp:   time.Now(),
		Message: msg,
	}))

}

func (l *commandLogger) Error(cmd string, err error) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    ErrorLevel,
		// TimeStamp:   time.Now(),
		Message: err.Error(),
	}))

}

func (l *commandLogger) Fail(cmd string, err error) {
	_, _ = fmt.Fprint(os.Stdout, l.formatter(LogFormatterParams{
		CommandName: cmd,
		LogLevel:    FailLevel,
		// TimeStamp:   time.Now(),
		Message: err.Error(),
	}))

}

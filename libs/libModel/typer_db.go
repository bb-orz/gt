package libModel

import (
	"strings"
)

// 需要转换成Struct的类型转换函数切片
var TypeWrappers = []typeWrapper{i64TypeWrapper, byteTypeWrapper, intTypeWrapper, float64TypeWrapper, stringTypeWrapper, timeTypeWrapper}

/* 数据库类型与go数据类型之间的转换 */
const (
	CTypeInt64   = "int64"
	CTypeInt     = "int"
	CTypeUInt    = "uint"
	CTypeString  = "string"
	CTypeFloat64 = "float64"
	CTypeTime    = "time.Time"
	CTypeInt8    = "int8"
	CTypeUInt64  = "uint64"
	CTypeByte    = "byte"
	CUnsigned    = "unsigned"
)

type typer interface {
	Type() string
	Match() bool
}

type typeWrapper func(string) typer

func i64TypeWrapper(s string) typer {
	s = strings.ToLower(s)
	u := uint64Type(s)
	if u.Match() {
		return u
	}
	return int64Type(s)
}

func byteTypeWrapper(s string) typer {
	s = strings.ToLower(s)
	b := byteType(s)
	if b.Match() {
		return b
	}
	return int8Type(s)
}

func intTypeWrapper(s string) typer {
	s = strings.ToLower(s)
	u := uintType(s)
	if u.Match() {
		return u
	}
	return intType(s)
}

func stringTypeWrapper(s string) typer {
	return stringType(strings.ToLower(s))
}

func float64TypeWrapper(s string) typer {
	return float64Type(strings.ToLower(s))
}

func timeTypeWrapper(s string) typer {
	return timeType(s)
}

type int64Type string

func (i64 int64Type) Type() string {
	return CTypeInt64
}

func (i64 int64Type) Match() bool {
	if strings.Contains(string(i64), "bigint") {
		return true
	}
	return false
}

type uint64Type string

func (ui64 uint64Type) Type() string {
	return CTypeUInt64
}

func (ui64 uint64Type) Match() bool {
	s := string(ui64)
	return strings.Contains(s, CUnsigned) && int64Type(s).Match()
}

type byteType string

func (b byteType) Type() string {
	return CTypeByte
}

func (b byteType) Match() bool {
	s := string(b)
	return strings.Contains(s, CUnsigned) && int8Type(s).Match()
}

type int8Type string

func (b int8Type) Type() string {
	return CTypeInt8
}

func (b int8Type) Match() bool {
	return strings.Contains(string(b), "tinyint")
}

type uintType string

func (ui uintType) Type() string {
	return CTypeUInt
}

func (ui uintType) Match() bool {
	s := string(ui)
	return strings.Contains(s, CUnsigned) && intType(s).Match()
}

type intType string

func (i intType) Type() string {
	return CTypeInt
}

func (i intType) Match() bool {
	return strings.Contains(string(i), "int")
}

type stringType string

func (s stringType) Type() string {
	return CTypeString
}

func (s stringType) Match() bool {
	var supportType = []string{"char", "text"}
	ss := string(s)
	for _, t := range supportType {
		if strings.Contains(ss, t) {
			return true
		}
	}
	return false
}

type float64Type string

func (f64 float64Type) Type() string {
	return CTypeFloat64
}

func (f64 float64Type) Match() bool {
	return strings.Contains(string(f64), "float") || strings.Contains(string(f64), "decimal")
}

type timeType string

func (t timeType) Type() string {
	return CTypeTime
}

func (t timeType) Match() bool {
	return t == "timestamp" || t == "date" || t == "datetime"
}

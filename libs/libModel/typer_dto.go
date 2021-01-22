package libModel

import (
	"strings"
)

// 需要转换成Struct的类型转换函数切片
var TyperDTOWrappers = []typerDTOWrapper{i64TyperDTOWrapper, byteTyperDTOWrapper, intTyperDTOWrapper, float64TyperDTOWrapper, stringTyperDTOWrapper, timeTyperDTOWrapper}

/* 数据库类型与validate类型之间的转换 */
const (
	DTypeInt64   = "numeric"
	DTypeInt     = "numeric"
	DTypeUInt    = "numeric,gt=0"
	DTypeString  = "alphanum"
	DTypeFloat64 = "numeric"
	DTypeTime    = "numeric"
	DTypeInt8    = "numeric"
	DTypeUInt64  = "numeric,gt=0"
	DTypeByte    = "alphanum"
	DUnsigned    = "alphanum"
)

type typerDTO interface {
	TypeDTO() string
	MatchDTO() bool
}

type typerDTOWrapper func(string) typerDTO

func i64TyperDTOWrapper(s string) typerDTO {
	s = strings.ToLower(s)
	u := uint64TyperDTO(s)
	if u.MatchDTO() {
		return u
	}
	return int64TyperDTO(s)
}

func byteTyperDTOWrapper(s string) typerDTO {
	s = strings.ToLower(s)
	b := byteTyperDTO(s)
	if b.MatchDTO() {
		return b
	}
	return int8TyperDTO(s)
}

func intTyperDTOWrapper(s string) typerDTO {
	s = strings.ToLower(s)
	u := uintTyperDTO(s)
	if u.MatchDTO() {
		return u
	}
	return intTyperDTO(s)
}

func stringTyperDTOWrapper(s string) typerDTO {
	return stringTyperDTO(strings.ToLower(s))
}

func float64TyperDTOWrapper(s string) typerDTO {
	return float64TyperDTO(strings.ToLower(s))
}

func timeTyperDTOWrapper(s string) typerDTO {
	return timeTyperDTO(s)
}

type int64TyperDTO string

func (i64 int64TyperDTO) TypeDTO() string {
	return DTypeInt64
}

func (i64 int64TyperDTO) MatchDTO() bool {
	if strings.Contains(string(i64), "bigint") {
		return true
	}
	return false
}

type uint64TyperDTO string

func (ui64 uint64TyperDTO) TypeDTO() string {
	return DTypeUInt64
}

func (ui64 uint64TyperDTO) MatchDTO() bool {
	s := string(ui64)
	return strings.Contains(s, DUnsigned) && int64TyperDTO(s).MatchDTO()
}

type byteTyperDTO string

func (b byteTyperDTO) TypeDTO() string {
	return DTypeByte
}

func (b byteTyperDTO) MatchDTO() bool {
	s := string(b)
	return strings.Contains(s, DUnsigned) && int8TyperDTO(s).MatchDTO()
}

type int8TyperDTO string

func (b int8TyperDTO) TypeDTO() string {
	return DTypeInt8
}

func (b int8TyperDTO) MatchDTO() bool {
	return strings.Contains(string(b), "tinyint")
}

type uintTyperDTO string

func (ui uintTyperDTO) TypeDTO() string {
	return DTypeUInt
}

func (ui uintTyperDTO) MatchDTO() bool {
	s := string(ui)
	return strings.Contains(s, DUnsigned) && intTyperDTO(s).MatchDTO()
}

type intTyperDTO string

func (i intTyperDTO) TypeDTO() string {
	return DTypeInt
}

func (i intTyperDTO) MatchDTO() bool {
	return strings.Contains(string(i), "int")
}

type stringTyperDTO string

func (s stringTyperDTO) TypeDTO() string {
	return DTypeString
}

func (s stringTyperDTO) MatchDTO() bool {
	var supportType = []string{"char", "text"}
	ss := string(s)
	for _, t := range supportType {
		if strings.Contains(ss, t) {
			return true
		}
	}
	return false
}

type float64TyperDTO string

func (f64 float64TyperDTO) TypeDTO() string {
	return DTypeFloat64
}

func (f64 float64TyperDTO) MatchDTO() bool {
	return strings.Contains(string(f64), "float") || strings.Contains(string(f64), "decimal")
}

type timeTyperDTO string

func (t timeTyperDTO) TypeDTO() string {
	return DTypeTime
}

func (t timeTyperDTO) MatchDTO() bool {
	return t == "timestamp" || t == "date" || t == "datetime"
}

package errors

import ()

const (
	ErrCodeConf = 35001
)

var (
	ErrConf = newConf(ErrCodeConf, "配置错误") //一些极少出现，且不适合返回的错误
)

func newConf(code int32, detail string) error {
	return New("", "", code)
}

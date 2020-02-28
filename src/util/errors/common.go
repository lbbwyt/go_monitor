package errors

func EntityIsNil() error {
	return &Error{
		Code:   EntityIsNilCode,
		Status: StatusText(EntityIsNilCode),
	}
}

const (
	EntityIsNilCode = 100001 // RFC 7231, 6.2.1

)

var statusText = map[int]string{
	EntityIsNilCode: "实体为空",
}

func StatusText(code int) string {
	return statusText[code]
}

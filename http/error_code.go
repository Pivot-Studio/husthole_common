package http

import "fmt"

type ErrorCode struct {
	CodeInner int
	Msg       string
}

func (e *ErrorCode) Error() string {
	return "E" + fmt.Sprint(e.CodeInner) + ":" + e.Msg
}

func (e *ErrorCode) Code() int {
	return e.CodeInner
}

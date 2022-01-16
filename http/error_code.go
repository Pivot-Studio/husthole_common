package http

import "fmt"

/**
 * @Author: nick
 * @Date: 2022/01/16 4:45 PM
 * @Version: 1.0.0
 */

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

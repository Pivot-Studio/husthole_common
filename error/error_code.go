package error

import "fmt"

type Error struct {
	code int
	msg  string
}

func (e *Error) Error() string {
	return "E" + fmt.Sprint(e.code) + ":" + e.msg
}

func (e *Error) Code() int {
	return e.code
}

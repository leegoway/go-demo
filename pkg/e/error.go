package e

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type ErrorCoder interface {
	Error() string // 打log等错误信息
	Code() uint32  // 错误码，第一次就正确生成
	Msg() string   // 给用户的提示信息
	Where() string // 第一次生成这个错的地方, 第一次: 当newCoder和wrap一个非errorCoder的时候
}

type ErrorCode struct {
	code  uint32
	msg   string
	where string
}

func (e *ErrorCode) Code() uint32 {
	return e.code
}

// 不带code的错误消息
func (e *ErrorCode) Msg() string {
	return e.msg
}

// 错误，附带code
func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s", e.code, e.msg)
}

func (e *ErrorCode) Where() string {
	return e.where
}

func New(msg string) *ErrorCode {
	where := caller(1, false)
	return &ErrorCode{code: UNKNOWN, msg: msg, where: where}
}

func NewCoder(code uint32, msg string, extMsg ...string) *ErrorCode {
	if msg == "" {
		msg = MsgFlags[code]
	}
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}
	where := caller(1, false)
	return &ErrorCode{code: code, msg: msg, where: where}
}

func NewCodere(code uint32, err error, extMsg ...string) *ErrorCode {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}
	where := caller(1, false)
	return &ErrorCode{code: code, msg: msg, where: where}
}

func Wrap(err error, extMsg ...string) *ErrorCode {
	var msg string
	var code uint32
	var where string
	switch v := err.(type) {
	case ErrorCoder:
		msg = v.Msg()
		code = v.Code()
		where = v.Where()
	default:
		msg = v.Error()
		code = UNKNOWN
		where = caller(1, false)
	}
	if len(extMsg) != 0 {
		msg = strings.Join(extMsg, " : ") + " : " + msg
	}
	return &ErrorCode{code: code, msg: msg, where: where}
}

//找到代码报错的位置
func caller(calldepth int, short bool) string {
	_, file, line, ok := runtime.Caller(calldepth + 1)
	if !ok {
		file = "???"
		line = 0
	} else if short {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}
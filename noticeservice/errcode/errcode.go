package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code int                    //`json:"code"`
	msg  string                 // `json:"msg"`
	data map[string]interface{} //`json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}
func GetError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

//给Error类定义方法
func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息:%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() map[string]interface{} {
	return e.data
}

func (e *Error) WithDetails(datas map[string]interface{}) *Error {
	e.data = datas
	// for _, d := range details {
	// 	e.data = append(e.details, d)
	// }
	return e
}

func (e *Error) StatusCode() int {
	switch e.code {
	case ServerError.Code():
		fallthrough
	case NgIOTServerError.Code():
		return http.StatusInternalServerError
	case NetError.Code():
		fallthrough
	case RegisterNetError.Code():
		fallthrough
	case FilerNetError.Code():
		fallthrough
	case RPCNetError.Code():
		fallthrough
	// case NbIOTNetError.Code():
	// 	fallthrough
	case NgIOTNetError.Code():
		return http.StatusOK
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedPlatNotExist.Code():
		fallthrough
	case UnauthorizedUserNotExist.Code():
		fallthrough
	case UnauthorizedUserDisabled.Code():
		fallthrough
	case UnauthorizedRoleNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTimeout.Code():
		fallthrough
	case HaveNewToken.Code():
		fallthrough
	case UnauthorizedGenerate.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	default:
		return http.StatusOK
	}
}

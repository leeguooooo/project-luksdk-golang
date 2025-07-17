package luksdkerrors

import (
	"errors"
	"fmt"
)

var (
	LukSDKErrorInternalError = NewLuksdkError(100000, "LukSDK: 服务器内部异常")
	LukSDKErrorParamError    = NewLuksdkError(100001, "LukSDK: 参数错误")
	LukSDKErrorRetryError    = NewLuksdkError(100002, "LukSDK: 请稍后再试")
	LukSDKErrorContentError  = NewLuksdkError(100003, "LukSDK: 资源不存在")
	LukSDKErrorChannelError  = NewLuksdkError(100004, "LukSDK: 渠道已禁用")
	LukSDKErrorSignError     = NewLuksdkError(100005, "LukSDK: 签名校验失败")
	LukSDKErrorLoginError    = NewLuksdkError(100006, "LukSDK: 未登录或 Token 已过期")
	LukSDKErrorCallbackError = NewLuksdkError(100007, "LukSDK: 渠道方回调地址响应解析失败")
)

func NewLuksdkError(code int, msg string) *LuksdkError {
	return &LuksdkError{
		code: code,
		msg:  msg,
	}
}

type LuksdkError struct {
	code int
	msg  string
}

func (e *LuksdkError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.msg)
}

func (e *LuksdkError) Code() int {
	return e.code
}

func (e *LuksdkError) Message() string {
	return e.msg
}

func (e *LuksdkError) MessageP() *string {
	return &e.msg
}

func (e *LuksdkError) With(m ...any) *LuksdkError {
	var msg []string
	switch len(m) {
	case 0:
		return e
	case 1:
		msg = append(msg, fmt.Sprintf("%v", m[0]))
	default:
		for _, v := range m {
			msg = append(msg, fmt.Sprintf("%v", v))
		}
	}
	return NewLuksdkError(e.code, fmt.Errorf("%w: %v", e, msg).Error())
}

func ParseErrorCode(err error) int {
	var e *LuksdkError
	if errors.As(err, &e) {
		return e.code
	}
	return LukSDKErrorInternalError.Code()
}

func ConvertError(err error) *LuksdkError {
	if err == nil {
		return LukSDKErrorInternalError.With("nil")
	}

	var v *LuksdkError
	if !errors.As(err, &v) {
		return LukSDKErrorInternalError.With(err)
	} else {
		return v
	}
}

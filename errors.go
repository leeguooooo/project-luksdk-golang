package luksdk

import "errors"

var (
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrInvalidParams = regError(1000, "invalid params") // 参数有误
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrInvalidChannel = regError(1001, "invalid channel") // 渠道有误
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrInvalidChannelOrder = regError(1002, "invalid channel request") // 渠道请求异常
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrInvalidSignature = regError(1003, "invalid signature") // 签名有误
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrInvalidGame = regError(1004, "invalid game") // 游戏有误
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrChannelDataException = regError(1005, "channel data exception") // 渠道返回数据异常
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrRepeatOrder = regError(1006, "repeat order") // 重复下订单
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrOrderFailed = regError(1007, "order failed") // 下单失败
	// Deprecated: 请使用 github.com/CFGameTech/project-luksdk-golang/luksdk 替代
	ErrOrderNotExist = regError(1008, "order not exist") // 订单不存在
)

var errorMap = make(map[error]int)

func regError(code int, msg string) error {
	err := errors.New(msg)
	errorMap[err] = code
	return err
}

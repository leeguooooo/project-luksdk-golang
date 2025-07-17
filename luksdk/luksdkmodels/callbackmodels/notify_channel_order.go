package callbackmodels

import "github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"

type NotifyChannelOrderRequest struct {
	// 订单列表
	Data []NotifyChannelOrderRequestDatum `json:"data"`
	// 随机字符串
	Nonce *string `json:"nonce"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type NotifyChannelOrderRequestDatum struct {
	// App ID
	AppId int64 `json:"c_id"`
	// 用户 ID
	UserId string `json:"c_uid"`
	// 奖励货币量，一些场景用户获得奖励需要发放的货币量
	CoinsAward *int64 `json:"coins_award"`
	// 预扣货币量，订单创建时的预扣货币量
	CoinsCost *int64 `json:"coins_cost"`
	// 游戏 ID
	GameId int64 `json:"g_id"`
	// 游戏订单 ID
	GameOrderID string `json:"game_order_id"`
	// 订单状态
	Status int64 `json:"status"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Token，长期有效访问令牌
	Token string `json:"token"`
}

type NotifyChannelOrderResponse struct {
	// 请求状态码，当值为 0 时表示请求成功
	Code int `json:"code"`
	// 订单响应列表
	Data []NotifyChannelOrderResponseDatum `json:"data"`
	// 请求状态说明
	Msg *string `json:"msg"`
}

type NotifyChannelOrderResponseDatum struct {
	// 用户 ID
	UserId string `json:"c_uid"`
	// 剩余货币量
	Coins *int64 `json:"coins"`
	// 渠道方订单 ID
	OrderID string `json:"order_id"`
}

func (req *NotifyChannelOrderRequest) Response() *NotifyChannelOrderResponse {
	return &NotifyChannelOrderResponse{}
}

func (resp *NotifyChannelOrderResponse) WithError(err error) *NotifyChannelOrderResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *NotifyChannelOrderResponse) WithCode(code int) *NotifyChannelOrderResponse {
	resp.Code = code
	return resp
}

func (resp *NotifyChannelOrderResponse) WithMsg(msg string) *NotifyChannelOrderResponse {
	resp.Msg = &msg
	return resp
}

func (resp *NotifyChannelOrderResponse) WithData(data []NotifyChannelOrderResponseDatum) *NotifyChannelOrderResponse {
	resp.Data = data
	return resp
}

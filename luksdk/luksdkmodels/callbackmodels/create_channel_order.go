package callbackmodels

import "github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"

type CreateChannelOrderRequest struct {
	Data      []CreateChannelOrderRequestDatum `json:"data"`
	Nonce     *string                          `json:"nonce"`
	Sign      string                           `json:"sign"`
	Timestamp *int64                           `json:"timestamp,omitempty"`
}

type CreateChannelOrderRequestDatum struct {
	AppId       int64   `json:"c_id"`
	GameId      int64   `json:"g_id"`
	RoomID      string  `json:"room_id"`
	UserId      string  `json:"c_uid"`
	CoinsCost   *int64  `json:"coins_cost"`
	EXT         *string `json:"ext,omitempty"`
	GameOrderID string  `json:"game_order_id"`
	Timestamp   *int64  `json:"timestamp,omitempty"`
	Token       string  `json:"token"`
}

type CreateChannelOrderResponse struct {
	Code int                               `json:"code"`
	Msg  *string                           `json:"msg"`
	Data []CreateChannelOrderResponseDatum `json:"data"`
}

type CreateChannelOrderResponseDatum struct {
	AppId   string `json:"c_uid"`
	Coins   *int64 `json:"coins"`
	OrderID string `json:"order_id"`
	Status  int64  `json:"status"`
}

func (req *CreateChannelOrderRequest) Response() *CreateChannelOrderResponse {
	return &CreateChannelOrderResponse{}
}

func (resp *CreateChannelOrderResponse) WithError(err error) *CreateChannelOrderResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *CreateChannelOrderResponse) WithCode(code int) *CreateChannelOrderResponse {
	resp.Code = code
	return resp
}

func (resp *CreateChannelOrderResponse) WithMsg(msg string) *CreateChannelOrderResponse {
	resp.Msg = &msg
	return resp
}

func (resp *CreateChannelOrderResponse) WithData(data []CreateChannelOrderResponseDatum) *CreateChannelOrderResponse {
	resp.Data = data
	return resp
}

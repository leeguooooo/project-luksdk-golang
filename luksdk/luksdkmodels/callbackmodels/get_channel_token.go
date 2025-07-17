package callbackmodels

import "github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"

type GetChannelTokenRequest struct {
	AppId     int64   `json:"c_id"`
	UserId    string  `json:"c_uid"`
	Code      *string `json:"code"`
	Timestamp *int64  `json:"timestamp"`
	Sign      string  `json:"sign"`
}

type GetChannelTokenResponse struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
	Data struct {
		Token    string `json:"token"`
		LeftTime int64  `json:"left_time"`
	} `json:"data,omitempty"`
}

func (req *GetChannelTokenRequest) Response() *GetChannelTokenResponse {
	return &GetChannelTokenResponse{}
}

func (resp *GetChannelTokenResponse) WithError(err error) *GetChannelTokenResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *GetChannelTokenResponse) WithCode(code int) *GetChannelTokenResponse {
	resp.Code = code
	return resp
}

func (resp *GetChannelTokenResponse) WithMsg(msg string) *GetChannelTokenResponse {
	resp.Msg = &msg
	return resp
}

func (resp *GetChannelTokenResponse) WithData(token string, leftTime int64) *GetChannelTokenResponse {
	resp.Data.Token = token
	resp.Data.LeftTime = leftTime
	return resp
}

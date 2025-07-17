package callbackmodels

import "github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"

type RefreshChannelTokenRequest struct {
	AppId     int64  `json:"c_id"`
	UserId    string `json:"c_uid"`
	Timestamp *int64 `json:"timestamp"`
	Sign      string `json:"sign"`
	Token     string `json:"token"`
	LeftTime  int64  `json:"left_time"`
}

type RefreshChannelTokenResponse struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
	Data struct {
		Token    string `json:"token"`
		LeftTime int64  `json:"left_time"`
	} `json:"data,omitempty"`
}

func (req *RefreshChannelTokenRequest) Response() *RefreshChannelTokenResponse {
	return &RefreshChannelTokenResponse{}
}

func (resp *RefreshChannelTokenResponse) WithError(err error) *RefreshChannelTokenResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *RefreshChannelTokenResponse) WithData(token string, leftTime int64) *RefreshChannelTokenResponse {
	resp.Data.Token = token
	resp.Data.LeftTime = leftTime
	return resp
}

func (resp *RefreshChannelTokenResponse) WithCode(code int) *RefreshChannelTokenResponse {
	resp.Code = code
	return resp
}

func (resp *RefreshChannelTokenResponse) WithMsg(msg string) *RefreshChannelTokenResponse {
	resp.Msg = &msg
	return resp
}

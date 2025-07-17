package callbackmodels

import (
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels"
)

type GetChannelUserInfoRequest struct {
	AppId     int64  `json:"c_id"`
	UserId    string `json:"c_uid"`
	GameId    int64  `json:"g_id"`
	Sign      string `json:"sign"`
	Timestamp *int64 `json:"timestamp,omitempty"`
	Token     string `json:"token"`
}

type GetChannelUserInfoResponse struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
	Data struct {
		UserId   string                 `json:"c_uid"`
		Avatar   string                 `json:"avatar"`
		Name     string                 `json:"name"`
		Coins    int64                  `json:"coins"`
		Identity *luksdkmodels.Identity `json:"identity,omitempty"`
	} `json:"data,omitempty"`
}

func (req *GetChannelUserInfoRequest) Response() *GetChannelUserInfoResponse {
	resp := &GetChannelUserInfoResponse{}
	resp.Data.UserId = req.UserId
	return resp
}

func (resp *GetChannelUserInfoResponse) WithError(err error) *GetChannelUserInfoResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *GetChannelUserInfoResponse) WithCode(code int) *GetChannelUserInfoResponse {
	resp.Code = code
	return resp
}

func (resp *GetChannelUserInfoResponse) WithMsg(msg string) *GetChannelUserInfoResponse {
	resp.Msg = &msg
	return resp
}

func (resp *GetChannelUserInfoResponse) WithUserId(userId string) *GetChannelUserInfoResponse {
	resp.Data.UserId = userId
	return resp
}

func (resp *GetChannelUserInfoResponse) WithAvatar(avatar string) *GetChannelUserInfoResponse {
	resp.Data.Avatar = avatar
	return resp
}

func (resp *GetChannelUserInfoResponse) WithName(name string) *GetChannelUserInfoResponse {
	resp.Data.Name = name
	return resp
}

func (resp *GetChannelUserInfoResponse) WithCoins(coins int64) *GetChannelUserInfoResponse {
	resp.Data.Coins = coins
	return resp
}

func (resp *GetChannelUserInfoResponse) WithIdentity(identity luksdkmodels.Identity) *GetChannelUserInfoResponse {
	resp.Data.Identity = &identity
	return resp
}

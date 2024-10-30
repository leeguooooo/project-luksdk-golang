package luksdk

import "testing"

func TestSDK_VerifySignature(t *testing.T) {
	sdk := New("fa7ad21fdbe10218024f88538a86")

	request := &GetChannelTokenRequest{
		CId:       1010997,
		CUid:      "123",
		Code:      "",
		Timestamp: 1730255315,
		Sign:      "DF32DD6C2C939FE896738F80EFF4E80B",
	}

	if err := sdk.VerifySignature(request.Sign, request); err != nil {
		panic(err)
	}
}

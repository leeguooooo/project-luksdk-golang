package luksdk

import "testing"

func TestSDK_VerifySignature(t *testing.T) {
	sdk := New("fa7ad21fdbe10218024f88538a86")

	request := &CreateChannelOrderRequest{
		Sign: "DF32DD6C2C939FE896738F80EFF4E80B",
		Data: []*CreateChannelOrderRequestEntry{
			{
				CId:         1,
				CRoomId:     "room_id",
				CUid:        "uid",
				CoinsCost:   100,
				GId:         1,
				GameOrderId: "order_id",
				ScoreCost:   100,
			},
			{
				CId:         2,
				CRoomId:     "room_id",
				CUid:        "uid",
				CoinsCost:   100,
				GId:         2,
				GameOrderId: "order_id",
				ScoreCost:   100,
			},
		},
	}

	if err := sdk.VerifySignature(request.Sign, request); err != nil {
		panic(err)
	}
}

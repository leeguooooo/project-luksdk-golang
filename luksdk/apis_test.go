package luksdk_test

import (
	"fmt"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels/apimodels"
	"testing"
	"time"
)

func TestApis_GetGameServiceList(t *testing.T) {
	resp, err := LukSDK.Apis.GetGameServiceList(apimodels.GetGameServiceListRequest{
		Sign:      "",
		Timestamp: &NowUnix,
	})
	checkFatal(t, err)
	assertionInt(t, resp.Code, 0)
	t.Log(resp)
}

func TestApis_QueryNotifyEvent(t *testing.T) {
	resp, err := LukSDK.Apis.QueryNotifyEvent(apimodels.QueryNotifyEventRequest{
		EndAt:     nil,
		GameID:    GameId,
		PageNo:    1,
		PageSize:  1,
		RoomID:    nil,
		Sign:      "",
		StartAt:   nil,
		Timestamp: nil,
		Type:      nil,
	})
	checkFatal(t, err)
	assertionInt(t, resp.Code, 0)
	t.Log(resp)
}

func TestApis_QueryOrder(t *testing.T) {
	resp, err := LukSDK.Apis.QueryOrder(apimodels.QueryOrderRequest{
		AppOrderNo:  nil,
		GameID:      GameId,
		GameOrderNo: nil,
		Nonce:       nil,
		Sign:        "",
		Timestamp:   nil,
	})
	checkFatal(t, err)
	assertionInt(t, resp.Code, 100001)
	t.Log(resp)
}

func TestApis_PublishControlEvent(t *testing.T) {
	t.Run("JoinGame", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			JoinGame("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("LeaveGame", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			LeaveGame("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("ChangeReadyStatus", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			ChangeReadyStatus("user", true)
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("KickPlayer", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			KickPlayer("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("StartGame", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			StartGame()
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("ForceCloseGame", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			ForceCloseGame()
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("ChangeRoomSetting", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			ChangeRoomSetting("")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("ChangeUserIdentity", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			ChangeUserIdentity("user", luksdkmodels.IdentityAdmin)
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("ChangeUserIdentity", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			SyncRoomSeat()
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("RefreshUserInfo", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			RefreshUserInfo("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("QuickStartGame", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			WithRoomId("0").
			QuickStartGame([]string{"user"})
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("IssueProps", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			IssueProps("user", fmt.Sprintf("luksdk-go-%d", time.Now().UnixMilli()), nil)
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("FetchBagStatus", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			FetchBagStatus("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("QueryIssuePropStatus", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			QueryIssuePropStatus("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("QueryIssuePropStatus", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			QueryIssuePropStatus("user")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("EquipProp", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			EquipProp("user", "prop")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
	t.Run("EquipProp", func(t *testing.T) {
		builder := apimodels.NewControlEventBuilder(AppId, GameId)
		_, err := builder.
			UnequipProp("user", "prop")
		checkFatal(t, err)
		resp, err := LukSDK.Apis.PublishControlEvent(builder.Build())
		checkFatal(t, err)
		assertionInt(t, resp.Code, 100003)
		t.Log(resp)
	})
}

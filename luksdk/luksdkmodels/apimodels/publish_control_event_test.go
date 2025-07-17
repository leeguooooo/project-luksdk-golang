package apimodels

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels"
)

func TestControlEventBuilder_JoinGame(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	builder.WithRoomId("room_123").WithSign("test_signature")

	request, err := builder.JoinGame("user_001",
		WithAutoStartNum(4),
		WithReady(true),
		WithSeat(1),
	)

	if err != nil {
		t.Fatalf("JoinGame failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeJoinGame {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeJoinGame, request.Type)
	}

	if request.AppId != 12345 {
		t.Errorf("Expected AppId 12345, got %d", request.AppId)
	}

	if request.GameId != 67890 {
		t.Errorf("Expected GameId 67890, got %d", request.GameId)
	}

	if *request.RoomId != "room_123" {
		t.Errorf("Expected RoomId 'room_123', got %s", *request.RoomId)
	}

	// 验证 Data 字段是否为有效的 JSON
	var joinGameData PublishControlEventRequestJoinGame
	err = json.Unmarshal([]byte(request.Data), &joinGameData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if joinGameData.UserId != "user_001" {
		t.Errorf("Expected UserId 'user_001', got %s", joinGameData.UserId)
	}

	if *joinGameData.AutoStartNum != 4 {
		t.Errorf("Expected AutoStartNum 4, got %d", *joinGameData.AutoStartNum)
	}

	if !*joinGameData.Ready {
		t.Errorf("Expected Ready true, got %v", *joinGameData.Ready)
	}

	if *joinGameData.Seat != 1 {
		t.Errorf("Expected Seat 1, got %d", *joinGameData.Seat)
	}
}

func TestControlEventBuilder_LeaveGame(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	request, err := builder.LeaveGame("user_001")

	if err != nil {
		t.Fatalf("LeaveGame failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeLeaveGame {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeLeaveGame, request.Type)
	}

	var leaveGameData PublishControlEventRequestLeaveGame
	err = json.Unmarshal([]byte(request.Data), &leaveGameData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if leaveGameData.UserId != "user_001" {
		t.Errorf("Expected UserId 'user_001', got %s", leaveGameData.UserId)
	}
}

func TestControlEventBuilder_KickPlayer(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	request, err := builder.KickPlayer("user_002",
		WithOpUserId("admin_001"),
		WithReason("违规行为"),
	)

	if err != nil {
		t.Fatalf("KickPlayer failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeKickPlayer {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeKickPlayer, request.Type)
	}

	var kickPlayerData PublishControlEventRequestKickPlayer
	err = json.Unmarshal([]byte(request.Data), &kickPlayerData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if kickPlayerData.UserId != "user_002" {
		t.Errorf("Expected UserId 'user_002', got %s", kickPlayerData.UserId)
	}

	if *kickPlayerData.OpUserId != "admin_001" {
		t.Errorf("Expected OpUserId 'admin_001', got %s", *kickPlayerData.OpUserId)
	}

	if *kickPlayerData.Reason != "违规行为" {
		t.Errorf("Expected Reason '违规行为', got %s", *kickPlayerData.Reason)
	}
}

func TestControlEventBuilder_StartGame(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	request, err := builder.StartGame(
		WithForce(true),
		WithStartOpUserId("admin_001"),
		WithStartEXT("game_mode:normal"),
	)

	if err != nil {
		t.Fatalf("StartGame failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeStartGame {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeStartGame, request.Type)
	}

	var startGameData PublishControlEventRequestStartGame
	err = json.Unmarshal([]byte(request.Data), &startGameData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if !*startGameData.Force {
		t.Errorf("Expected Force true, got %v", *startGameData.Force)
	}

	if *startGameData.OpUserId != "admin_001" {
		t.Errorf("Expected OpUserId 'admin_001', got %s", *startGameData.OpUserId)
	}

	if *startGameData.StartEXT != "game_mode:normal" {
		t.Errorf("Expected StartEXT 'game_mode:normal', got %s", *startGameData.StartEXT)
	}
}

func TestControlEventBuilder_IssueProps(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	propDetails := []PublishControlEventRequestIssuePropsDetail{
		NewPropDetail("prop_001", 10, WithDuration(3600), WithDurationReset(false)),
		NewPropDetail("prop_002", 5),
	}

	request, err := builder.IssueProps("user_001", "unique_123", propDetails,
		WithExtra("活动奖励"),
	)

	if err != nil {
		t.Fatalf("IssueProps failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeUserItemGrant {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeUserItemGrant, request.Type)
	}

	var issuePropsData PublishControlEventRequestIssueProps
	err = json.Unmarshal([]byte(request.Data), &issuePropsData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if issuePropsData.UserId != "user_001" {
		t.Errorf("Expected UserId 'user_001', got %s", issuePropsData.UserId)
	}

	if issuePropsData.UniqueID != "unique_123" {
		t.Errorf("Expected UniqueID 'unique_123', got %s", issuePropsData.UniqueID)
	}

	if *issuePropsData.Extra != "活动奖励" {
		t.Errorf("Expected Extra '活动奖励', got %s", *issuePropsData.Extra)
	}

	if len(issuePropsData.Details) != 2 {
		t.Errorf("Expected 2 details, got %d", len(issuePropsData.Details))
	}

	// 验证第一个道具详情
	if issuePropsData.Details[0].PropID != "prop_001" {
		t.Errorf("Expected PropID 'prop_001', got %s", issuePropsData.Details[0].PropID)
	}

	if issuePropsData.Details[0].Num != 10 {
		t.Errorf("Expected Num 10, got %d", issuePropsData.Details[0].Num)
	}

	if *issuePropsData.Details[0].Duration != 3600 {
		t.Errorf("Expected Duration 3600, got %d", *issuePropsData.Details[0].Duration)
	}

	if *issuePropsData.Details[0].DurationReset {
		t.Errorf("Expected DurationReset false, got %v", *issuePropsData.Details[0].DurationReset)
	}
}

func TestControlEventBuilder_ChangeUserIdentity(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	request, err := builder.ChangeUserIdentity("user_001", luksdkmodels.IdentityAdmin)

	if err != nil {
		t.Fatalf("ChangeUserIdentity failed: %v", err)
	}

	if request.Type != luksdkmodels.ControlEventTypeChangePlayerIdentity {
		t.Errorf("Expected type %d, got %d", luksdkmodels.ControlEventTypeChangePlayerIdentity, request.Type)
	}

	var changeIdentityData PublishControlEventRequestChangeUserIdentity
	err = json.Unmarshal([]byte(request.Data), &changeIdentityData)
	if err != nil {
		t.Fatalf("Failed to unmarshal Data: %v", err)
	}

	if changeIdentityData.UserId != "user_001" {
		t.Errorf("Expected UserId 'user_001', got %s", changeIdentityData.UserId)
	}

	if *changeIdentityData.Identity != luksdkmodels.IdentityAdmin {
		t.Errorf("Expected Identity %d, got %d", luksdkmodels.IdentityAdmin, *changeIdentityData.Identity)
	}
}

func TestControlEventBuilder_WithTimestamp(t *testing.T) {
	builder := NewControlEventBuilder(12345, 67890)
	customTimestamp := time.Now().Unix() - 3600 // 1小时前
	builder.WithTimestamp(customTimestamp)

	request, err := builder.LeaveGame("user_001")
	if err != nil {
		t.Fatalf("LeaveGame failed: %v", err)
	}

	if *request.Timestamp != customTimestamp {
		t.Errorf("Expected Timestamp %d, got %d", customTimestamp, *request.Timestamp)
	}
}

func TestNewPropDetail(t *testing.T) {
	// 测试基本道具详情创建
	detail1 := NewPropDetail("prop_001", 5)
	if detail1.PropID != "prop_001" {
		t.Errorf("Expected PropID 'prop_001', got %s", detail1.PropID)
	}
	if detail1.Num != 5 {
		t.Errorf("Expected Num 5, got %d", detail1.Num)
	}

	// 测试带选项的道具详情创建
	detail2 := NewPropDetail("prop_002", 10,
		WithDuration(7200),
		WithDurationReset(true),
	)
	if detail2.PropID != "prop_002" {
		t.Errorf("Expected PropID 'prop_002', got %s", detail2.PropID)
	}
	if detail2.Num != 10 {
		t.Errorf("Expected Num 10, got %d", detail2.Num)
	}
	if *detail2.Duration != 7200 {
		t.Errorf("Expected Duration 7200, got %d", *detail2.Duration)
	}
	if !*detail2.DurationReset {
		t.Errorf("Expected DurationReset true, got %v", *detail2.DurationReset)
	}
}

func TestHelperFunctions(t *testing.T) {
	// 测试辅助函数
	intPtr := int64Ptr(123)
	if *intPtr != 123 {
		t.Errorf("Expected 123, got %d", *intPtr)
	}

	boolPtr := boolPtr(true)
	if !*boolPtr {
		t.Errorf("Expected true, got %v", *boolPtr)
	}

	strPtr := stringPtr("test")
	if *strPtr != "test" {
		t.Errorf("Expected 'test', got %s", *strPtr)
	}
}

func TestMarshalEventData(t *testing.T) {
	data := PublishControlEventRequestLeaveGame{
		UserId: "user_001",
	}

	jsonStr, err := marshalEventData(data)
	if err != nil {
		t.Fatalf("marshalEventData failed: %v", err)
	}

	var unmarshaled PublishControlEventRequestLeaveGame
	err = json.Unmarshal([]byte(jsonStr), &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if unmarshaled.UserId != "user_001" {
		t.Errorf("Expected UserId 'user_001', got %s", unmarshaled.UserId)
	}
}

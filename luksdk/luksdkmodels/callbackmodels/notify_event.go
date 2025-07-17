package callbackmodels

import (
	"encoding/json"

	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"
)

type NotifyEventRequest struct {
	// App ID
	AppID int64 `json:"app_id"`
	// 游戏 ID
	GameID int64 `json:"game_id"`
	// 房间 ID
	RoomID string `json:"room_id"`
	// 游戏局次 ID
	RoundID string `json:"round_id"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
	// 事件类型，根据事件类型解析为对应的事件数据
	Type int `json:"type"`
	// 事件数据
	Data string `json:"data"`
}

type NotifyEventRequestStartGame struct {
	// 操作用户 ID
	OpUserID *string `json:"op_user_id"`
	// 开始游戏透传数据
	StartEXT *string `json:"start_ext"`
	// 游戏开始时间，秒级时间戳
	StartUnixSEC int64 `json:"start_unix_sec"`
	// 参与游戏的玩家 ID 列表
	UserIDS []string `json:"user_ids"`
}

type NotifyEventRequestGameOver struct {
	// 本局消耗货币数
	CostCoins *int64 `json:"cost_coins"`
	// 结束扩展信息，该字段可能包含不同游戏说附带的额外结束信息
	EndEXT *string `json:"end_ext"`
	// 游戏结束类型
	EndType *int64 `json:"end_type,omitempty"`
	// 游戏结束时间，秒级时间戳
	EndUnixSEC int64 `json:"end_unix_sec"`
	// 操作用户 ID
	OpUserID *string `json:"op_user_id"`
	// 开始游戏透传数据
	StartEXT *string `json:"start_ext"`
	// 游戏开始时间，秒级时间戳
	StartUnixSEC int64 `json:"start_unix_sec"`
	// 玩家游戏结果
	UserResults []NotifyEventRequestGameOverUserResult `json:"user_results"`
}

type NotifyEventRequestGameOverUserResult struct {
	// 是否逃跑
	Escape *bool `json:"escape"`
	// 玩家排名
	Rank int64 `json:"rank"`
	// 玩家得分
	Score *int64 `json:"score"`
	// 玩家状态
	Status int `json:"status"`
	// 是否托管
	Trust *bool `json:"trust"`
}

type NotifyEventRoomUserChange struct {
	// 是否游戏进行中
	Gaming *bool `json:"gaming"`
	// 观众用户 ID 列表
	ObUserIDS []string `json:"ob_user_ids"`
	// 玩家准备状态，string => boolean
	PlayerState map[string]bool `json:"player_state"`
	// 变更类型
	Type int `json:"type"`
	// 变更的用户 ID 列表
	UserIDS []string `json:"user_ids"`
}

type NotifyEventRoomSettingChange struct {
	// 操作用户 ID
	OpUserID *string `json:"op_user_id"`
	// 房间设置
	Setting *string `json:"setting"`
}

type NotifyEventRoomSeatSync struct {
	// 座位用户 ID 集合，string（转 stirng 的座位号） => string
	SeatUserIDS map[int]string `json:"seat_user_ids"`
}

type NotifyEventRoomGameFeature struct {
	// 游戏特色事件 JSON
	Feature *string `json:"feature"`
}

type NotifyEventRoomPropEquip struct {
	// 装配的道具 ID
	EquippedPropID *string `json:"equipped_prop_id"`
	// 卸下的道具 ID
	UnequippedPropID *string `json:"unequipped_prop_id"`
	// 用户 ID
	UserID string `json:"user_id"`
}

type NotifyEventResponse struct {
	// 请求状态码，当值为 0 时表示请求成功
	Code int `json:"code"`
	// 请求状态说明
	Msg *string `json:"msg"`
}

func (req *NotifyEventRequest) ParseStartGameData() (*NotifyEventRequestStartGame, error) {
	var startGame NotifyEventRequestStartGame
	err := json.Unmarshal([]byte(req.Data), &startGame)
	if err != nil {
		return nil, err
	}
	return &startGame, nil
}

func (req *NotifyEventRequest) ParseGameEndData() (*NotifyEventRequestGameOver, error) {
	var gameEnd NotifyEventRequestGameOver
	err := json.Unmarshal([]byte(req.Data), &gameEnd)
	if err != nil {
		return nil, err
	}
	return &gameEnd, nil
}

func (req *NotifyEventRequest) ParseRoomUserChangeData() (*NotifyEventRoomUserChange, error) {
	var roomUserChange NotifyEventRoomUserChange
	err := json.Unmarshal([]byte(req.Data), &roomUserChange)
	if err != nil {
		return nil, err
	}
	return &roomUserChange, nil
}

func (req *NotifyEventRequest) ParseRoomSettingChangeData() (*NotifyEventRoomSettingChange, error) {
	var roomSettingChange NotifyEventRoomSettingChange
	err := json.Unmarshal([]byte(req.Data), &roomSettingChange)
	if err != nil {
		return nil, err
	}
	return &roomSettingChange, nil
}

func (req *NotifyEventRequest) ParseRoomSeatSyncData() (*NotifyEventRoomSeatSync, error) {
	var roomSeatSync NotifyEventRoomSeatSync
	err := json.Unmarshal([]byte(req.Data), &roomSeatSync)
	if err != nil {
		return nil, err
	}
	return &roomSeatSync, nil
}

func (req *NotifyEventRequest) ParseRoomGameFeatureData() (*NotifyEventRoomGameFeature, error) {
	var roomGameFeature NotifyEventRoomGameFeature
	err := json.Unmarshal([]byte(req.Data), &roomGameFeature)
	if err != nil {
		return nil, err
	}
	return &roomGameFeature, nil
}

func (req *NotifyEventRequest) ParseRoomPropEquipData() (*NotifyEventRoomPropEquip, error) {
	var roomPropEquip NotifyEventRoomPropEquip
	err := json.Unmarshal([]byte(req.Data), &roomPropEquip)
	if err != nil {
		return nil, err
	}
	return &roomPropEquip, nil
}

func (req *NotifyEventRequest) Response() *NotifyEventResponse {
	return &NotifyEventResponse{}
}

func (resp *NotifyEventResponse) WithError(err error) *NotifyEventResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *NotifyEventResponse) WithCode(code int) *NotifyEventResponse {
	resp.Code = code
	return resp
}

func (resp *NotifyEventResponse) WithMsg(msg string) *NotifyEventResponse {
	resp.Msg = &msg
	return resp
}

package callbackmodels

import (
	"encoding/json"

	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"
)

type NotifyGameRequest struct {
	// App ID
	AppId int64 `json:"c_id"`
	// 通知数据
	Data string `json:"data"`
	// 房间透传字段，每个用户加入房间时采用最新值作为透传内容
	EXT *string `json:"ext"`
	// 游戏 ID
	GameId int64 `json:"g_id"`
	// 通知类型枚举，需根据不同的通知类型，解析为对应的通知数据
	NotifyType int64 `json:"notify_type"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type NotifyGameRequestGameStartBeforeData struct {
	// 游戏当前设置
	GameSetting *string `json:"game_setting"`
	// 操作
	NotifyAction int64 `json:"notify_action"`
	// 是否已准备，字典键指向用户 ID
	PlayerReadyStatus map[string]bool `json:"player_ready_status"`
	// 房间 ID
	RoomID string `json:"room_id"`
	// 游戏局次 ID
	RoundID string `json:"round_id"`
}

type NotifyGameRequestGameRunning struct {
	// 操作
	NotifyAction int64 `json:"notify_action"`
	// 玩家数量
	PlayerNum *int64 `json:"player_num"`
	// 玩家 ID 列表
	PlayerUids []string `json:"player_uids"`
	// 房间 ID
	RoomID string `json:"room_id"`
	// 游戏局次 ID
	RoundID string `json:"round_id"`
}

type NotifyGameRequestGameEnd struct {
	// 是否强制结束游戏
	IsForceEnd bool `json:"is_force_end"`
	// 操作
	NotifyAction int64 `json:"notify_action"`
	// 玩家分数，字典键指向用户 ID
	PlayerScore map[string]int64 `json:"player_score"`
	// 排名，根据排名排序的玩家 ID 数组
	Rank []string `json:"rank"`
	// 房间 ID
	RoomID string `json:"room_id"`
	// 游戏局次 ID
	RoundID string `json:"round_id"`
}

type NotifyGameResponse struct {
	// 请求状态码，当值为 0 时表示请求成功
	Code int `json:"code"`
	// 请求状态说明
	Msg *string `json:"msg"`
}

func (req *NotifyGameRequest) ParseGameStartBeforeData() (*NotifyGameRequestGameStartBeforeData, error) {
	var data NotifyGameRequestGameStartBeforeData
	err := json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (req *NotifyGameRequest) ParseGameRunningData() (*NotifyGameRequestGameRunning, error) {
	var data NotifyGameRequestGameRunning
	err := json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (req *NotifyGameRequest) ParseGameEndData() (*NotifyGameRequestGameEnd, error) {
	var data NotifyGameRequestGameEnd
	err := json.Unmarshal([]byte(req.Data), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (req *NotifyGameRequest) Response() *NotifyGameResponse {
	return &NotifyGameResponse{}
}

func (resp *NotifyGameResponse) WithError(err error) *NotifyGameResponse {
	e := luksdkerrors.ConvertError(err)
	resp.Code, resp.Msg = e.Code(), e.MessageP()
	return resp
}

func (resp *NotifyGameResponse) WithCode(code int) *NotifyGameResponse {
	resp.Code = code
	return resp
}

func (resp *NotifyGameResponse) WithMsg(msg string) *NotifyGameResponse {
	resp.Msg = &msg
	return resp
}

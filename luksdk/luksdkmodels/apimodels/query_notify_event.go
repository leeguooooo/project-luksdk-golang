package apimodels

type QueryNotifyEventRequest struct {
	// App ID
	AppId int64 `json:"app_id"`
	// 结束时间，可基于秒级结束时间戳过滤
	EndAt *int64 `json:"end_at"`
	// 游戏 ID
	GameID int64 `json:"game_id"`
	// 页码
	PageNo int64 `json:"page_no"`
	// 每页数量
	PageSize int64 `json:"page_size"`
	// 房间 ID，可基于房间 ID 过滤
	RoomID *string `json:"room_id"`
	// 请求签名
	Sign string `json:"sign"`
	// 开始时间，可基于秒级开始时间戳过滤
	StartAt *int64 `json:"start_at"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
	// 事件类型，可基于事件类型过滤，具体参考游戏通知事件接口
	Type *int64 `json:"type"`
}

type QueryNotifyEventResponse struct {
	// Code 0 表示成功
	Code int `json:"code"`
	// Msg 消息
	Msg *string `json:"msg"`
	// Data 数据
	Data struct {
		// List 列表
		List []QueryNotifyEventResponseEvent `json:"list"`
	} `json:"data"`
}

type QueryNotifyEventResponseEvent struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

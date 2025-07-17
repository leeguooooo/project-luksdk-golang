package apimodels

type GetGameServiceListRequest struct {
	// App ID
	AppId int64 `json:"c_id"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type GetGameServiceListResponse struct {
	// 请求状态码，当值为 0 时表示请求成功
	Code int `json:"code"`
	// 响应数据
	Data struct {
		GameList []GetGameServiceListResponseData `json:"game_list"`
	} `json:"data"`
	// 请求状态说明
	Msg *string `json:"msg"`
}

type GetGameServiceListResponseData struct {
	// 游戏图标
	Icon string `json:"g_icon"`
	// 游戏 ID
	GameId int64 `json:"g_id"`
	// 游戏名称
	Name string `json:"g_name"`
	// 游戏在线加载 URL
	URL string `json:"g_url"`
}

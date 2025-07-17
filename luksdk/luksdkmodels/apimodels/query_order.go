package apimodels

type QueryOrderRequest struct {
	// App ID
	AppId int64 `json:"app_id"`
	// 渠道订单 ID，和 game_order_no 二选一
	AppOrderNo *string `json:"app_order_no"`
	// 游戏 ID
	GameID int64 `json:"game_id"`
	// 游戏订单 ID，和 app_order_no 二选一
	GameOrderNo *string `json:"game_order_no"`
	// 随机字符串
	Nonce *string `json:"nonce"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type QueryOrderResponse struct {
	// 请求状态码，当值为 0 时表示请求成功
	Code int `json:"code"`
	// 响应数据
	Data QueryOrderResponseData `json:"data"`
	// 请求状态说明
	Msg *string `json:"msg"`
}

// 响应数据
type QueryOrderResponseData struct {
	// 房间抽成值
	AnchorDraw *int64 `json:"anchor_draw"`
	// App ID
	AppID int64 `json:"app_id"`
	// 渠道订单 ID
	AppOrderID string `json:"app_order_id"`
	// 订单奖励金额
	CoinsAward *int64 `json:"coins_award"`
	// 订单消耗金额
	CoinsCost *int64 `json:"coins_cost"`
	// 官方抽成值
	CoinsOfficialDraw *int64 `json:"coins_official_draw"`
	// 订单时间，秒级时间戳
	CreateTime int64 `json:"create_time"`
	// 订单扩展信息
	EXT *string `json:"ext"`
	// 订单盈余
	Gain *int64 `json:"gain"`
	// 游戏订单 ID
	GameOrderID string `json:"game_order_id"`
	// 订单道具 ID
	ItemID *string `json:"item_id"`
	// 道具数量
	Num *int64 `json:"num"`
	// 游戏支付通知状态
	PayGameStatus int64 `json:"pay_game_status"`
	// 订单状态
	PayStatus int64 `json:"pay_status"`
	// 用户 ID
	UserID string `json:"user_id"`
}

package luksdk

import (
	"encoding/json"
	"strings"
)

type NotifyType = int
type Action = int

const (
	NotifyTypeStartBefore NotifyType = iota + 1 // 游戏开始前状态
	NotifyTypeGaming                            // 游戏开始中状态
	NotifyTypeEnd                               // 游戏结束状态
)

const (
	ActionJoinGame      Action = iota + 1 // 加入游戏操作
	ActionExitGame                        // 退出游戏操作
	ActionSettingGame                     // 设置游戏操作
	ActionKickOut                         // 踢人操作
	ActionStartGame                       // 开始游戏操作
	ActionPrepare                         // 准备操作
	ActionCancelPrepare                   // 取消准备操作
	ActionGameEnd                         // 游戏结束操作
)

type GetChannelTokenRequest struct {
	CId       int    `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`                     // 渠道ID
	CUid      string `json:"c_uid" form:"c_uid" uri:"c_uid" xml:"c_uid"`                 // 渠道用户id
	Code      string `json:"code" form:"code" uri:"code" xml:"code"`                     // 渠道的用户临时的code
	Timestamp int64  `json:"timestamp" form:"timestamp" uri:"timestamp" xml:"timestamp"` // 时间戳
	Sign      string `json:"sign" form:"sign" uri:"sign" xml:"sign"`                     // 加密签名
}

type GetChannelTokenResponse struct {
	Token    string `json:"token"`     // 用户令牌
	LeftTime int64  `json:"left_time"` // 剩余秒数(避免时区问题用剩余秒数) 注意：不能为0
}

type RefreshChannelTokenRequest struct {
	CId       int    `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`                     // 渠道ID
	CUid      string `json:"c_uid" form:"c_uid" uri:"c_uid" xml:"c_uid"`                 // 渠道用户id
	Token     string `json:"token" form:"token" uri:"token" xml:"token"`                 // sdk平台令牌
	Timestamp int64  `json:"timestamp" form:"timestamp" uri:"timestamp" xml:"timestamp"` // 时间戳
	Sign      string `json:"sign" form:"sign" uri:"sign" xml:"sign"`                     // 加密签名
	LeftTime  int64  `json:"left_time" form:"left_time" uri:"left_time" xml:"left_time"` // 设置剩余秒数
}

type RefreshChannelTokenResponse struct {
	Token    string `json:"token"`     // 用户令牌
	LeftTime int64  `json:"left_time"` // 剩余秒数(避免时区问题用剩余秒数) 注意：不能为0
}

type GetChannelUserInfoRequest struct {
	CId   int    `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`     // 渠道ID
	CUid  string `json:"c_uid" form:"c_uid" uri:"c_uid" xml:"c_uid"` // 渠道用户id
	Token string `json:"token" form:"token" uri:"token" xml:"token"` // 用户token
	Sign  string `json:"sign" form:"sign" uri:"sign" xml:"sign"`     // 加密签名
}

type GetChannelUserInfoResponse struct {
	CUid   string `json:"c_uid"`  // 渠道用户id
	Name   string `json:"name"`   // 用户昵称
	Avatar string `json:"avatar"` // 用户头像
	Coins  int64  `json:"coins"`  // 用户金币
}

type CreateChannelOrderRequest struct {
	Sign string                            `json:"sign" form:"sign" uri:"sign" xml:"sign"` // 加密签名
	Data []*CreateChannelOrderRequestEntry `json:"data" form:"data" uri:"data" xml:"data"` // 订单条目
}

type CreateChannelOrderRequestEntry struct {
	CId         int    `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`                                     // 渠道ID
	CUid        string `json:"c_uid" form:"c_uid" uri:"c_uid" xml:"c_uid"`                                 // 渠道用户id
	CRoomId     string `json:"c_room_id" form:"c_room_id" uri:"c_room_id" xml:"c_room_id"`                 // 渠道房间id
	GId         int    `json:"g_id" form:"g_id" uri:"g_id" xml:"g_id"`                                     // 游戏id
	CoinsCost   int64  `json:"coins_cost" form:"coins_cost" uri:"coins_cost" xml:"coins_cost"`             // 下单金额
	ScoreCost   int64  `json:"score_cost" form:"score_cost" uri:"score_cost" xml:"score_cost"`             // 下单积分
	GameOrderId string `json:"game_order_id" form:"game_order_id" uri:"game_order_id" xml:"game_order_id"` // 游戏方生成的唯一id
	Token       string `json:"token" form:"token" uri:"token" xml:"token"`                                 // 用户token
	Timestamp   int64  `json:"timestamp" form:"timestamp" uri:"timestamp" xml:"timestamp"`                 // 时间戳
}

type CreateChannelOrderResponse []*CreateChannelOrderResponseEntry

type CreateChannelOrderResponseEntry struct {
	CUid    string `json:"c_uid"`      // 渠道用户id
	OrderId string `json:"order_id"`   // 渠道订单id
	Coins   int64  `json:"coins_cost"` // 用户当前金币
	Status  int    `json:"status"`     // 付款状态 1成功  0失败
}

type NotifyChannelOrderRequest struct {
	Sign string                            `json:"sign" form:"sign" uri:"sign" xml:"sign"` // 加密签名
	Data []*NotifyChannelOrderRequestEntry `json:"data" form:"data" uri:"data" xml:"data"`
}

type NotifyChannelOrderRequestEntry struct {
	CId         int    `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`                                     // 渠道ID
	CUid        string `json:"c_uid" form:"c_uid" uri:"c_uid" xml:"c_uid"`                                 // 渠道用户id
	GId         int    `json:"g_id" form:"g_id" uri:"g_id" xml:"g_id"`                                     // 游戏id
	GameOrderId string `json:"game_order_id" form:"game_order_id" uri:"game_order_id" xml:"game_order_id"` // 游戏方生成的唯一id
	Token       string `json:"token" form:"token" uri:"token" xml:"token"`                                 // 用户token
	CoinsCost   int64  `json:"coins_cost" form:"coins_cost" uri:"coins_cost" xml:"coins_cost"`             // 下注时消耗的金币
	CoinsAward  int64  `json:"coins_award" form:"coins_award" uri:"coins_award" xml:"coins_award"`         // 下注开出的金币 可能为0或者负数,即没获胜的情况下
	ScoreCost   int64  `json:"score_cost" form:"score_cost" uri:"score_cost" xml:"score_cost"`             // 下注时消耗的积分
	ScoreAward  int64  `json:"score_award" form:"score_award" uri:"score_award" xml:"score_award"`         // 下注开出的积分 可能为0或者负数,即没获胜的情况下
	Timestamp   int64  `json:"timestamp" form:"timestamp" uri:"timestamp" xml:"timestamp"`                 // 时间戳
}

type NotifyChannelOrderResponse []*NotifyChannelOrderResponseEntry

type NotifyChannelOrderResponseEntry struct {
	CUid    string `json:"c_uid"`    // 渠道用户id
	OrderId string `json:"order_id"` // 渠道订单id
	Coins   int64  `json:"coins"`    // 当前用户剩下金币
	Score   int64  `json:"score"`    // 当前用户剩下积分
}

type NotifyGameRequest struct {
	CId        int        `json:"c_id" form:"c_id" uri:"c_id" xml:"c_id"`                             // 渠道ID
	GId        int        `json:"g_id" form:"g_id" uri:"g_id" xml:"g_id"`                             // 游戏id
	NotifyType NotifyType `json:"notify_type" form:"notify_type" uri:"notify_type" xml:"notify_type"` // 游戏通知状态，参考data数据说明
	Ext        string     `json:"ext" form:"ext" uri:"ext" xml:"ext"`                                 // 渠道方透传数据/游戏房间拉起时传入游戏客户端，后续房间的通知会透传该数据,可选项,如果不接sdk接口，默认为空
	Data       string     `json:"data" form:"data" uri:"data" xml:"data"`                             // 游戏数据，以notify_type的类型做相对应解析
	Timestamp  int64      `json:"timestamp" form:"timestamp" uri:"timestamp" xml:"timestamp"`         // 时间戳
	Sign       string     `json:"sign" form:"sign" uri:"sign" xml:"sign"`                             // 加密签名
}

type NotifyGameRequestStartBefore struct {
	RoomId            int             `json:"room_id"`             // 房间id
	RoundId           int             `json:"round_id"`            // 回合ID
	PlayerReadyStatus map[string]bool `json:"player_ready_status"` // 玩家准备状态 playerId:isReady
	NotifyAction      Action          `json:"notify_action"`       // 游戏通知操作，具体的游戏操作 参考游戏通知操作说明
	GameSetting       string          `json:"game_setting"`        // 游戏当前设置/具体设置信息参考对应游戏,可选项,部分游戏没有游戏设置
}

type NotifyGameRequestGaming struct {
	RoomId       int      `json:"room_id"`       // 房间id
	RoundId      int      `json:"round_id"`      // 回合ID
	PlayerNum    int      `json:"player_num"`    // 玩家数量
	PlayerUids   []string `json:"player_uids"`   // 玩家uids数组
	NotifyAction Action   `json:"notify_action"` // 游戏通知操作，具体的游戏操作 参考游戏通知操作说明
}

type NotifyGameRequestEnd struct {
	RoomId       int      `json:"room_id"`       // 房间id
	RoundId      int      `json:"round_id"`      // 回合ID
	Rank         []string `json:"rank"`          // 排名 playerId
	IsForceEnd   bool     `json:"is_force_end"`  // 是否为强制结束
	NotifyAction Action   `json:"notify_action"` // 游戏通知操作，具体的游戏操作 参考游戏通知操作说明
}

type NotifyGameResponse struct{}

func (req *NotifyGameRequest) GetStartBefore() (*NotifyGameRequestStartBefore, error) {
	var data = new(NotifyGameRequestStartBefore)
	return data, json.Unmarshal([]byte(req.Data), data)
}

func (req *NotifyGameRequest) GetGaming() (*NotifyGameRequestGaming, error) {
	var data = new(NotifyGameRequestGaming)
	return data, json.Unmarshal([]byte(req.Data), data)
}

func (req *NotifyGameRequest) GetEnd() (*NotifyGameRequestEnd, error) {
	var data = new(NotifyGameRequestEnd)
	return data, json.Unmarshal([]byte(req.Data), data)
}

type RequestHandler[Q, T any] func(request Q) (T, error)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// WithError 设置响应的错误信息
func (r *Response[T]) WithError(err error, msg ...string) *Response[T] {
	if err == nil {
		return r
	}
	r.Code = errorMap[err]
	r.Msg = err.Error()
	if len(msg) > 0 {
		r.Msg = strings.Join(append([]string{r.Msg}, msg...), ", ")
	}
	if r.Code == 0 {
		r.Code = -1
	}
	return r
}

// WithData 设置响应的数据
func (r *Response[T]) WithData(data T) *Response[T] {
	r.Data = data
	if r.Code == 0 {
		r.Msg = "成功"
	}
	return r
}

// Suc 判断是否成功
func (r *Response[T]) Suc() bool {
	return r.Code == 0
}

func generateHandler[Req, Res any](signSecret, requestSign string, request Req, successHandler ...RequestHandler[Req, Res]) *Response[Res] {
	verify := signature(signSecret, request)
	response := new(Response[Res])
	if verify != requestSign {
		return response.WithError(ErrInvalidSignature, requestSign, verify)
	}

	var err error
	for _, h := range successHandler {
		response.Data, err = h(request)
		if err != nil {
			return response.WithError(ErrChannelDataException, err.Error())
		}
	}

	return response
}

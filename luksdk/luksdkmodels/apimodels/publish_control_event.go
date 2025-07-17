package apimodels

import (
	"encoding/json"
	"time"

	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels"
)

type PublishControlEventRequest struct {
	// App ID
	AppId int64 `json:"app_id"`
	// 游戏 ID
	GameId int64 `json:"game_id"`
	// 房间 ID
	RoomId *string `json:"room_id"`
	// 请求签名
	Sign string `json:"sign"`
	// 秒级时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`
	// 事件类型枚举
	Type luksdkmodels.ControlEventType `json:"type"`
	// 事件数据 JSON
	Data string `json:"data"`
}

type PublishControlEventRequestJoinGame struct {
	// 自动开始游戏人数，当座位人数满足时将自动触发开始游戏
	AutoStartNum *int64 `json:"auto_start_num"`
	// 是否准备
	Ready *bool `json:"ready"`
	// 座位号，需要加入的座位号，当为 0 时将顺序选择空位入座
	Seat *int64 `json:"seat"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestLeaveGame struct {
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestChangeReadyStatus struct {
	// 是否准备
	IsPrepare *bool `json:"is_prepare"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestKickPlayer struct {
	// 操作者用户 ID，当填写该值后将会校验操作人状态及权限是否满足
	OpUserId *string `json:"op_user_id"`
	// 踢出原因
	Reason *string `json:"reason"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestStartGame struct {
	// 是否强制开始，若强制开始，将忽略玩家准备状态
	Force *bool `json:"force"`
	// 操作者用户 ID，当填写该值后将会校验操作人状态及权限是否满足
	OpUserId *string `json:"op_user_id"`
	// 开始游戏透传参数，该参数将在游戏开始通知事件和游戏结束通知事件进行携带
	StartEXT *string `json:"start_ext"`
}

type PublishControlEventRequestForceCloseGame struct {
	// 是否清理游戏座位，若清理，将会结束游戏后对所有座位上玩家执行离开游戏操作
	ClearSeat *bool `json:"clear_seat"`
	// 操作者用户 ID，当填写该值后将会校验操作人状态及权限是否满足
	OpUserId *string `json:"op_user_id"`
}

type PublishControlEventRequestChangeRoomSetting struct {
	// 操作者用户 ID，当填写该值后将会校验操作人状态及权限是否满足
	OpUserId *string `json:"op_user_id"`
	// 房间设置
	RoomSetting *string `json:"room_setting"`
}

type PublishControlEventRequestChangeUserIdentity struct {
	// 服务端用户身份，默认情况下游戏将采用客户端接口带入的用户身份，当使用服务端身份后，客户端身份将无效
	Identity *luksdkmodels.Identity `json:"identity,omitempty"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestRoomSeatSync struct {
}

type PublishControlEventRequestRefreshUserInfo struct {
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestQuickStartGame struct {
	// 变更房间设置
	Setting *string `json:"setting"`
	// 是否开始游戏
	StartGame *bool `json:"start_game"`
	// 加入游戏的用户 ID 列表，需确保用户已经拉起游戏
	UserIdS []string `json:"user_ids"`
}

type PublishControlEventRequestIssueProps struct {
	// 发放道具详情
	Details []PublishControlEventRequestIssuePropsDetail `json:"details"`
	// 附加信息，被记录的附加信息，可用于三方协查
	Extra *string `json:"extra"`
	// 用作幂等的唯一 ID
	UniqueID string `json:"unique_id"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestIssuePropsDetail struct {
	// 有效时长，秒级的有效时长，仅特定道具有效，当小于 0 时表示为永久
	Duration *int64 `json:"duration"`
	// 是否重置有效时长，仅时效性道具有效，移除后再发放（用于重置过期时间）
	DurationReset *bool `json:"duration_reset"`
	// 发放数量
	Num int64 `json:"num"`
	// 道具 ID
	PropID string `json:"prop_id"`
}

type PublishControlEventRequestFetchBagStatus struct {
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestFetchBagStatusResponse struct {
	// 背包道具列表
	Props []PublishControlEventRequestFetchBagStatusResponseProp `json:"props"`
}

type PublishControlEventRequestFetchBagStatusResponseProp struct {
	// 过期时间，秒级时间戳，当值为 -1 时表示永久
	ExpireTime *int64 `json:"expire_time"`
	// 是否已装备
	IsEquipped *bool `json:"is_equipped"`
	// 道具数量
	Num int64 `json:"num"`
	// 道具 ID
	PropID string `json:"prop_id"`
	// 道具类型
	Type int64 `json:"type"`
}

type PublishControlEventRequestQueryIssuePropStatus struct {
	// 唯一 ID
	UniqueID string `json:"unique_id"`
}

type PublishControlEventRequestQueryIssuePropStatusResponse struct {
	// App ID
	AppId int64 `json:"app_id"`
	// 记录创建秒级时间戳
	CreatedTime *int64 `json:"created_time"`
	// 道具详情列表
	Details []PublishControlEventRequestQueryIssuePropStatusResponseDetailsItem `json:"details"`
	// 附加信息，被记录的附加信息，用于三方协查
	Extra *string `json:"extra"`
	// 游戏 ID
	GameId int64 `json:"game_id"`
	// 发放状态
	Status int64 `json:"status"`
	// 唯一 ID
	UniqueID string `json:"unique_id"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestQueryIssuePropStatusResponseDetailsItem struct {
	// 有效时长，秒级的有效时长，仅特定道具有效，当小于 0 时表示为永久
	Duration *int64 `json:"duration"`
	// 是否重置有效时长，仅时效性道具有效，移除后再发放（用于重置过期时间）
	DurationReset *bool `json:"duration_reset"`
	// 发放数量
	Num int64 `json:"num"`
	// 道具 ID
	PropID string `json:"prop_id"`
}

type PublishControlEventRequestEquippedProp struct {
	// 道具 ID
	EquippedPropID string `json:"equipped_prop_id"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventRequestUnequippedProp struct {
	// 道具 ID
	UnequippedPropID string `json:"unequipped_prop_id"`
	// 用户 ID
	UserId string `json:"user_id"`
}

type PublishControlEventResponse struct {
	// 响应码
	Code int `json:"code"`
	// 响应信息
	Msg string `json:"msg"`
	// 响应数据
	Data json.RawMessage `json:"data"`
}

// ControlEventBuilder 控制事件构造器
type ControlEventBuilder struct {
	request *PublishControlEventRequest
}

// NewControlEventBuilder 创建控制事件构造器
func NewControlEventBuilder(appId, gameId int64) *ControlEventBuilder {
	return &ControlEventBuilder{
		request: &PublishControlEventRequest{
			AppId:     appId,
			GameId:    gameId,
			Timestamp: int64Ptr(time.Now().Unix()),
		},
	}
}

// WithRoomId 设置房间ID
func (b *ControlEventBuilder) WithRoomId(roomId string) *ControlEventBuilder {
	b.request.RoomId = &roomId
	return b
}

// WithSign 设置签名
func (b *ControlEventBuilder) WithSign(sign string) *ControlEventBuilder {
	b.request.Sign = sign
	return b
}

// WithTimestamp 设置时间戳
func (b *ControlEventBuilder) WithTimestamp(timestamp int64) *ControlEventBuilder {
	b.request.Timestamp = &timestamp
	return b
}

// Build 构建最终的请求
func (b *ControlEventBuilder) Build() PublishControlEventRequest {
	return *b.request
}

// 辅助函数
func int64Ptr(v int64) *int64 {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func stringPtr(v string) *string {
	return &v
}

// marshalEventData 将事件数据序列化为JSON字符串
func marshalEventData(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JoinGame 创建加入游戏控制事件
func (b *ControlEventBuilder) JoinGame(userId string, options ...func(*PublishControlEventRequestJoinGame)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestJoinGame{
		UserId: userId,
	}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeJoinGame
	b.request.Data = dataJSON
	return b.request, nil
}

// WithAutoStartNum 设置自动开始游戏人数
func WithAutoStartNum(num int64) func(*PublishControlEventRequestJoinGame) {
	return func(req *PublishControlEventRequestJoinGame) {
		req.AutoStartNum = &num
	}
}

// WithReady 设置是否准备
func WithReady(ready bool) func(*PublishControlEventRequestJoinGame) {
	return func(req *PublishControlEventRequestJoinGame) {
		req.Ready = &ready
	}
}

// WithSeat 设置座位号
func WithSeat(seat int64) func(*PublishControlEventRequestJoinGame) {
	return func(req *PublishControlEventRequestJoinGame) {
		req.Seat = &seat
	}
}

// LeaveGame 创建离开游戏控制事件
func (b *ControlEventBuilder) LeaveGame(userId string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestLeaveGame{
		UserId: userId,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeLeaveGame
	b.request.Data = dataJSON
	return b.request, nil
}

// ChangeReadyStatus 创建改变准备状态控制事件
func (b *ControlEventBuilder) ChangeReadyStatus(userId string, isPrepare bool) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestChangeReadyStatus{
		UserId:    userId,
		IsPrepare: &isPrepare,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeChangeReady
	b.request.Data = dataJSON
	return b.request, nil
}

// KickPlayer 创建踢出玩家控制事件
func (b *ControlEventBuilder) KickPlayer(userId string, options ...func(*PublishControlEventRequestKickPlayer)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestKickPlayer{
		UserId: userId,
	}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeKickPlayer
	b.request.Data = dataJSON
	return b.request, nil
}

// WithOpUserId 设置操作者用户ID
func WithOpUserId(opUserId string) func(*PublishControlEventRequestKickPlayer) {
	return func(req *PublishControlEventRequestKickPlayer) {
		req.OpUserId = &opUserId
	}
}

// WithReason 设置踢出原因
func WithReason(reason string) func(*PublishControlEventRequestKickPlayer) {
	return func(req *PublishControlEventRequestKickPlayer) {
		req.Reason = &reason
	}
}

// StartGame 创建开始游戏控制事件
func (b *ControlEventBuilder) StartGame(options ...func(*PublishControlEventRequestStartGame)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestStartGame{}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeStartGame
	b.request.Data = dataJSON
	return b.request, nil
}

// WithForce 设置是否强制开始
func WithForce(force bool) func(*PublishControlEventRequestStartGame) {
	return func(req *PublishControlEventRequestStartGame) {
		req.Force = &force
	}
}

// WithStartOpUserId 设置开始游戏操作者用户ID
func WithStartOpUserId(opUserId string) func(*PublishControlEventRequestStartGame) {
	return func(req *PublishControlEventRequestStartGame) {
		req.OpUserId = &opUserId
	}
}

// WithStartEXT 设置开始游戏透传参数
func WithStartEXT(startEXT string) func(*PublishControlEventRequestStartGame) {
	return func(req *PublishControlEventRequestStartGame) {
		req.StartEXT = &startEXT
	}
}

// ForceCloseGame 创建强制结束游戏控制事件
func (b *ControlEventBuilder) ForceCloseGame(options ...func(*PublishControlEventRequestForceCloseGame)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestForceCloseGame{}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeForceEndGame
	b.request.Data = dataJSON
	return b.request, nil
}

// WithClearSeat 设置是否清理游戏座位
func WithClearSeat(clearSeat bool) func(*PublishControlEventRequestForceCloseGame) {
	return func(req *PublishControlEventRequestForceCloseGame) {
		req.ClearSeat = &clearSeat
	}
}

// WithCloseOpUserId 设置强制结束游戏操作者用户ID
func WithCloseOpUserId(opUserId string) func(*PublishControlEventRequestForceCloseGame) {
	return func(req *PublishControlEventRequestForceCloseGame) {
		req.OpUserId = &opUserId
	}
}

// ChangeRoomSetting 创建修改房间设置控制事件
func (b *ControlEventBuilder) ChangeRoomSetting(roomSetting string, options ...func(*PublishControlEventRequestChangeRoomSetting)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestChangeRoomSetting{
		RoomSetting: &roomSetting,
	}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeChangeRoomSetting
	b.request.Data = dataJSON
	return b.request, nil
}

// WithRoomSettingOpUserId 设置房间设置操作者用户ID
func WithRoomSettingOpUserId(opUserId string) func(*PublishControlEventRequestChangeRoomSetting) {
	return func(req *PublishControlEventRequestChangeRoomSetting) {
		req.OpUserId = &opUserId
	}
}

// ChangeUserIdentity 创建修改用户身份控制事件
func (b *ControlEventBuilder) ChangeUserIdentity(userId string, identity luksdkmodels.Identity) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestChangeUserIdentity{
		UserId:   userId,
		Identity: &identity,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeChangePlayerIdentity
	b.request.Data = dataJSON
	return b.request, nil
}

// SyncRoomSeat 创建同步房间座位控制事件
func (b *ControlEventBuilder) SyncRoomSeat() (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestRoomSeatSync{}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeSyncRoomSeat
	b.request.Data = dataJSON
	return b.request, nil
}

// RefreshUserInfo 创建刷新用户信息控制事件
func (b *ControlEventBuilder) RefreshUserInfo(userId string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestRefreshUserInfo{
		UserId: userId,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeRefreshUserInfo
	b.request.Data = dataJSON
	return b.request, nil
}

// QuickStartGame 创建快捷开始游戏控制事件
func (b *ControlEventBuilder) QuickStartGame(userIds []string, options ...func(*PublishControlEventRequestQuickStartGame)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestQuickStartGame{
		UserIdS: userIds,
	}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeQuickStartGame
	b.request.Data = dataJSON
	return b.request, nil
}

// WithQuickStartSetting 设置快捷开始游戏房间设置
func WithQuickStartSetting(setting string) func(*PublishControlEventRequestQuickStartGame) {
	return func(req *PublishControlEventRequestQuickStartGame) {
		req.Setting = &setting
	}
}

// WithQuickStartGame 设置是否开始游戏
func WithQuickStartGame(startGame bool) func(*PublishControlEventRequestQuickStartGame) {
	return func(req *PublishControlEventRequestQuickStartGame) {
		req.StartGame = &startGame
	}
}

// IssueProps 创建发放道具控制事件
func (b *ControlEventBuilder) IssueProps(userId, uniqueID string, details []PublishControlEventRequestIssuePropsDetail, options ...func(*PublishControlEventRequestIssueProps)) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestIssueProps{
		UserId:   userId,
		UniqueID: uniqueID,
		Details:  details,
	}

	for _, option := range options {
		option(data)
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeUserItemGrant
	b.request.Data = dataJSON
	return b.request, nil
}

// WithExtra 设置附加信息
func WithExtra(extra string) func(*PublishControlEventRequestIssueProps) {
	return func(req *PublishControlEventRequestIssueProps) {
		req.Extra = &extra
	}
}

// NewPropDetail 创建道具详情
func NewPropDetail(propID string, num int64, options ...func(*PublishControlEventRequestIssuePropsDetail)) PublishControlEventRequestIssuePropsDetail {
	detail := PublishControlEventRequestIssuePropsDetail{
		PropID: propID,
		Num:    num,
	}

	for _, option := range options {
		option(&detail)
	}

	return detail
}

// WithDuration 设置道具有效时长
func WithDuration(duration int64) func(*PublishControlEventRequestIssuePropsDetail) {
	return func(detail *PublishControlEventRequestIssuePropsDetail) {
		detail.Duration = &duration
	}
}

// WithDurationReset 设置是否重置有效时长
func WithDurationReset(reset bool) func(*PublishControlEventRequestIssuePropsDetail) {
	return func(detail *PublishControlEventRequestIssuePropsDetail) {
		detail.DurationReset = &reset
	}
}

// FetchBagStatus 创建获取背包状态控制事件
func (b *ControlEventBuilder) FetchBagStatus(userId string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestFetchBagStatus{
		UserId: userId,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeGetUserItemBag
	b.request.Data = dataJSON
	return b.request, nil
}

// QueryIssuePropStatus 创建查询道具发放状态控制事件
func (b *ControlEventBuilder) QueryIssuePropStatus(uniqueID string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestQueryIssuePropStatus{
		UniqueID: uniqueID,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeQueryItemGrantStatus
	b.request.Data = dataJSON
	return b.request, nil
}

// EquipProp 创建装备道具控制事件
func (b *ControlEventBuilder) EquipProp(userId, propID string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestEquippedProp{
		UserId:         userId,
		EquippedPropID: propID,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeAssemblePlayerItem
	b.request.Data = dataJSON
	return b.request, nil
}

// UnequipProp 创建卸下道具控制事件
func (b *ControlEventBuilder) UnequipProp(userId, propID string) (*PublishControlEventRequest, error) {
	data := &PublishControlEventRequestUnequippedProp{
		UserId:           userId,
		UnequippedPropID: propID,
	}

	dataJSON, err := marshalEventData(data)
	if err != nil {
		return nil, err
	}

	b.request.Type = luksdkmodels.ControlEventTypeUnassemblePlayerItem
	b.request.Data = dataJSON
	return b.request, nil
}

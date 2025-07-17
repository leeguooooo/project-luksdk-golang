package luksdkmodels

type Identity = int64

const (
	IdentityOwner  Identity = 1
	IdentityAdmin  Identity = 2
	IdentityNormal Identity = 3
)

type ControlEventType = int64

// 事件数据 JSON
//
// 控制玩家加入游戏
//
// 控制玩家离开游戏
//
// 改变玩家准备状态
//
// 踢出特定玩家
//
// 开始游戏
//
// 强制结束游戏
//
// 修改房间设置
//
// 修改玩家身份
//
// 主动拉取同步房间座位事件
//
// 刷新用户信息
//
// 快捷开始游戏
//
// 用户道具发放
//
// 获取用户背包状态
//
// 查询道具发放状态
//
// 装配玩家道具
//
// 卸下玩家道具
const (
	// 控制玩家加入游戏
	ControlEventTypeJoinGame ControlEventType = 1
	// 控制玩家离开游戏
	ControlEventTypeLeaveGame ControlEventType = 2
	// 改变玩家准备状态
	ControlEventTypeChangeReady ControlEventType = 3
	// 踢出特定玩家
	ControlEventTypeKickPlayer ControlEventType = 4
	// 开始游戏
	ControlEventTypeStartGame ControlEventType = 5
	// 强制结束游戏
	ControlEventTypeForceEndGame ControlEventType = 6
	// 修改房间设置
	ControlEventTypeChangeRoomSetting ControlEventType = 7
	// 修改玩家身份
	ControlEventTypeChangePlayerIdentity ControlEventType = 8
	// 主动拉取同步房间座位事件
	ControlEventTypeSyncRoomSeat ControlEventType = 9
	// 刷新用户信息
	ControlEventTypeRefreshUserInfo ControlEventType = 10
	// 快捷开始游戏
	ControlEventTypeQuickStartGame ControlEventType = 11
)

const (
	// 用户道具发放
	ControlEventTypeUserItemGrant ControlEventType = 1000
	// 获取用户背包状态
	ControlEventTypeGetUserItemBag ControlEventType = 1001
	// 查询道具发放状态
	ControlEventTypeQueryItemGrantStatus ControlEventType = 1002
	// 装配玩家道具
	ControlEventTypeAssemblePlayerItem ControlEventType = 1003
	// 卸下玩家道具
	ControlEventTypeUnassemblePlayerItem ControlEventType = 1004
)

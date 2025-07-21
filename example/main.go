package main

import (
	"errors"
	"github.com/CFGameTech/project-luksdk-golang/luksdk"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkerrors"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels/apimodels"
	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels/callbackmodels"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
)

func main() {
	// 构建 LukSDK 实例
	sdk := luksdk.NewLukSDKWithConfigurators(luksdk.ConfiguratorFN(func(config *luksdk.Config) {
		config.WithAppId(0).WithDomain("https://xxx.xxx.xx")
	}))

	// 普通请求
	resp, err := sdk.Apis.GetGameServiceList(apimodels.GetGameServiceListRequest{
		AppId:     0,   // AppId 为 0 的情况下默认采用配置值
		Sign:      "",  // 签名未填写的情况下默认生成
		Timestamp: nil, // 时间戳为 nil 的情况下默认采用当前时间的秒级时间戳
	})

	switch {
	case err != nil:
		// 请求错误
	case resp.Code != 0:
		// 请求异常，可根据错误码和业务场景执行不同逻辑
	}

	// 回调请求
	router := gin.New()
	router.POST("/prefix/get_channel_user_info", func(context *gin.Context) {
		var request callbackmodels.GetChannelUserInfoRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			// 以标准错误码响应，并携带额外错误信息
			context.JSON(http.StatusOK, new(callbackmodels.GetChannelUserInfoResponse).WithError(luksdkerrors.LukSDKErrorParamError.With("参数解析失败", err)))
			return
		}

		// 假设业务失败
		if rand.IntN(100) < 50 {
			// 将采用 luksdkerrors.LukSDKErrorInternalError 错误
			context.JSON(http.StatusOK, new(callbackmodels.GetChannelUserInfoResponse).WithError(errors.New("内部异常")))
			return
		}

		// 响应结果
		context.JSON(http.StatusOK, new(callbackmodels.GetChannelUserInfoResponse).
			WithUserId(request.UserId).
			WithName("nickname").
			WithAvatar("https://aaa.bbb.ccc/avatar.png"))
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}

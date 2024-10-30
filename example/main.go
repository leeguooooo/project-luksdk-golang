package main

import (
	luksdk "github.com/CFGameTech/project-luksdk-golang"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	sdk := luksdk.New("fa7ad21fdbe10218024f88538a86")
	app := gin.New()
	defer func() {
		if err := app.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	app.POST("/sdk/get_channel_token", func(context *gin.Context) {
		var request = new(luksdk.GetChannelTokenRequest)
		var response = new(luksdk.Response[*luksdk.GetChannelTokenResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("get_channel_token", "request", request, "response", response)
			return
		}

		response = sdk.GetChannelToken(request, func(request *luksdk.GetChannelTokenRequest) (*luksdk.GetChannelTokenResponse, error) {
			resp := &luksdk.GetChannelTokenResponse{
				Token:    "my-token",
				LeftTime: 7200,
			}
			return resp, nil
		})

		context.JSON(200, response)
		slog.Info("get_channel_token", "request", request, "response", response)
	})

	app.POST("/sdk/refresh_channel_token", func(context *gin.Context) {
		var request = new(luksdk.RefreshChannelTokenRequest)
		var response = new(luksdk.Response[*luksdk.RefreshChannelTokenResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("refresh_channel_token", "request", request, "response", response)
			return
		}

		response = sdk.RefreshChannelToken(request, func(request *luksdk.RefreshChannelTokenRequest) (*luksdk.RefreshChannelTokenResponse, error) {
			resp := &luksdk.RefreshChannelTokenResponse{
				Token:    "my-token",
				LeftTime: 7200,
			}
			return resp, nil
		})

		context.JSON(200, response)
		slog.Info("refresh_channel_token", "request", request, "response", response)
	})

	app.POST("/sdk/get_channel_user_info", func(context *gin.Context) {
		var request = new(luksdk.GetChannelUserInfoRequest)
		var response = new(luksdk.Response[*luksdk.GetChannelUserInfoResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("get_channel_user_info", "request", request, "response", response)
			return
		}

		response = sdk.GetChannelUserInfo(request, func(request *luksdk.GetChannelUserInfoRequest) (*luksdk.GetChannelUserInfoResponse, error) {
			resp := &luksdk.GetChannelUserInfoResponse{
				CUid:   request.CUid,
				Name:   "my-name",
				Avatar: "",
				Coins:  100000,
			}
			return resp, nil
		})

		context.JSON(200, response)
		slog.Info("get_channel_user_info", "request", request, "response", response)
	})

	app.POST("/sdk/create_channel_order", func(context *gin.Context) {
		var request = new(luksdk.CreateChannelOrderRequest)
		var response = new(luksdk.Response[luksdk.CreateChannelOrderResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("create_channel_order", "request", request, "response", response)
			return
		}

		response = sdk.CreateChannelOrder(request, func(request *luksdk.CreateChannelOrderRequest) (luksdk.CreateChannelOrderResponse, error) {
			var resp luksdk.CreateChannelOrderResponse
			for _, datum := range request.Data {
				resp = append(resp, &luksdk.CreateChannelOrderResponseEntry{
					CUid:    datum.CUid,
					OrderId: datum.GameOrderId,
					Coins:   100000,
					Status:  1,
				})
			}
			return resp, nil
		})

		context.JSON(200, response)
		slog.Info("create_channel_order", "request", request, "response", response)
	})

	app.POST("/sdk/notify_channel_order", func(context *gin.Context) {
		var request = new(luksdk.NotifyChannelOrderRequest)
		var response = new(luksdk.Response[luksdk.NotifyChannelOrderResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("notify_channel_order", "request", request, "response", response)
			return
		}

		response = sdk.NotifyChannelOrder(request, func(request *luksdk.NotifyChannelOrderRequest) (luksdk.NotifyChannelOrderResponse, error) {
			var resp luksdk.NotifyChannelOrderResponse
			for _, datum := range request.Data {
				resp = append(resp, &luksdk.NotifyChannelOrderResponseEntry{
					CUid:    datum.CUid,
					OrderId: datum.GameOrderId,
					Coins:   100000,
					Score:   100000,
				})
			}
			return resp, nil
		})

		context.JSON(200, response)
		slog.Info("notify_channel_order", "request", request, "response", response)
	})

	app.POST("/sdk/notify_game", func(context *gin.Context) {
		var request = new(luksdk.NotifyGameRequest)
		var response = new(luksdk.Response[*luksdk.NotifyGameResponse])
		if err := context.ShouldBind(request); err != nil {
			context.JSON(400, response.WithError(err))
			slog.Info("notify_game", "request", request, "response", response)
			return
		}

		response = sdk.NotifyGame(request, func(request *luksdk.NotifyGameRequest) (*luksdk.NotifyGameResponse, error) {
			return new(luksdk.NotifyGameResponse), nil
		})

		context.JSON(200, response)
		slog.Info("notify_game", "request", request, "response", response)
	})
}

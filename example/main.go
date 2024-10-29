package main

import (
	luksdk "github.com/CFGameTech/project-luksdk-golang"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	sdk := luksdk.New("@")
	app := gin.New()
	defer func() {
		if err := app.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	app.POST("/go/get_channel_token", func(context *gin.Context) {
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
}

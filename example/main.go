package main

import (
	"encoding/json"
	"fmt"
	luksdk "github.com/kercylan98/server-golang-sdk"
)

func main() {
	// 初始化 SDK
	sdk := luksdk.New("123456")

	// 来自 SDK 请求的参数结构
	request := &luksdk.GetChannelTokenRequest{
		CId:       1000,
		CUid:      "123456789",
		Timestamp: 167456789,
	}
	request.Sign = sdk.GenerateSignature(request)

	// 处理请求
	resp := sdk.GetChannelToken(request, func(request *luksdk.GetChannelTokenRequest) (*luksdk.GetChannelTokenResponse, error) {
		// 业务逻辑
		return &luksdk.GetChannelTokenResponse{
			Token:    "token", // 生成 Token
			LeftTime: 7200,    // 设置 Token 过期时间
		}, nil
	})

	// 将 resp 作为 JSON 写入 HTTP 响应
	b, _ := json.Marshal(resp)
	fmt.Println(string(b))
}

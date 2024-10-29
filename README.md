# 介绍
本项目为 Golang 版本的 LukSDK，可直接引入使用，其中提供了需接入接口的通用实现，仅需结合业务逻辑将其返回即可。

> 仅需将 HTTP 请求转换为对应结构体后调用相关函数并填充返回值即可，关于参数的校验等行为交由 SDK 内部处理。

# Go Mod
可通过以下方式引入依赖

```shell
go get -u github.com/CFGameTech/project-luksdk-golang
```

# 示例代码
```go
package main

import (
	"encoding/json"
	"fmt"
	luksdk "github.com/kerylan98/server-golang-sdk"
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
```
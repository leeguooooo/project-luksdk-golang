# LukSDK for Go

本项目为 Golang 版本的 LukSDK，提供游戏平台接入所需的数据结构定义、API 调用接口和控制事件构造工具。

## 主要功能

- **API 接口封装**: 提供标准化的请求/响应结构体定义
- **控制事件构造**: 使用构造器模式简化复杂事件参数的组装
- **回调结构定义**: 提供完整的回调数据结构，供开发者自行实现业务逻辑
- **类型安全**: 完整的 Go 类型定义，减少运行时错误

## 安装

```bash
go get -u github.com/CFGameTech/project-luksdk-golang
```

## 示例代码

参考示例文件：[SDK 初始化和 API 调用示例](./example/main.go)

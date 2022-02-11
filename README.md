企业微信应用消息发送
============
[![Go Report Card](https://goreportcard.com/badge/github.com/dongfg/notify)](https://goreportcard.com/report/github.com/dongfg/notify)
> [官方API文档](https://work.weixin.qq.com/api/doc#90001/90143/90372)

## 使用
### Use as a library
```shell script
go get github.com/dongfg/notify
```
简单示例：
```go
package main

import (
	"fmt"
	"github.com/dongfg/notify"
)

func main() {
	n := notify.New("", 1000001, "") // your config

	result, err := n.Send(notify.MessageReceiver{
		ToUser: "@all",
	}, notify.Text{Content: "Simple Message"}, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

### 命令行
从 [Release](https://github.com/dongfg/notify/releases) 页面下载二进制文件，或者通过 go 命令安装:
```shell script
GO111MODULE="off" go get github.com/dongfg/notify/cmd/notify
```
全局参数，支持环境变量设置：
```shell script
GLOBAL OPTIONS:
   --corpID value     企业ID [$CORP_ID]
   --agentID value    应用agentID [$AGENT_ID]
   --appSecret value  应用secret [$APP_SECRET]
```
发送文本消息：
```shell script
CORP_ID=your corp id AGENT_ID=your agent id APP_SECRET=your app secret notify text "to user" "message content"
```
详细参数:
```shell script
USAGE:
   notify text [command options] [arguments...]

OPTIONS:
   --toUser value, --user value, -u value    目标用户
   --toParty value, --party value, -p value  目标部门
   --toTag value, --tag value, -t value      目标标签
   --message value, -m value                 消息内容
```
## 测试情况
- [x] Text
- [x] Image
- [x] Voice
- [x] Video
- [x] File
- [x] TextCard
- [x] News
- [x] MpNews
- [x] Markdown: 仅支持在企业微信查看
- [x] TaskCard: 仅支持在企业微信查看
- [ ] MiniProgram: 未测试

## 命令行支持情况
- [x] Text
- [x] Image
- [x] Voice
- [x] Video
- [x] File
- [x] TextCard 
- [x] News
- [ ] MpNews
- [x] Markdown
- [ ] TaskCard
- [ ] MiniProgram

## Todo-List
- [ ] token 异常过期处理
- [x] 命令行发送其他类型消息
- [ ] 命令行 token 缓存
- [x] 命令行 completion

企业微信应用消息发送
============
[![Go Report Card](https://goreportcard.com/badge/github.com/dongfg/notify)](https://goreportcard.com/report/github.com/dongfg/notify)
> [企业微信官方API文档](https://work.weixin.qq.com/api/doc#90001/90143/90372)

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

全局参数：

```text
Flags:
      --config string      config file (default is $HOME/.notify.yaml)
      --corpID string      企业ID，https://work.weixin.qq.com/wework_admin/frame#profile
      --agentID int        应用agentID，应用页面查看
      --appSecret string   应用secret，应用页面查看
  -u, --user string        指定接收消息的成员，成员ID列表，多个接收者用‘|’分隔，最多支持1000个。特殊情况：指定为 @all，则向该企业应用的全部成员发送
  -p, --party string       指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数
  -t, --tag string         指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数
  -v, --verbose            verbose mode
```

全局参数配置文件 ``.notify.yaml`` 内容参考，可通过 ``--config`` 指定，若未指定会从 ``$HOME`` 及当前目录查找:

```yaml
appSecret: LTyXNttXXXXXXXXXXXXXyqtQ9Uw
agentID: 1000002
corpID: ww7XXXXXXXXXd9
user: "@all"
party: "1|2|3"
tag: "t1|t2|t3"
```

详细参数:

```text
企业微信应用消息发送

Usage:
  notify [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  file        发送文件消息
  help        Help about any command
  image       发送图片消息
  markdown    发送 markdown 消息
  news        发送图文消息
  text        发送文本消息
  textcard    发送文本卡片消息
  upload      上传临时素材
  video       发送视频消息
  voice       发送语音消息

Flags:
      --config string      config file (default is $HOME/.notify.yaml)
      --corpID string      企业ID，https://work.weixin.qq.com/wework_admin/frame#profile
      --agentID int        应用agentID，应用页面查看
      --appSecret string   应用secret，应用页面查看
  -u, --user string        指定接收消息的成员，成员ID列表，多个接收者用‘|’分隔，最多支持1000个。特殊情况：指定为 @all，则向该企业应用的全部成员发送
  -p, --party string       指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数
  -t, --tag string         指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数
  -v, --verbose            verbose mode
  -h, --help               help for notify

Use "notify [command] --help" for more information about a command.
```

发送文本消息：

```shell script
notify text "some text from command line"
```

发送文件：

```shell script
notify file ./README.md
```

发送文本卡片：

```shell
notify textcard --title 恭喜你中奖了 --description 明天不用上班 --url www.baidu.com --btntxt 让我看看
```

其他消息请查看帮助 ``notify [command] --help``

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

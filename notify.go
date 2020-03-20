/**
企业微信应用消息发送接口封装，接口文档见：
https://work.weixin.qq.com/api/doc/90001/90143/90372
id转译说明：
https://work.weixin.qq.com/api/doc/90001/90143/90372#id%E8%BD%AC%E8%AF%91%E8%AF%B4%E6%98%8E
*/
package notify

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MessageReceiver 消息接收者 ToUser、ToParty、ToTag 至少一个
type MessageReceiver struct {
	// ToUser 非必填，成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）。
	// 特殊情况：指定为@all，则向关注该企业应用的全部成员发送
	ToUser string `json:"touser"`
	// ToParty 非必填，指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为”@all”时忽略本参数
	ToParty string `json:"toparty"`
	// ToTag 非必填，指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。 当touser为”@all”时忽略本参数
	ToTag string `json:"totag"`
}

// MessageOptions 消息配置包括加密、id转译、重复检查等
type MessageOptions struct {
	// Safe 表示是否是保密消息，默认否
	Safe bool `json:"safe"`
	// EnableIdTrans 表示是否开启id转译，默认否
	EnableIdTrans bool `json:"enable_id_trans"`
	// EnableDuplicateCheck 表示是否开启重复消息检查，默认否
	EnableDuplicateCheck bool `json:"enable_duplicate_check"`
	// DuplicateCheckInterval 表示是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// TextMessage 文本消息
type TextMessage struct {
	// Content 消息内容，最长不超过2048个字节，超过将截断（支持id转译）
	Content string `json:"content"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	// 图片媒体文件id，可以调用上传临时素材接口获取
	MediaID string `json:"media_id"`
}

// Notify reference to call send method
type Notify struct {
	corpID    string
	agentID   uint
	appSecret string
}

type message struct {
}

// New 创建客户端， corpID 企业ID，在企业信息页面查看, agentID + appSecret 在应用页面查看
func New(corpID string, agentID uint, appSecret string) *Notify {
	return &Notify{
		corpID, agentID, appSecret,
	}
}

// Send message with options to receiver, options can be nil
func (n *Notify) Send(receiver *MessageReceiver, message interface{}, options *MessageOptions) (err error) {
	if receiver == nil {
		err = errors.New("message receiver can not be nil")
		return
	}
	if message == nil {
		err = errors.New("message can not be nil")
		return
	}

	msgBody := make(map[string]interface{})

	if len(receiver.ToUser) > 0 {
		msgBody["touser"] = receiver.ToUser
	}
	if len(receiver.ToParty) > 0 {
		msgBody["toparty"] = receiver.ToParty
	}
	if len(receiver.ToTag) > 0 {
		msgBody["totag"] = receiver.ToTag
	}
	if len(msgBody) == 0 {
		err = errors.New("message receiver not set, set at least one")
		return
	}

	msgBody["agentid"] = n.agentID
	if options != nil {
		if options.Safe {
			msgBody["safe"] = 1
		}
		if options.EnableIdTrans {
			msgBody["enable_id_trans"] = 1
		}
		if options.EnableDuplicateCheck {
			msgBody["enable_duplicate_check"] = 1
			if options.DuplicateCheckInterval != 0 {
				msgBody["duplicate_check_interval"] = options.DuplicateCheckInterval
			}
		}
	}

	switch v := message.(type) {
	case TextMessage:
		msgBody["msgtype"] = "text"
		msgBody["text"] = v
	case ImageMessage:
		msgBody["msgtype"] = "image"
		msgBody["image"] = v
	default:
		err = fmt.Errorf("unrecognized message type: %T", v)
	}
	jsonBody, err := json.Marshal(msgBody)
	fmt.Println(string(jsonBody))
	return
}

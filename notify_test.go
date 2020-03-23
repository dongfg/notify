package notify

import (
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func newNotifyFromEnv() *Notify {
	corpID, hasCorpID := os.LookupEnv("corpID")
	agentIDStr, hasAgentID := os.LookupEnv("agentID")
	appSecret, hasAppSecret := os.LookupEnv("appSecret")
	agentID, err := strconv.Atoi(agentIDStr)
	if !(hasCorpID && hasAgentID && hasAppSecret) || err != nil {
		panic("please set environment correctly")
	}

	return New(corpID, int64(agentID), appSecret)
}

func TestNotify_Send(t *testing.T) {
	type args struct {
		receiver MessageReceiver
		message  interface{}
		options  *MessageOptions
	}
	tests := []struct {
		name    string
		args    args
		want    MessageResult
		wantErr bool
	}{
		{
			name:    "UnknownType",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: "Simple String", options: nil},
			want:    MessageResult{},
			wantErr: true,
		},
		{
			name:    "Text",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: Text{Content: "TestNotify_Send_Text"}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name:    "Image",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: Image{MediaID: "2lUfpG6A6TxyH7WJbtArKH3N40q5dF9tmV6Ib2e2tvCutkqEKGxrmExSbwSmLzv2Q"}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name:    "Voice",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: Voice{MediaID: "2yazpc6Y-vSYmF24KP8N9b1jZ5nD9wnvkOR0amyZWn5o"}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name:    "Video",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: Video{MediaID: "23NH1YLqekQ3FPAXN_uxn9A39MItpX2SEgZ7xmaKGyCc-pnqvOM62eH7zVJHZ9Xz2"}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name:    "File",
			args:    args{receiver: MessageReceiver{ToUser: "@all"}, message: File{MediaID: "2WZTACgAdDh2NpUAs0hcPt6tsXBN1lZ7X2JELMxMO4k7auzDOsbOAAl6SO5y4kyh7"}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name: "TextCard",
			args: args{receiver: MessageReceiver{ToUser: "@all"}, message: TextCard{
				Title:       "放假通知",
				Description: "清明节放假通知",
				URL:         "https://work.weixin.qq.com/",
			}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name: "News",
			args: args{receiver: MessageReceiver{ToUser: "@all"}, message: News{Articles: []NewsArticle{
				{
					Title:       "中秋节礼品领取",
					Description: "今年中秋节公司有豪礼相送",
					URL:         "https://work.weixin.qq.com/",
					PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
				},
			}}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name: "MpNews",
			args: args{receiver: MessageReceiver{ToUser: "@all"}, message: MpNews{Articles: []MpNewsArticle{
				{
					Title:            "中秋节礼品领取",
					ThumbMediaID:     "2lUfpG6A6TxyH7WJbtArKH3N40q5dF9tmV6Ib2e2tvCutkqEKGxrmExSbwSmLzv2Q",
					Author:           "UnitTest",
					ContentSourceURL: "https://work.weixin.qq.com/",
					Content:          "今年中秋节公司有豪礼相送",
					Digest:           "今年中秋节公司有豪礼相送",
				},
			}}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
		{
			name: "Markdown",
			args: args{receiver: MessageReceiver{ToUser: "@all"}, message: Markdown{Content: `
您的会议室已经预定，稍后会同步到 *邮箱*
>**事项详情**
>事　项：<font color=\"info\">开会</font>
>组织者：@miglioguan
>参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang
>
>会议室：<font color=\"info\">广州TIT 1楼 301</font>
>日　期：<font color=\"warning\">2018年5月18日</font>
>时　间：<font color=\"comment\">上午9:00-11:00</font>
>
>请准时参加会议。
>
>如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)
`}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},

		{
			name: "TaskCard",
			args: args{receiver: MessageReceiver{ToUser: "@all"}, message: TaskCard{
				Title:       "赵明登的礼物申请",
				Description: "礼品：A31茶具套装<br>用途：赠与小黑科技张总经理",
				TaskID:      "notify_" + strconv.Itoa(rand.Intn(99999)),
				Buttons: []TaskCardButton{
					{
						Key:         "k1",
						Name:        "批准",
						ReplaceName: "已批准",
						Color:       "red",
						IsBold:      true,
					},
					{
						Key:         "k2",
						Name:        "驳回",
						ReplaceName: "已驳回",
					},
				},
			}, options: nil},
			want:    MessageResult{ErrorCode: 0, ErrorMsg: "ok"},
			wantErr: false,
		},
	}
	n := newNotifyFromEnv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := n.Send(tt.args.receiver, tt.args.message, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotify_getToken(t *testing.T) {
	n := newNotifyFromEnv()
	t.Run("ValidConfig", func(t *testing.T) {
		err := n.getToken()
		if err != nil {
			t.Errorf("getToken() error = %v, want no error", err)
		}
	})
}

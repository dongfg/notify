package notify

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		corpID    string
		agentID   int64
		appSecret string
	}
	fakeCorpID := "fakeCorpID"
	fakeAgentID := rand.Int63()
	fakeAppSecret := "fakeAppSecret"
	tests := []struct {
		name string
		args args
		want *Notify
	}{
		{
			name: "Basic",
			args: struct {
				corpID    string
				agentID   int64
				appSecret string
			}{corpID: fakeCorpID, agentID: fakeAgentID, appSecret: fakeAppSecret},
			want: &Notify{corpID: fakeCorpID, agentID: fakeAgentID, appSecret: fakeAppSecret,},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.corpID, tt.args.agentID, tt.args.appSecret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotify_Send(t *testing.T) {
	type fields struct {
		corpID         string
		agentID        int64
		appSecret      string
		token          string
		tokenExpiresAt int64
	}
	type args struct {
		receiver *MessageReceiver
		message  interface{}
		options  *MessageOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MessageResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				corpID:         tt.fields.corpID,
				agentID:        tt.fields.agentID,
				appSecret:      tt.fields.appSecret,
				token:          tt.fields.token,
				tokenExpiresAt: tt.fields.tokenExpiresAt,
			}
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
	type fields struct {
		corpID         string
		agentID        int64
		appSecret      string
		token          string
		tokenExpiresAt int64
	}
	type args struct {
		corpID    string
		appSecret string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				corpID:         tt.fields.corpID,
				agentID:        tt.fields.agentID,
				appSecret:      tt.fields.appSecret,
				token:          tt.fields.token,
				tokenExpiresAt: tt.fields.tokenExpiresAt,
			}
			if err := n.getToken(tt.args.corpID, tt.args.appSecret); (err != nil) != tt.wantErr {
				t.Errorf("getToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotify_sendInternal(t *testing.T) {
	type fields struct {
		corpID         string
		agentID        int64
		appSecret      string
		token          string
		tokenExpiresAt int64
	}
	type args struct {
		msgBody map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MessageResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				corpID:         tt.fields.corpID,
				agentID:        tt.fields.agentID,
				appSecret:      tt.fields.appSecret,
				token:          tt.fields.token,
				tokenExpiresAt: tt.fields.tokenExpiresAt,
			}
			got, err := n.sendInternal(tt.args.msgBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendInternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendInternal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotify_sendMessage(t *testing.T) {
	type fields struct {
		corpID         string
		agentID        int64
		appSecret      string
		token          string
		tokenExpiresAt int64
	}
	type args struct {
		msgBody map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MessageResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				corpID:         tt.fields.corpID,
				agentID:        tt.fields.agentID,
				appSecret:      tt.fields.appSecret,
				token:          tt.fields.token,
				tokenExpiresAt: tt.fields.tokenExpiresAt,
			}
			got, err := n.sendMessage(tt.args.msgBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

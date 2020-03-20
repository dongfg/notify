package main

import (
	"fmt"
	"github.com/dongfg/notify"
)

func main() {
	client := notify.New("11", 22, "33")
	err := client.Send(&notify.MessageReceiver{
		ToUser: "@all",
	}, notify.TextMessage{Content: "I am content"}, &notify.MessageOptions{
		Safe:                   true,
		EnableIdTrans:          true,
		EnableDuplicateCheck:   true,
		DuplicateCheckInterval: 1400,
	})
	if err != nil {
		fmt.Println(err)
	}
}

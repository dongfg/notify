package main

import (
	"fmt"
	"github.com/dongfg/notify"
)

func main() {
	client := notify.New("ww0cb42f7ec6df90f7", 1000002, "eA-NzUtglvuyBIeJQjdJjwOBe29XbvI5XlO2FpSd9BA")
	result, err := client.Send(&notify.MessageReceiver{
		ToUser: "@all",
	}, notify.Text{Content: "I am content"}, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

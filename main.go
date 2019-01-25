package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

// Config struct
type Config struct {
	CorpID     string `json:"corpId"`
	CorpSecret string `json:"corpSecret"`
	AgentID    int    `json:"agentId"`
}

// Message struct
type Message struct {
	ToUser      string       `json:"touser"`
	MsgType     string       `json:"msgtype"`
	AgentID     int          `json:"agentid"`
	Safe        int          `json:"safe"`
	MsgText     *MsgText     `json:"text,omitempty"`
	MsgImage    *MsgImage    `json:"image,omitempty"`
	MsgVoice    *MsgVoice    `json:"voice,omitempty"`
	MsgVideo    *MsgVideo    `json:"video,omitempty"`
	MsgFile     *MsgFile     `json:"file,omitempty"`
	MsgTextCard *MsgTextCard `json:"textcard,omitempty"`
	MsgNews     *MsgNews     `json:"news,omitempty"`
	MsgMpNews   *MsgMpNews   `json:"mpnews,omitempty"`
	MsgMarkdown *MsgMarkdown `json:"markdown,omitempty"`
}

// MsgText MsgType: "text"
type MsgText struct {
	Content string `json:"content"`
}

// MsgImage MsgType: "image"
type MsgImage struct {
	MediaID string `json:"media_id"`
}

// MsgVoice MsgType: "voice"
type MsgVoice struct {
	MediaID string `json:"media_id"`
}

// MsgVideo MsgType: "video"
type MsgVideo struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// MsgFile MsgType: "file"
type MsgFile struct {
	MediaID string `json:"media_id"`
}

// MsgTextCard MsgType: "textcard"
type MsgTextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	BtnTxt      string `json:"btntxt,omitempty"`
}

// MsgNews MsgType: "news"
type MsgNews struct {
	Articles []*NewsArticle `json:"articles"`
}

// NewsArticle MsgNews MsgType: "news"
type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl,omitempty"`
}

// MsgMpNews MsgType: "mpnews"
type MsgMpNews struct {
	Articles []*MpNewsArticle `json:"articles"`
}

// MpNewsArticle MsgNews MsgType: "mpnews"
type MpNewsArticle struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author,omitempty"`
	ContentSourceURL string `json:"content_source_url,omitempty"`
	Content          string `json:"content"`
	Digest           string `json:"digest,omitempty"`
}

// MsgMarkdown MsgType: "markdown"
type MsgMarkdown struct {
	Content string `json:"content"`
}

func main() {
	cli.AppHelpTemplate = `NAME:
	{{.Name}} - {{.Usage}}
USAGE:
	notify --config/-c config.json --message/m message.json
	`

	app := cli.NewApp()
	app.Version = "2019.1"
	app.Name = "notify"
	app.Usage = "企业微信消息推送"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config json",
		},
		cli.StringFlag{
			Name:  "message, m",
			Usage: "message json",
		},
	}

	app.Action = func(c *cli.Context) error {
		if !c.GlobalIsSet("config") || !c.GlobalIsSet("message") {
			cli.ShowAppHelp(c)
			return nil
		}

		configFile, err := os.Open(c.GlobalString("config"))
		if err != nil {
			return nil
		}
		defer configFile.Close()

		configByte, _ := ioutil.ReadAll(configFile)
		var config Config
		json.Unmarshal(configByte, &config)

		message := Message{}
		message.AgentID = config.AgentID
		msgJSON, _ := json.Marshal(message)

		fmt.Println(string(msgJSON))

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

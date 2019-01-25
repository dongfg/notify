package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

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

// AccessToken auth api
type AccessToken struct {
	ErrCode   int    `json:"errcode,omitempty"`
	ErrMsg    string `json:"errmsg,omitempty"`
	Token     string `json:"access_token,omitempty"`
	ExpiresIn int32  `json:"expires_in,omitempty"`
	ExpiresAt int32  `json:"expires_at,omitempty"`
}

var client = &http.Client{Timeout: 10 * time.Second}

func accessToken(config Config) (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	cachePath := usr.HomeDir + "/.notify"
	expired := false
	if _, err := os.Stat(cachePath); !os.IsNotExist(err) {
		cacheFile, err := os.Open(cachePath)
		if err != nil {
			log.Fatal(err)
		}
		defer cacheFile.Close()

		cacheByte, _ := ioutil.ReadAll(cacheFile)
		var accessToken AccessToken
		json.Unmarshal(cacheByte, &accessToken)
		if accessToken.ExpiresAt > int32(time.Now().Unix()) {
			return accessToken.Token, nil
		}
		expired = true
	}

	if _, err := os.Stat(cachePath); expired || os.IsNotExist(err) {
		res, err := client.Get(fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", config.CorpID, config.CorpSecret))
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		accessToken := AccessToken{}
		json.NewDecoder(res.Body).Decode(&accessToken)

		if accessToken.ErrMsg != "ok" {
			panic(accessToken.ErrMsg)
		} else {
			accessToken.ExpiresAt = int32(time.Now().Unix()) + accessToken.ExpiresIn
			accessTokenJSON, _ := json.Marshal(accessToken)

			cacheFile, err := os.Create(cachePath)
			if err != nil {
				log.Fatal(err)
			}
			cacheFile.Write(accessTokenJSON)
			cacheFile.Sync()

			return accessToken.Token, nil
		}
	}

	return "", err
}

func readJSON(path string, target interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileByte, &target)
}

func sendMessage(config Config, message Message) {
	message.AgentID = config.AgentID

	accessToken, err := accessToken(config)
	if err != nil {
		log.Fatal(err)
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(message)

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", accessToken), body)
	res, _ := client.Do(req)
	io.Copy(os.Stdout, res.Body)
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

		var config Config
		if readJSON(c.GlobalString("config"), &config) != nil {
			log.Fatal("fail to parse " + c.GlobalString("config"))
		}

		var message Message
		if readJSON(c.GlobalString("message"), &message) != nil {
			log.Fatal("fail to parse " + c.GlobalString("message"))
		}

		sendMessage(config, message)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

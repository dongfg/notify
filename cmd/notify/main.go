package main

import (
	"errors"
	"fmt"
	"github.com/dongfg/notify"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Version:              "v1.1.0",
		Name:                 "notify",
		HelpName:             "notify",
		Usage:                "企业微信应用消息发送",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "corpID",
				Usage:   "企业ID",
				EnvVars: []string{"CORP_ID"},

				Required: true,
			},
			&cli.Int64Flag{
				Name:     "agentID",
				Usage:    "应用agentID",
				EnvVars:  []string{"AGENT_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "appSecret",
				Usage:    "应用secret",
				EnvVars:  []string{"APP_SECRET"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "text",
				Usage: "发送文本消息",
				Before: func(c *cli.Context) error {
					if c.Args().Len() == 2 {
						err := c.Set("toUser", c.Args().Get(0))
						if err != nil {
							return err
						}
						err = c.Set("message", c.Args().Get(1))
						return err
					}
					if !c.IsSet("toUser") && !c.IsSet("toParty") && !c.IsSet("toTag") {
						return errors.New("请至少指定一个发送目标")
					}
					if !c.IsSet("message") {
						return errors.New("请指定消息内容")
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "toUser",
						Usage:   "目标用户",
						Aliases: []string{"user", "u"},
					},
					&cli.StringFlag{
						Name:    "toParty",
						Usage:   "目标部门",
						Aliases: []string{"party", "p"},
					},
					&cli.StringFlag{
						Name:    "toTag",
						Usage:   "目标标签",
						Aliases: []string{"tag", "t"},
					},
					&cli.StringFlag{
						Name:    "message",
						Usage:   "消息内容",
						Aliases: []string{"m"},
					},
				},
				Action: func(c *cli.Context) error {
					n := notify.New(c.String("corpID"), c.Int64("agentID"), c.String("appSecret"))
					result, err := n.Send(notify.MessageReceiver{
						ToUser:  c.String("toUser"),
						ToParty: c.String("toParty"),
						ToTag:   c.String("toTag"),
					}, notify.Text{Content: c.String("message")}, nil)
					if err == nil {
						fmt.Println(result.ErrorMsg)
					}
					return err
				},
			},
			{
				Name:  "send",
				Usage: "通过配置文件发送消息",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "file",
						Usage:    "配置文件",
						Aliases:  []string{"f"},
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("WIP")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

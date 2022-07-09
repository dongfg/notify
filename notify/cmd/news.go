package cmd

import (
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// newsCmd represents the news command
var newsCmd = &cobra.Command{
	Use:   "news",
	Short: "发送图文消息",
	Long:  `发送图文消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return err
		}
		picurl, err := cmd.Flags().GetString("picurl")
		if err != nil {
			return err
		}
		return sendMessage(notify.News{Articles: []notify.NewsArticle{
			{
				Title:       title,
				Description: description,
				URL:         url,
				PicURL:      picurl,
			},
		}})
	},
}

func init() {
	rootCmd.AddCommand(newsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newsCmd.PersistentFlags().String("title", "", "标题，不超过128个字节，超过会自动截断")
	newsCmd.PersistentFlags().String("description", "", "非必填。描述，不超过512个字节，超过会自动截断")
	newsCmd.PersistentFlags().String("url", "", "非必填。点击后跳转的链接")
	newsCmd.PersistentFlags().String("picurl", "", "非必填。图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150")
	_ = newsCmd.MarkPersistentFlagRequired("title")

	newsCmd.Flags().SortFlags = false
	newsCmd.PersistentFlags().SortFlags = false

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

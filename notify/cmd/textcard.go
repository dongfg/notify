package cmd

import (
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// textcardCmd represents the textcard command
var textcardCmd = &cobra.Command{
	Use:   "textcard",
	Short: "发送文本卡片消息",
	Long:  `发送文本卡片消息`,
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
		btntxt, err := cmd.Flags().GetString("btntxt")
		if err != nil {
			return err
		}
		return sendMessage(notify.TextCard{
			Title:       title,
			Description: description,
			URL:         url,
			BtnTxt:      btntxt,
		})
	},
}

func init() {
	rootCmd.AddCommand(textcardCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	textcardCmd.PersistentFlags().String("title", "", "标题，不超过128个字节，超过会自动截断")
	textcardCmd.PersistentFlags().String("description", "", "描述，不超过512个字节，超过会自动截断")
	textcardCmd.PersistentFlags().String("url", "", "点击后跳转的链接")
	textcardCmd.PersistentFlags().String("btntxt", "", "非必填。按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断")
	_ = textcardCmd.MarkPersistentFlagRequired("title")
	_ = textcardCmd.MarkPersistentFlagRequired("description")
	_ = textcardCmd.MarkPersistentFlagRequired("url")

	textcardCmd.Flags().SortFlags = false
	textcardCmd.PersistentFlags().SortFlags = false

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textcardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

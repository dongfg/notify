package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// markdownCmd represents the markdown command
var markdownCmd = &cobra.Command{
	Use:   "markdown <markdown content>",
	Short: "发送 markdown 消息",
	Long:  `发送 markdown 消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请输入要发送的MD文本")
		}
		return sendMessage(notify.Markdown{Content: args[0]})
	},
}

func init() {
	rootCmd.AddCommand(markdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

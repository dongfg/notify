package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file <file path>",
	Short: "发送文件消息",
	Long:  `发送文件消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请指定要发送的文件")
		}
		file := args[0]
		media, err := client.Upload(notify.UploadMedia{
			Type: "file",
			Path: file,
		})
		if err != nil {
			return err
		}
		return sendMessage(notify.File{MediaID: media.MediaID})
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video <video path>",
	Short: "发送视频消息",
	Long:  `发送视频消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请指定要发送的视频")
		}
		video := args[0]
		media, err := client.Upload(notify.UploadMedia{
			Type: "video",
			Path: video,
		})
		if err != nil {
			return err
		}
		return sendMessage(notify.Video{MediaID: media.MediaID})
	},
}

func init() {
	rootCmd.AddCommand(videoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

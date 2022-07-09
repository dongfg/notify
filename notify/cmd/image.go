package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image <image path>",
	Short: "发送图片消息",
	Long:  `发送图片消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请指定要发送的图片")
		}
		image := args[0]
		media, err := client.Upload(notify.UploadMedia{
			Type: "image",
			Path: image,
		})
		if err != nil {
			return err
		}
		return sendMessage(notify.Image{MediaID: media.MediaID})
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

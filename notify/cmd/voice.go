package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// voiceCmd represents the voice command
var voiceCmd = &cobra.Command{
	Use:   "voice <voice path>",
	Short: "发送语音消息",
	Long:  `发送语音消息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("请指定要发送的音频")
		}
		voice := args[0]
		media, err := client.Upload(notify.UploadMedia{
			Type: "voice",
			Path: voice,
		})
		if err != nil {
			return err
		}
		return sendMessage(notify.Voice{MediaID: media.MediaID})
	},
}

func init() {
	rootCmd.AddCommand(voiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

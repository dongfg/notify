package cmd

import (
	"fmt"
	"github.com/dongfg/notify"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传临时素材",
	Long:  `上传临时素材`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		file, err := cmd.Flags().GetString("file")
		if err == nil && file != "" {
			return nil
		}
		image, err := cmd.Flags().GetString("image")
		if err == nil && image != "" {
			return nil
		}
		voice, err := cmd.Flags().GetString("voice")
		if err == nil && voice != "" {
			return nil
		}
		video, err := cmd.Flags().GetString("video")
		if err == nil && video != "" {
			return nil
		}
		return fmt.Errorf("请指定要上传的文件, 图片（image）、语音（voice）、视频（video）或普通文件（file）")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		media := notify.UploadMedia{
			Type: "",
			Path: "",
		}
		file, err := cmd.Flags().GetString("file")
		if err == nil && file != "" {
			media.Type = "file"
			media.Path = file
		}
		image, err := cmd.Flags().GetString("image")
		if err == nil && image != "" {
			media.Type = "image"
			media.Path = image
		}
		voice, err := cmd.Flags().GetString("voice")
		if err == nil && voice != "" {
			media.Type = "voice"
			media.Path = voice
		}
		video, err := cmd.Flags().GetString("video")
		if err == nil && video != "" {
			media.Type = "video"
			media.Path = video
		}
		r, err := client.Upload(media)
		if err != nil {
			return err
		}
		if r.ErrorCode == 0 {
			fmt.Println("上传成功, MediaID: " + r.MediaID)
		} else {
			fmt.Println(r.ErrorMsg)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	uploadCmd.PersistentFlags().String("file", "", "普通文件")
	uploadCmd.PersistentFlags().String("image", "", "图片")
	uploadCmd.PersistentFlags().String("voice", "", "语音")
	uploadCmd.PersistentFlags().String("video", "", "视频")

	uploadCmd.Flags().SortFlags = false
	uploadCmd.PersistentFlags().SortFlags = false

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

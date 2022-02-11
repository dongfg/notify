package cmd

import (
	"fmt"
	"github.com/dongfg/notify"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

//  cfgFile global config file path
var cfgFile string

var (
	corpID    string
	agentID   int64
	appSecret string
	client    *notify.Notify
	receiver  notify.MessageReceiver
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "notify",
	Short: "企业微信应用消息发送",
	Long:  `企业微信应用消息发送`,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		client = notify.New(corpID, agentID, appSecret)
		receiver = notify.MessageReceiver{
			ToUser:  viper.GetString("user"),
			ToParty: viper.GetString("party"),
			ToTag:   viper.GetString("tag"),
		}
		if cmd.Use == "upload" {
			return nil
		}

		if viper.GetString("user") == "" && viper.GetString("party") == "" && viper.GetString("tag") == "" {
			return fmt.Errorf(`请指定发送对象：user、party 或 tag`)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.notify.yaml)")

	rootCmd.PersistentFlags().StringVarP(&corpID, "corpID", "", "", "企业ID，https://work.weixin.qq.com/wework_admin/frame#profile")
	rootCmd.PersistentFlags().Int64VarP(&agentID, "agentID", "", 0, "应用agentID，应用页面查看")
	rootCmd.PersistentFlags().StringVarP(&appSecret, "appSecret", "", "", "应用secret，应用页面查看")
	_ = rootCmd.MarkPersistentFlagRequired("corpID")
	_ = rootCmd.MarkPersistentFlagRequired("agentID")
	_ = rootCmd.MarkPersistentFlagRequired("appSecret")

	rootCmd.PersistentFlags().StringP("user", "u", "", "指定接收消息的成员，成员ID列表，多个接收者用‘|’分隔，最多支持1000个。特殊情况：指定为 @all，则向该企业应用的全部成员发送")
	rootCmd.PersistentFlags().StringP("party", "p", "", "指定接收消息的部门，部门ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数")
	rootCmd.PersistentFlags().StringP("tag", "t", "", "指定接收消息的标签，标签ID列表，多个接收者用‘|’分隔，最多支持100个。当 user 为 @all 时忽略本参数")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode")

	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	_ = viper.BindPFlag("corpID", rootCmd.PersistentFlags().Lookup("corpID"))
	_ = viper.BindPFlag("agentID", rootCmd.PersistentFlags().Lookup("agentID"))
	_ = viper.BindPFlag("appSecret", rootCmd.PersistentFlags().Lookup("appSecret"))

	_ = viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	_ = viper.BindPFlag("party", rootCmd.PersistentFlags().Lookup("party"))
	_ = viper.BindPFlag("tag", rootCmd.PersistentFlags().Lookup("tag"))

	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	viper.SetEnvPrefix("NOTIFY")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".notify")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
		postInitCommands(rootCmd.Commands())
	}
}

func postInitCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		presetRequiredFlags(cmd)
		if cmd.HasSubCommands() {
			postInitCommands(cmd.Commands())
		}
	}
}

func presetRequiredFlags(cmd *cobra.Command) {
	_ = viper.BindPFlags(cmd.Flags())
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			_ = cmd.Flags().Set(f.Name, viper.GetString(f.Name))
		}
	})
}

func sendMessage(message interface{}) error {
	r, err := client.Send(receiver, message, nil)
	if err != nil {
		return err
	}
	if r.ErrorCode == 0 {
		fmt.Println("发送成功")
	} else {
		fmt.Println(r.ErrorMsg)
	}
	return nil
}

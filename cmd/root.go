package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wangyw15/Chate/server"
	"github.com/wangyw15/Chate/util"
)

var rootCmd = &cobra.Command{
	Use:   "chate",
	Short: "A tool for getting images from animate onlineshop by jan code",
	Long:  "A tool for getting images from animate onlineshop by jan code",
	Run: func(cmd *cobra.Command, args []string) {
		parsed, err := url.Parse(viper.GetString("proxy"))
		if err == nil {
			util.SetProxy(parsed)
		}
		server.Start(viper.GetString("host"), viper.GetInt32("port"))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	// allows the program to run from explorer
	cobra.MousetrapHelpText = ""

	// config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	// flags
	rootCmd.PersistentFlags().Int32P("port", "p", 8080, "port to listen")
	rootCmd.PersistentFlags().String("host", "127.0.0.1", "host to listen")
	rootCmd.PersistentFlags().String("proxy", "", "proxy to use")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("proxy", rootCmd.PersistentFlags().Lookup("proxy"))
}

package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/wangyw15/Chate/server"
	"github.com/wangyw15/Chate/util"
)

var port int32
var host string
var proxy string

var rootCmd = &cobra.Command{
	Use:   "chate",
	Short: "A tool for getting images from animate onlineshop by jan code",
	Long:  "A tool for getting images from animate onlineshop by jan code",
	Run: func(cmd *cobra.Command, args []string) {
		parsed, err := url.Parse(proxy)
		if err != nil {
			panic(err)
		}
		util.SetProxy(parsed)
		server.Start(port)
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
	rootCmd.PersistentFlags().Int32VarP(&port, "port", "p", 8080, "port to listen")
	rootCmd.PersistentFlags().StringVar(&host, "host", "127.0.0.1", "host to listen")
	rootCmd.PersistentFlags().StringVar(&proxy, "proxy", "", "proxy to use")
}

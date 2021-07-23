package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	gitlabToken string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "help",
	Short: "tool to remove users from all gitlab repositories and groups accessible with token",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.AutomaticEnv()
	RootCmd.PersistentFlags().StringVarP(&gitlabToken, "token", "t", viper.GetString("gitlab_token"), "gitlab access token")
}

package cmd

import (
	"github.com/allanger/gitlab-remove-user/users"
	"github.com/spf13/cobra"
)

var (
	user string
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "get user id by email or username",
	Run: func(cmd *cobra.Command, args []string) {
		users.Search(user, gitlabToken)
	},
}

func init() {
	searchCmd.Flags().StringVarP(&user, "user", "u", "", "gitlab user email or username")
	RootCmd.AddCommand(searchCmd)
}

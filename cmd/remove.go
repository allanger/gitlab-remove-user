package cmd

import (
	"github.com/allanger/gitlab-remove-user/remove"
	"github.com/spf13/cobra"
)

var (
	dryrun = true
	userID int
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove user from gitlab",
	Long: `
	use --dry-run to check first
`,

	Run: func(cmd *cobra.Command, args []string) {
		remove.Remove(userID, dryrun, gitlabToken)
	},
}

func init() {
	removeCmd.Flags().IntVarP(&userID, "user", "u", -1, "gitlab user id")
	removeCmd.Flags().BoolVar(&dryrun, "dry", true, "use for dry run")
	RootCmd.AddCommand(removeCmd)
}

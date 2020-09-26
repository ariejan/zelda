package cmd

import (
	"github.com/spf13/cobra"
)

var authorizeCmd = &cobra.Command{
	Use:     "authorize",
	Aliases: []string{"a", "auth", "token"},
	Short:   "Retrieve a new API token",
	Long:    `This will request a new API token. You only need to use this if you require your token. All commands handle authorization themselves.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := API.FetchToken()
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

func init() {
	RootCmd.AddCommand(authorizeCmd)
}

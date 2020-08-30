package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var webhooksRootCmd = &cobra.Command{
	Use:     "webhooks",
	Aliases: []string{"w", "webh", "wh"},
	Short:   "Manage thy webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var webhooksGetCmd = &cobra.Command{
	Use:     "list [installationID]",
	Aliases: []string{"l", "ls"},
	Short:   "Retrieve an overview of your webhooks for this installation.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])

		response, err := API.FetchWebhooks(installationID)
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

var webhookGetCmd = &cobra.Command{
	Use:     "get [installationID] [webhookID]",
	Aliases: []string{"g"},
	Short:   "Retrieve info about a specific webhook",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])
		webhookID := parseIntArg(args[1])

		response, err := API.FetchWebhook(installationID, webhookID)
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

var webhookCreateCmd = &cobra.Command{
	Use:     "create [installationID] [endpoint] [token]",
	Aliases: []string{"g"},
	Short:   "Create a new webhook for the specified installation with `endpoint` and `token`",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])

		response, err := API.CreateWebhook(installationID, args[1], args[2])
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

var webhookDeleteCmd = &cobra.Command{
	Use:     "delete [installationID] [webhookID]",
	Aliases: []string{"d", "del"},
	Short:   "Delete the specified webhook",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])
		webhookID := parseIntArg(args[1])

		_, err := API.DeleteWebhook(installationID, webhookID)
		if err != nil {
			er(err)
		}
	},
}

func init() {
	webhooksRootCmd.AddCommand(webhooksGetCmd)
	webhooksRootCmd.AddCommand(webhookGetCmd)
	webhooksRootCmd.AddCommand(webhookCreateCmd)
	webhooksRootCmd.AddCommand(webhookDeleteCmd)

	RootCmd.AddCommand(webhooksRootCmd)
}

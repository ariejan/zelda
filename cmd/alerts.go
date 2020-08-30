package cmd

import (
	"github.com/spf13/cobra"
)

var alertsCmd = &cobra.Command{
	Use:     "alerts [installationID]",
	Aliases: []string{"a", "alert"},
	Short:   "Retrieve an overview of elements with a not-normal status for the specified installation",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])

		response, err := API.FetchAlerts(installationID)
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

func init() {
	RootCmd.AddCommand(alertsCmd)
}

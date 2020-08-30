package cmd

import (
	"github.com/spf13/cobra"
)

var zonesCmd = &cobra.Command{
	Use:     "zones [installationID]",
	Aliases: []string{"z", "zone"},
	Short:   "Retrieve an overview of all the zones for the specified installation.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])

		response, err := API.FetchZones(installationID)
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

func init() {
	RootCmd.AddCommand(zonesCmd)
}

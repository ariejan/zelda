package cmd

import (
	"github.com/spf13/cobra"
)

var elementsCmd = &cobra.Command{
	Use:     "elements [installationID] [zoneID]",
	Aliases: []string{"e", "elem"},
	Short:   "Retrieve an overview of all the elements for the specified installation and zone.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		installationID := parseIntArg(args[0])
		zoneID := parseIntArg(args[1])

		response, err := API.FetchElements(installationID, zoneID)
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

func init() {
	RootCmd.AddCommand(elementsCmd)
}

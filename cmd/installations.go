package cmd

import (
	"github.com/spf13/cobra"
)

var installationsCmd = &cobra.Command{
	Use:     "installations",
	Aliases: []string{"i", "inst"},
	Short:   "Retrieve an overview of your installations.",
	Long:    `This will request and up-to-date list of all of your installations.`,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := API.FetchInstallations()
		if err != nil {
			er(err)
		}

		prettyPrintJSON(response)
	},
}

func init() {
	RootCmd.AddCommand(installationsCmd)
}

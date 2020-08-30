package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping Hertek Connect Link to validate connectivity and authentication",
	Long:  `This will authenticate with the Hertek Connect Link API and send call the ping endpoint to validate you can successfully use the API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sending PING...")

		response, err := API.Ping()
		if err != nil {
			er(err)
		}

		fmt.Printf("==> %s\n", response)
	},
}

func init() {
	RootCmd.AddCommand(pingCmd)
}

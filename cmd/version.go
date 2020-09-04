package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// AppVersion is the current branch of tag for this build.
	AppVersion string
	// BuildTime is the time of the build.
	BuildTime string
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Zelda",
	Long:  `All software has versions. This is Zelda's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Zelda - Hertek Connect Link Command Line Interface %s -- %s\n", AppVersion, BuildTime)
	},
}

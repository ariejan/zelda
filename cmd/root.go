package cmd

import (
	"fmt"
	"os"

	"github.com/ariejan/link/core"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	// RootCmd is where CLI commands come in
	RootCmd = &cobra.Command{
		Use:   "link",
		Short: "Link is a CLI to Hertek Connect Link",
		Long:  `LInk is a CLI to Hertek Connect Link - a valid Integrator account with Hertek Connect is required.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
		},
	}

	// API is what we use to communicate with the API
	API *core.ConnectLinkAPI
)

// Execute the `link` command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/link.yaml)")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configHome, err := homedir.Expand("~/.config")
		if err != nil {
			er(err)
		}

		viper.AddConfigPath(configHome)
		viper.SetConfigName("link")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot find config file at ~/.config/link.yaml. Please check the README.")
		os.Exit(1)
	} else {
		fmt.Printf("--- Using config found at: %s\n", viper.ConfigFileUsed())
	}

	API = core.NewConnectLinkAPI(
		viper.GetString("link.server_url"),
		viper.GetString("link.username"),
		viper.GetString("link.password"),
	)
}

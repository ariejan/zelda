package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ariejan/zelda/core"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
    "github.com/tidwall/pretty"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	enableColor bool
	disableColor bool

	// RootCmd is where CLI commands come in
	RootCmd = &cobra.Command{
		Use:   "zelda",
		Short: "Zelda is a CLI to Hertek Connect Link",
		Long:  `Zelda is a CLI to Hertek Connect Link - a valid Integrator account with Hertek Connect is required.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
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
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file (default is $HOME/.config/zelda.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&enableColor, "color", "c", true, "enable colour output of JSON (default: true)")
	RootCmd.PersistentFlags().BoolVarP(&disableColor, "no-color", "", false, "disable colour output of JSON (default: false)")
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
		viper.SetConfigName("zelda")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Cannot find config file at ~/.config/zelda.yaml. Please check the README.")
		os.Exit(1)
	} else {
		fmt.Printf("--- Using config found at: %s\n", viper.ConfigFileUsed())
	}

	API = core.NewConnectLinkAPI(
		viper.GetString("zelda.server_url"),
		viper.GetString("zelda.username"),
		viper.GetString("zelda.password"),
	)
}

func prettyPrintJSON(data string) {
	result := pretty.Pretty([]byte(data))

	if (enableColor && !disableColor) {
		result = pretty.Color(result, nil)
	}

	fmt.Println("==> JSON response: ")
	fmt.Println(string(result))
	fmt.Println()
}

func parseIntArg(arg string) int {
	result, err := strconv.Atoi(arg)
	if err != nil {
		er(err)
	}

	return result
}

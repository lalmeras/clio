package main

import (
	"fmt"
	"os"
	"path"
	"restishovh/auth"
	"restishovh/openapi"

	"github.com/spf13/cobra"
)

func Execute() {
	var configFile string
	var account string
	var rootCmd = &cobra.Command{
		Use:   "restish-ovh",
		Short: "Restish external tool for OVH API description and authentication",
		Args:  cobra.NoArgs,
	}
	authCommand := &cobra.Command{
		Use:   "auth",
		Short: "restish external-command for OVH authentication",
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Authenticate(configFile, account)
		},
		Args: cobra.NoArgs,
	}
	accountParameters(authCommand, &configFile, &account)
	rootCmd.AddCommand(authCommand)

	loginCommand := &cobra.Command{
		Use:   "login",
		Short: "Perform login and retrieve a customerKey",
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Login(configFile, account)
		},
		Args: cobra.NoArgs,
	}
	accountParameters(loginCommand, &configFile, &account)
	rootCmd.AddCommand(loginCommand)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "openapi-convert",
		Short: "Openapi 3 conversion",
		RunE:  openapi.Generate,
		Args:  cobra.NoArgs,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func accountParameters(cmd *cobra.Command, configFile *string, account *string) {
	cmd.Flags().StringVarP(
		configFile,
		"config", "c",
		path.Join(os.Getenv("HOME"), ".restish/ovh-auth.json"),
		"Path to configuration file (JSON credentials)")
	cmd.Flags().StringVarP(
		account,
		"account", "a",
		"default",
		"Account alias")
}

func main() {
	Execute()
}

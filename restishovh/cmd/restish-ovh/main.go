package main

import (
	"fmt"
	"os"
	"restishovh/auth"
	"restishovh/openapi"

	"github.com/spf13/cobra"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "restish-ovh",
		Short: "Restish external tool for OVH API description and authentication",
		Args:  cobra.NoArgs,
	}
	rootCmd.AddCommand(&cobra.Command{
		Use:   "auth",
		Short: "restic external-command for OVH authentication",
		RunE:  auth.Authenticate,
		Args:  cobra.NoArgs,
	})
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

func main() {
	Execute()
}

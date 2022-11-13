package cmd

import (
	"fmt"
	"time"

	"github.com/lalmeras/clio/pkg/types"
	"github.com/lalmeras/clio/pkg/util"
	"github.com/spf13/cobra"
)

var authGroup = &cobra.Group{ID: "auth", Title: "Authentication"}

var authCmd = &cobra.Command{
	GroupID: authGroup.ID,
	Use:     "auth",
}

func init() {
	authCmd.AddCommand(currentCredentialsCmd)
	authCmd.AddCommand(detailsCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(timeCmd)
}

var currentCredentialsCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current credentials information.",
	Run: func(cmd *cobra.Command, args []string) {
		credential := types.ApiCredential{}
		util.Check(util.OvhGet("/auth/currentCredential", &credential))
		fmt.Printf("%+v\n", credential)
	},
}

var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Get current credentials details.",
	Run: func(cmd *cobra.Command, args []string) {
		details := types.Details{}
		util.Check(util.OvhGet("/auth/details", &details))
		fmt.Printf("%+v\n", details)
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Expire current credentials.",
	Run: func(cmd *cobra.Command, args []string) {
		util.Check(util.OvhPost("/auth/logout", nil, nil))
		fmt.Printf("Logout done\n")
	},
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the current time of the OVH servers.",
	Run: func(cmd *cobra.Command, args []string) {
		var response int64
		util.Check(util.OvhGet("/auth/time", &response))
		fmt.Printf("%s\n", time.Unix(response, 0).Format(time.RFC3339))
	},
}

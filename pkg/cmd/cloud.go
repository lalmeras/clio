package cmd

import (
	"fmt"

	"github.com/lalmeras/clio/pkg/types"
	"github.com/lalmeras/clio/pkg/util"
	"github.com/spf13/cobra"
)

var cloudGroup = &cobra.Group{ID: "cloud", Title: "Cloud"}

var cloudCmd = &cobra.Command{
	GroupID: cloudGroup.ID,
	Use:     "cloud",
	Short:   "Public Cloud API",
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Public Cloud API",
}

func init() {
	cloudCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectsCmd)
	projectCmd.AddCommand(projectGetCmd)
}

var projectsCmd = &cobra.Command{
	Use:   "list",
	Short: "Project management",
	Run: func(cmd *cobra.Command, args []string) {
		projects := []string{}
		util.Check(util.OvhGet("/cloud/project", &projects))
		fmt.Printf("%+v\n", projects)
	},
}

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Project detail",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]
		url := "/cloud/project/" + projectID
		project := types.Project{}
		util.Check(util.OvhGet(url, &project))
		fmt.Printf("%+v\n", project)
	},
}

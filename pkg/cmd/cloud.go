package cmd

import (
	"fmt"
	"net/url"

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

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Public Cloud API",
}

type ImageFilter struct {
	FlavorType string
	OSType     string
	Region     string
}

var ProjectID string
var imageFilter = ImageFilter{}

func init() {
	cloudCmd.AddCommand(projectsCmd)
	cloudCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectGetCmd)
	projectCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imageListCmd)
	imageCmd.AddCommand(imageGetCmd)

	projectCmd.PersistentFlags().StringVarP(&ProjectID, "project", "p", "", "Project ID")

	imageListCmd.Flags().StringVarP(&imageFilter.FlavorType, "flavor-type", "f", "", "Flavor type filter")
	imageListCmd.Flags().StringVarP(&imageFilter.OSType, "os-type", "o", "", "OS type filter")
	imageListCmd.Flags().StringVarP(&imageFilter.Region, "region", "r", "", "Region filter")

	projectCmd.MarkPersistentFlagRequired("project")
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Project management",
	Run: func(cmd *cobra.Command, args []string) {
		projects := []string{}
		util.Check(util.OvhGet("/cloud/project", &projects))
		for _, projectID := range projects {
			project, err := Project(projectID)
			util.Check(err)
			fmt.Printf("%+v\n", *project)
		}
	},
}

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Project detail",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		project, err := Project(ProjectID)
		util.Check(err)
		fmt.Printf("%+v\n", *project)
	},
}

var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "Project management",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		values := url.Values{}
		if imageFilter.FlavorType != "" {
			values.Set("flavorType", imageFilter.FlavorType)
		}
		if imageFilter.OSType != "" {
			values.Set("osType", imageFilter.OSType)
		}
		if imageFilter.Region != "" {
			values.Set("region", imageFilter.Region)
		}
		murl := "/cloud/project/" + ProjectID + "/image?" + values.Encode()
		images := []types.Image{}
		util.Check(util.OvhGet(murl, &images))
		for _, image := range images {
			fmt.Printf("%+v\n", image)
		}
	},
}

var imageGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Project detail",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		image, err := Image(ProjectID, args[0])
		util.Check(err)
		fmt.Printf("%+v\n", *image)
	},
}

func Project(projectID string) (*types.Project, error) {
	url := "/cloud/project/" + projectID
	project := &types.Project{}
	if err := util.OvhGet(url, &project); err != nil {
		return nil, err
	}
	return project, nil
}

func Image(projectID string, imageID string) (*types.Image, error) {
	murl := "/cloud/project/" + projectID + "/image/" + imageID
	image := &types.Image{}
	if err := util.OvhGet(murl, &image); err != nil {
		return nil, err
	}
	return image, nil
}

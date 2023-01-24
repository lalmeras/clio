package cloud

import (
	"github.com/lalmeras/clio/introspect/cmd"
	"github.com/spf13/cobra"
)

func CloudCommand(rootCmd *cobra.Command) {
	var cloudCmd = &cobra.Command{
		Use:               "cloud",
		Short:             "Cloud API",
		ValidArgsFunction: ListEndpoints,
	}
	rootCmd.AddCommand(cloudCmd)
}

func ListEndpoints(cobraCmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	result := make([]string, 0)
	desc := cmd.DownloadDescription("cloud", "1.0")
	for _, a := range desc.Apis {
		result = append(result, a.Path)
	}
	return result, cobra.ShellCompDirectiveNoFileComp
}

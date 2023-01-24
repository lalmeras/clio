package main

import (
	"fmt"
	"os"

	"github.com/lalmeras/clio/introspect/cloud"
	"github.com/lalmeras/clio/introspect/cmd"
	"github.com/spf13/cobra"
)

func Execute() {
	var clioCmd = &cobra.Command{
		Use:   "clio",
		Short: "API introspection",
	}
	var completionCmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish]",
		Short: "Generate completion script",
		Run:   GenerateCompletion,
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	}
	cmd.ApiCommand(clioCmd)
	cloud.CloudCommand(clioCmd)
	clioCmd.AddCommand(completionCmd)
	if err := clioCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func GenerateCompletion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		cmd.Root().GenBashCompletionV2(os.Stdout, true)
	case "zsh":
		cmd.Root().GenZshCompletion(os.Stdout)
	case "fish":
		cmd.Root().GenFishCompletion(os.Stdout, true)
	case "powershell":
		cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	}
}

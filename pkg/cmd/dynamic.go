package cmd

import (
	"fmt"

	"github.com/lalmeras/clio/introspect/cmd"
	"github.com/lalmeras/clio/introspect/types"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var dynamicCommand = &cobra.Command{
	Use: "dynamic",
}

func register(services []*types.ApiDescription) {
	for _, v := range services {
		if g, ok := Get(v.Operations); ok {
			command := &cobra.Command{
				Use:   v.Path,
				Short: v.Description,
				Long:  g.Description,
				Run:   func(c *cobra.Command, args []string) {},
			}
			dynamicCommand.AddCommand(command)
			for _, p := range *g.Parameters {
				addParameter(command, v, g, p)
			}
		}
	}
}

func addParameter(command *cobra.Command, service *types.ApiDescription, operation *types.ApiOperation, parameter *types.ApiParameter) {
	if parameter.DataType == "string" || parameter.DataType == "password" || parameter.DataType == "ipBlock" || parameter.DataType == "datetime" || parameter.DataType == "uuid" {
		var param string
		command.Flags().StringVar(&param, parameter.Name, "", parameter.Description)
		return
	}
	if parameter.DataType == "long" {
		var param int64
		command.Flags().Int64Var(&param, parameter.Name, 0, parameter.Description)
		return
	}
	if parameter.DataType == "boolean" {
		var param bool
		command.Flags().BoolVar(&param, parameter.Name, false, parameter.Description)
		return
	}
	panic(fmt.Sprintf("Type not known %s\n", parameter.DataType))
}

var SUPPORTED_TYPES []string = []string{
	"string",
	"password",
	"long",
	"datetime",
	"uuid",
	"ipBlock",
	"boolean",
}

func Get(operations *[]*types.ApiOperation) (*types.ApiOperation, bool) {
	for _, v := range *operations {
		if v.HttpMethod == "GET" {
			for _, p := range *v.Parameters {
				if !slices.Contains(SUPPORTED_TYPES, p.DataType) {
					return nil, false
				}
			}
			return v, true
		}
	}
	return nil, false
}

func init() {
	desc := cmd.DownloadDescription("cloud")
	register(*desc.Apis)
}

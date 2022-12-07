package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lalmeras/clio/introspect/cmd"
	"github.com/lalmeras/clio/introspect/types"
	"github.com/lalmeras/clio/introspect/util"
	"github.com/lalmeras/clio/pkg/types_cloud"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var parameters *types_cloud.CmdParameters = &types_cloud.CmdParameters{}

var dynamicCommand = &cobra.Command{
	Use: "dynamic",
}

func register(services []*types.ApiDescription) {
	for _, v := range services {
		if g, ok := Get(v.Operations); ok {
			addCommand(v, g)
		}
	}
}

func Run(c *cobra.Command, args []string, service *types.ApiDescription, operation *types.ApiOperation) {
	url := service.Path
	for _, p := range *operation.Parameters {
		if p.ParamType == "path" {
			url = strings.ReplaceAll(url, fmt.Sprintf("{%s}", p.Name), reflect.ValueOf(parameters).Elem().FieldByName(util.VarName(operation, p)).String())
		}
	}
	var result interface{}
	if url == "/cloud/project" {
		result = []string{}
	} else {
		result = types_cloud.Cloud_Project{}
	}
	fmt.Printf("%s\n", url)
	util.Check(util.OvhGet(url, &result))
	fmt.Printf("%+v\n", result)
}

func addCommand(service *types.ApiDescription, operation *types.ApiOperation) {
	command := &cobra.Command{
		Use:   service.Path,
		Short: service.Description,
		Long:  operation.Description,
		Run:   func(c *cobra.Command, args []string) { Run(c, args, service, operation) },
	}
	dynamicCommand.AddCommand(command)
	for _, p := range *operation.Parameters {
		addParameter(command, service, operation, p)
	}
}

func addParameter(command *cobra.Command, service *types.ApiDescription, operation *types.ApiOperation, parameter *types.ApiParameter) {
	fn := util.VarName(operation, parameter)
	fv := reflect.ValueOf(parameters).Elem().FieldByName(fn)
	if parameter.DataType == "string" || parameter.DataType == "password" || parameter.DataType == "ipBlock" || parameter.DataType == "datetime" || parameter.DataType == "uuid" {
		param := fv.Addr().Interface().(*string)
		command.Flags().StringVar(param, parameter.Name, "", parameter.Description)
		return
	}
	if parameter.DataType == "long" {
		param := fv.Addr().Interface().(*int64)
		command.Flags().Int64Var(param, parameter.Name, 0, parameter.Description)
		return
	}
	if parameter.DataType == "boolean" {
		param := fv.Addr().Interface().(*bool)
		command.Flags().BoolVar(param, parameter.Name, false, parameter.Description)
		return
	}
	panic(fmt.Sprintf("Type not known %s\n", parameter.DataType))
}

func Get(operations *[]*types.ApiOperation) (*types.ApiOperation, bool) {
	for _, v := range *operations {
		if v.HttpMethod == "GET" {
			for _, p := range *v.Parameters {
				if !slices.Contains(types.SUPPORTED_TYPES, p.DataType) {
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
	register(desc.Apis)
}

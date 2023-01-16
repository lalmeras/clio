package cmd

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/lalmeras/clio/introspect/types"
	"github.com/lalmeras/clio/introspect/util"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// root URL for OVH API
const ROOT_URL = "https://api.ovh.com/"

//go:embed types.gotmpl
var STRUCT_TEMPLATE string

func Execute() {
	if err := apiCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// api command group
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API introspection",
}

// api download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download API descriptor",
	Args:  cobra.ExactArgs(1),
	Run:   ApiCommand,
}

// command arguments
var version string

// load command definition
func init() {
	apiCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&version, "version", "v", "1.0", "API version to load")
	downloadCmd.MarkFlagRequired("api")
}

// Download and decode an api description
func DownloadDescription(name string, version string) *types.ApiDescriptionDocument {
	url := ROOT_URL + version + "/" + name + ".json"
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		var result types.ApiDescriptionDocument
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&result)
		return &result
	}
}

// Filter and sort models from api descriptions; complexType.UnitAndValue are removed
// as their are handled with adhoc type definitions.
//
// Types are sorted by name so that code generation is deterministic.
func FilterModels(models map[string]*types.ApiModel) []string {
	sortedModels := make([]string, 0)
	for _, v := range maps.Keys(models) {
		if strings.HasPrefix(v, "complexType.UnitAndValue") {
			continue
		}
		sortedModels = append(sortedModels, v)
	}
	sort.Strings(sortedModels)
	return sortedModels
}

// download and extract api definition, then generate go types and variables for arguments handling.
func ApiCommand(cmd *cobra.Command, args []string) {
	result := DownloadDescription(args[0], "1.0")
	sortedModels := FilterModels(result.Models)

	// extract all parameters from api
	// iterate over endpoints/methods to populate parameter list
	parameters := make(types.ApiOperationsParameters, 0)
	parTypes := make([]string, 0)
	for _, s := range result.Apis {
		for _, o := range s.Operations {
			if "GET" != o.HttpMethod {
				continue
			}
			for _, p := range o.Parameters {
				if CanGenerate(o, p) {
					parameters = append(parameters, &types.ApiOperationParameter{
						Order:     util.VarName(o, p),
						VarName:   util.VarName(o, p),
						VarType:   util.VarType(o, p),
						Operation: o,
						Parameter: p})
				}
				if !slices.Contains(parTypes, p.FullType) {
					parTypes = append(parTypes, p.FullType)
				}
			}
		}
	}
	fmt.Printf("%+v\n", parTypes)
	// sort and deduplicate parameter list
	sort.SliceStable(parameters, parameters.Less)
	parameters = Unique(parameters)

	// open generated target file to store template generation result
	f, err := os.OpenFile(fmt.Sprintf("../pkg/types_%[1]s/%[1]s.go", args[0]), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	template := template.Must(template.New("struct").Parse(STRUCT_TEMPLATE))
	templateModel := types.ApiTemplate{
		Name:         args[0],
		SortedModels: sortedModels,
		Types:        result.Models,
		Imports:      []string{"time", "github.com/lalmeras/clio/pkg/types"},
		Parameters:   parameters,
	}
	if err := template.Execute(buf, &templateModel); err != nil {
		panic(err)
	}
	buf.Flush()
}

// deduplicate parameters from a sorted parameters list
func Unique(parameters []*types.ApiOperationParameter) []*types.ApiOperationParameter {
	result := make([]*types.ApiOperationParameter, 0)
	var last types.ApiOperationParameter
	for _, p := range parameters {
		// only append new values
		if last.Order == "" || last.Order != p.Order {
			result = append(result, p)
		}
		last = *p
	}
	return result
}

// Check if types is an expected type.
func CanGenerate(operation *types.ApiOperation, parameter *types.ApiParameter) bool {
	return slices.Contains(types.SUPPORTED_TYPES, parameter.DataType)
}

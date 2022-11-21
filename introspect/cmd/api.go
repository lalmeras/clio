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
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const ROOT_URL = "https://api.ovh.com/"

//go:embed types.gotmpl
var STRUCT_TEMPLATE string

func Execute() {
	if err := apiCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download API descriptor",
	Args:  cobra.ExactArgs(1),
	Run:   ApiCommand,
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API introspection",
}

var version string

func init() {
	apiCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&version, "version", "v", "1.0", "API version to load")
	downloadCmd.MarkFlagRequired("api")
}

func DownloadDescription(name string) *types.ApiDescriptionDocument {
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

func ApiCommand(cmd *cobra.Command, args []string) {
	result := DownloadDescription(args[0])
	sortedModels := FilterModels(result.Models)
	f, err := os.OpenFile(fmt.Sprintf("pkg/types_%[1]s/%[1]s.go", args[0]), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
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
	}
	if err := template.Execute(buf, &templateModel); err != nil {
		panic(err)
	}
	parTypes := make([]string, 0)
	for _, s := range *result.Apis {
		for _, o := range *s.Operations {
			if "GET" != o.HttpMethod {
				continue
			}
			for _, p := range *o.Parameters {
				if !slices.Contains(parTypes, p.FullType) {
					parTypes = append(parTypes, p.FullType)
				}
			}
		}
	}
	fmt.Printf("%+v\n", parTypes)
	buf.Flush()
}

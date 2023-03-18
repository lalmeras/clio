package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/lalmeras/clio/typovh/openapi"
	"github.com/lalmeras/clio/typovh/types"
	"github.com/spf13/cobra"
)

// root URL for OVH API
const ROOT_URL = "https://api.ovh.com/"

func Execute() {
	var genCmd = &cobra.Command{
		Use:   "clio",
		Short: "API introspection",
		Run:   Generate,
		Args:  cobra.NoArgs,
	}
	if err := genCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func Generate(cmd *cobra.Command, args []string) {
	result := DownloadApiListDescription("1.0")
	for _, a := range result.Apis {
		ApiGenerate(cmd, a.Path, result.BasePath+a.Path+".json")
	}
}

// download and extract api definition, then generate go types and variables for arguments handling.
func ApiGenerate(cmd *cobra.Command, relativePath string, url string) {
	result := DownloadApiDescription(url)
	o := openapi.OpenAPI{
		Openapi: "3.1",
		Info: &openapi.Info{
			Title:   "OVH API",
			Summary: relativePath,
			Version: result.ApiVersion,
		},
		Servers: []openapi.Server{
			{Url: ROOT_URL + "1.0"},
		},
	}
	paths := make(map[string]openapi.PathItem, 0)
	for _, api := range result.Apis {
		path := openapi.PathItem{}
		path.Description = api.Description
		for _, operation := range api.Operations {
			op := openapi.Operation{
				Description: operation.Description,
				Summary:     operation.Description,
				OperationId: api.Path + "-" + strings.ToLower(string(operation.HttpMethod)),
			}
			switch operation.HttpMethod {
			case types.GET:
				path.Get = &op
				break
			case types.POST:
				path.Post = &op
				break
			case types.DELETE:
				path.Delete = &op
				break
			case types.PUT:
				path.Put = &op
				break
			}
			params := make([]openapi.Parameter, 0)
			for _, p := range operation.Parameters {
				in := ""
				switch p.ParamType {
				case types.BODY:
					op.RequestBody = &openapi.RequestBody{
						Description: p.Description,
						Content:     map[string]openapi.MediaType{"application/json": {}},
						Required:    p.Required,
					}
					break
				default:
					in = "path"
					if p.ParamType == types.QUERY {
						in = "query"
					}
					param := openapi.Parameter{
						Name:        p.Name,
						Description: p.Description,
						Required:    p.Required,
						In:          in,
					}
					params = append(params, param)
				}
			}
			op.Parameters = params
			responses := map[string]openapi.Response{
				"200": {
					Content: map[string]openapi.MediaType{"application/json": {}},
				},
			}
			op.Responses = responses
		}
		paths[api.Path] = path
	}
	o.Paths = paths

	outputPath := path.Join("output", relativePath+".json")
	os.MkdirAll(path.Dir(outputPath), 0750)
	file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(o)
	if err != nil {
		panic(err)
	}
}

// Download and decode an api description
func DownloadApiListDescription(version string) *types.ApiListDocument {
	url := ROOT_URL + version + "/"
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		var result types.ApiListDocument
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&result)
		return &result
	}
}

// Download and decode an api description
func DownloadApiDescription(url string) *types.ApiDescriptionDocument {
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else {
		var result types.ApiDescriptionDocument
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&result)
		return &result
	}
}

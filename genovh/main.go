package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/lalmeras/clio/typovh/types"
	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	low "github.com/pb33f/libopenapi/datamodel/low"
	v3low "github.com/pb33f/libopenapi/datamodel/low/v3"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// root URL for OVH API
const ROOT_URL = "https://api.ovh.com/"

func execute() {
	var genCmd = &cobra.Command{
		Use:   "clio",
		Short: "Openapi 3 conversion",
		RunE:  generate,
		Args:  cobra.NoArgs,
	}
	if err := genCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	execute()
}

func generate(cmd *cobra.Command, args []string) error {
	result := DownloadApiListDescription("1.0")
	for _, a := range result.Apis {
		if err := openapiGenerate(cmd, a.Path, result.BasePath+a.Path+".yaml?format=openapi3"); err != nil {
			return err
		}
	}
	return nil
}

// Update openapi3 description from OVH to append generated operationId (from path)
func openapiGenerate(cmd *cobra.Command, relativePath string, url string) error {
	if result, err := DownloadOpenapiDescription(url); err != nil {
		return err
	} else {
		methodNames := []string{
			"get",
			"delete",
			"head",
			"options",
			"patch",
			"post",
			"put",
			"trace",
		}
		v3high, _ := result.BuildV3Model()
		for path, pathItem := range v3high.Model.Paths.PathItems {
			operations := listOperations(pathItem)
			lowOperations := listLowOperations(pathItem)
			for idx, operation := range operations {
				lowOperation := lowOperations[idx]
				methodName := methodNames[idx]
				operationId := buildOperationId(path, methodName)
				updateOperationId(operation, lowOperation, operationId)
			}
		}
		if output, err := result.Serialize(); err != nil {
			return err
		} else {
			outputPath := path.Join("output", relativePath+".yaml")
			os.MkdirAll(path.Dir(outputPath), 0750)
			if file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm); err != nil {
				return err
			} else {
				defer file.Close()
				_, err := file.Write(output)
				return err
			}
		}
	}
}

// Order must be consistent across listMethodNames, listOperations, listLowOperations
func listMethodNames() []string {
	return []string{
		"get",
		"delete",
		"head",
		"options",
		"patch",
		"post",
		"put",
		"trace",
	}
}

// Order must be consistent across listMethodNames, listOperations, listLowOperations
func listOperations(pathItem *v3.PathItem) []*v3.Operation {
	return []*v3.Operation{
		pathItem.Get,
		pathItem.Delete,
		pathItem.Head,
		pathItem.Options,
		pathItem.Patch,
		pathItem.Post,
		pathItem.Put,
		pathItem.Trace,
	}
}

// Order must be consistent across listMethodNames, listOperations, listLowOperations
func listLowOperations(pathItem *v3.PathItem) []low.NodeReference[*v3low.Operation] {
	return []low.NodeReference[*v3low.Operation]{
		pathItem.GoLow().Get,
		pathItem.GoLow().Delete,
		pathItem.GoLow().Head,
		pathItem.GoLow().Options,
		pathItem.GoLow().Patch,
		pathItem.GoLow().Post,
		pathItem.GoLow().Put,
		pathItem.GoLow().Trace,
	}
}

// Build and adapt operationId from path and http method name
func buildOperationId(path string, methodName string) string {
	operationId := path + "-" + methodName
	if strings.HasPrefix(operationId, "/") {
		operationId = operationId[1:]
	}
	operationId = strings.ReplaceAll(operationId, "/", "-")
	// replace parameter name delimiters
	operationId = strings.ReplaceAll(operationId, "{", "")
	operationId = strings.ReplaceAll(operationId, "}", "")
	return operationId
}

// Mutate an libopenapi document to add operationId field
func updateOperationId(operation *v3.Operation, lowOperation low.NodeReference[*v3low.Operation], operationId string) {
	if operation != nil {
		if operation.GoLow().OperationId.IsEmpty() {
			// no operationId; we need to add a yaml.Node
			opContent := lowOperation.ValueNode.Content
			opContent = append(opContent, &yaml.Node{Kind: yaml.ScalarNode, Value: "operationId"})
			opContent = append(opContent, &yaml.Node{Kind: yaml.ScalarNode, Value: operationId})
			lowOperation.ValueNode.Content = opContent
		} else {
			// operationId already here; we need to update it
			operation.GoLow().OperationId.Mutate(operationId)
		}
	}
}

// Download and decode API listing
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

// Download and decode an openapi description
func DownloadOpenapiDescription(url string) (libopenapi.Document, error) {
	if resp, err := http.Get(url); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		data := new(bytes.Buffer)
		data.ReadFrom(resp.Body)
		if d, err := libopenapi.NewDocument(data.Bytes()); err != nil {
			return nil, err
		} else {
			return d, nil
		}
	}
}

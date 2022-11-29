package util

import (
	"fmt"
	"os"

	"github.com/lalmeras/clio/introspect/types"
	"github.com/ovh/go-ovh/ovh"
)

func Check(err error) {
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		os.Exit(1)
	}
}

func OvhGet(url string, resType interface{}) error {
	client, err := ovh.NewDefaultClient()
	if err != nil {
		return err
	}
	return client.Get(url, resType)
}

func OvhPost(url string, reqBody interface{}, resType interface{}) error {
	client, err := ovh.NewDefaultClient()
	if err != nil {
		return err
	}
	return client.Post(url, reqBody, resType)
}

func VarName(operation *types.ApiOperation, parameter *types.ApiParameter) string {
	return fmt.Sprintf("%s_%s_%s", operation.HttpMethod, parameter.Name, parameter.DataType)
}

func VarType(operation *types.ApiOperation, parameter *types.ApiParameter) string {
	if parameter.DataType == "string" || parameter.DataType == "password" || parameter.DataType == "ipBlock" || parameter.DataType == "datetime" || parameter.DataType == "uuid" {
		return "string"
	}
	if parameter.DataType == "long" {
		return "int64"
	}
	if parameter.DataType == "boolean" {
		return "bool"
	}
	panic(fmt.Sprintf("Type not known %s\n", parameter.DataType))
}

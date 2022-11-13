package util

import (
	"fmt"
	"os"

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

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/lalmeras/clio/pkg/types"
	"github.com/lalmeras/clio/pkg/util"
	"github.com/spf13/cobra"
)

const ROOT_URL = "https://api.ovh.com/"

var apiGroup = &cobra.Group{ID: "api", Title: "API Introspection"}

var apiCmd = &cobra.Command{
	GroupID: cloudGroup.ID,
	Use:     "api",
	Short:   "API introspection",
}

var version string
var extractTypes []string

func init() {
	apiCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&version, "version", "v", "1.0", "API version to load")
	downloadCmd.Flags().StringArrayVarP(&extractTypes, "type", "t", make([]string, 0), "Type to extract")
	downloadCmd.MarkFlagRequired("api")
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download API descriptor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var result types.ApiDescriptionDocument
		url := ROOT_URL + version + "/" + args[0] + ".json"
		resp, err := http.Get(url)
		util.Check(err)
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&result)
		for _, t := range extractTypes {
			if val, ok := result.Models[t]; ok {
				var rendered bytes.Buffer
				template := template.Must(template.New("struct").Parse(STRUCT_TEMPLATE))
				if err := template.Execute(&rendered, val); err == nil {
					fmt.Printf(rendered.String())
				} else {
					fmt.Printf("%q\n", err)
				}
				fmt.Printf("Type %s found:\n%+v\n", t, val)
			}
		}
	},
}

const STRUCT_TEMPLATE = `
type {{ .ID }} struct {
{{- $fieldPadValue := .NamePadding }}
{{- $typePadValue := .TypePadding }}
{{- $model := . }}
{{- range $key, $value := .Properties }}
{{ $model.PadGoName $key $value $fieldPadValue }} {{ $model.PadGoType $key $value $typePadValue }} {{ $model.JsonTag $key }}
{{- end }}
}
`

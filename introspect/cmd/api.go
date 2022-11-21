package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"text/template"

	"github.com/lalmeras/clio/introspect/types"
	"github.com/lalmeras/clio/introspect/util"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

func Execute() {
	if err := apiCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const ROOT_URL = "https://api.ovh.com/"

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "API introspection",
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
		if len(extractTypes) == 1 && extractTypes[0] == "*" {
			extractTypes = maps.Keys(result.Models)
			sort.Strings(extractTypes)
		}
		ok := true
		f, err := os.OpenFile(fmt.Sprintf("pkg/types_%[1]s/%[1]s.go", args[0]), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		buf := bufio.NewWriter(f)
		fmt.Fprintf(buf, "package types_%s\n\n", args[0])
		fmt.Fprint(buf, `import (
	"time"
	"github.com/lalmeras/clio/pkg/types"
)`+"\n\n")
		for _, t := range extractTypes {
			if val, ok := result.Models[t]; ok {
				if val.ID == "UnitAndValue" {
					continue
				}
				var rendered bytes.Buffer
				template := template.Must(template.New("struct").Parse(STRUCT_TEMPLATE))
				if err := template.Execute(&rendered, &types.ApiTemplate{CurrentType: val, Types: result.Models}); err == nil {
					fmt.Fprint(buf, rendered.String())
				} else {
					fmt.Printf("%s.%s\n", val.Namespace, val.ID)
					fmt.Printf("%q\n", err)
					ok = false
				}
			}
		}
		if !ok {
			fmt.Printf("One or more error during generation.\n")
		}
		buf.Flush()
	},
}

const STRUCT_TEMPLATE = `
{{- $typeName := .GoTypeName .CurrentType }}
{{- $template := . }}
{{- $model := .CurrentType }}
{{- if gt (len .CurrentType.Properties) 0 }}
type {{ $typeName }} struct {
{{- $fieldPadValue := .NamePadding .CurrentType }}
{{- $typePadValue := .TypePadding .Types .CurrentType }}
{{- range $key, $value := .CurrentType.Properties }}
	{{ $template.PadGoName $key $value $fieldPadValue }} {{ $template.PadGoType $template.Types $key $value $typePadValue }} {{ $template.JsonTag $key }}
	{{- end }}
}
{{- else if .CurrentType.EnumType | eq "string" }}
type {{ $typeName }} string
const (
{{- range $key, $value := .CurrentType.Enum }}
	{{ $template.GoEnumSymbol $model $value }} {{ if $key | eq 0 }}{{ $typeName }}{{ end }} = "{{ $value }}"
{{- end }}
)
{{- else if .CurrentType.EnumType | eq "long" }}
type {{ $typeName }} int64
const (
{{- range $key, $value := .CurrentType.Enum }}
	{{ $template.GoEnumSymbol $model $value }} {{ if $key | eq 0 }}{{ $typeName }}{{ end }} = {{ $value }}
{{- end }}
)
{{- else if .CurrentType.EnumType | eq "double" }}
type {{ $typeName }} float64
const (
{{- range $key, $value := .CurrentType.Enum }}
	{{ $template.GoEnumSymbol $model $value }} {{ if $key | eq 0 }}{{ $typeName }}{{ end }} = {{ $value }}
{{- end }}
)
{{- end }}
`

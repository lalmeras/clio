package types

import (
	"fmt"
	"strings"
)

type ApiDescriptionDocument struct {
	ApiVersion   string                 `json:"apiversion"`
	Apis         []*ApiDescription      `json:"apis"`
	ResourcePath string                 `json:"resourcePath"`
	BasePath     string                 `json:"basePath"`
	Models       map[string]*ApiModel   `json:"models"`
	X            map[string]interface{} `json:"-"`
}

type ApiDescription struct {
	Path        string                 `json:"path"`
	Operations  []interface{}          `json:"operations"`
	Description string                 `json:"description"`
	X           map[string]interface{} `json:"-"`
}

type ApiModel struct {
	ID          string                 `json:"id"`
	Namespace   string                 `json:"namespace"`
	Description string                 `json:"description"`
	Enum        []string               `json:"enum"`
	EnumType    string                 `json:"enumType"`
	Properties  map[string]*ApiField   `json:"properties"`
	X           map[string]interface{} `json:"-"`
}

type ApiField struct {
	Type        string                 `json:"type"`
	FullType    string                 `json:"fullType"`
	CanBeNull   bool                   `json:"canBeNull"`
	ReadOnly    bool                   `json:"readOnly"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	X           map[string]interface{} `json:"-"`
}

type ApiOperation struct {
	ApiStatus        ApiStatus              `json:"apiStatus"`
	HttpMethod       string                 `json:"httpMethod"`
	Parameters       []*ApiParameter        `json:"parameters"`
	ResponseType     string                 `json:"responseType"`
	NoAuthentication bool                   `json:"noAuthentication"`
	Description      string                 `json:"description"`
	Scopes           []string               `json:"scopes"`
	X                map[string]interface{} `json:"-"`
}

type ApiParameter struct {
	Name        string                 `json:"product"`
	DataType    string                 `json:"dataType"`
	ParamType   string                 `json:"paramType"`
	FullType    string                 `json:"fullType"`
	Required    bool                   `json:"required"`
	Description string                 `json:"description"`
	X           map[string]interface{} `json:"-"`
}

type ApiStatus struct {
	Description string                 `json:"description"`
	Value       string                 `json:"value"`
	X           map[string]interface{} `json:"-"`
}

func (self *ApiModel) JsonTag(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

func (self *ApiModel) GoName(name string, field *ApiField) string {
	if goType, ok := GO_TYPES[field.Type]; ok {
		return goType
	}
	if strings.HasPrefix(field.Type, "cloud.") {
		return field.Type
	}
	panic(fmt.Sprintf("Type %s cannot be mapped", field.Type))
}

func (self *ApiModel) GoType(name string, field *ApiField) string {
	if field.Type == "datetime" {
		return "time.Time"
	}
	return field.Type
}

func (self *ApiModel) PadGoType(name string, field *ApiField, padding int) string {
	goType := self.GoType(name, field)
	if length := len(goType); length < padding {
		return goType + strings.Repeat(" ", padding-length)
	}
	return goType
}

func (self *ApiModel) PadGoName(name string, field *ApiField, padding int) string {
	goName := self.GoName(name, field)
	if length := len(goName); length < padding {
		return goName + strings.Repeat(" ", padding-length)
	}
	return goName
}

func (self *ApiModel) NamePadding() int {
	max := 0
	for name, val := range self.Properties {
		if nameLength := len(self.GoName(name, val)); nameLength > max {
			max = nameLength
		}
	}
	return max
}

func (self *ApiModel) TypePadding() int {
	max := 0
	for name, field := range self.Properties {
		if typeLength := len(self.GoType(name, field)); typeLength > max {
			max = typeLength
		}
	}
	return max
}

var GO_TYPES = map[string]string{
	"string":   "string",
	"boolean":  "bool",
	"datetime": "time.Time",
	"long":     "int64",
	"double":   "float64",
}

package types

import (
	"fmt"
	"strconv"
	"strings"
)

type ApiDescriptionDocument struct {
	ApiVersion   string                  `json:"apiversion"`
	Apis         *[]*ApiDescription      `json:"apis"`
	ResourcePath string                  `json:"resourcePath"`
	BasePath     string                  `json:"basePath"`
	Models       map[string]*ApiModel    `json:"models"`
	X            map[string]*interface{} `json:"-"`
}

type ApiDescription struct {
	Path        string                  `json:"path"`
	Operations  *[]*ApiOperation        `json:"operations"`
	Description string                  `json:"description"`
	X           map[string]*interface{} `json:"-"`
}

type ApiModel struct {
	ID          string                  `json:"id"`
	Namespace   string                  `json:"namespace"`
	Description string                  `json:"description"`
	Enum        *[]string               `json:"enum"`
	EnumType    string                  `json:"enumType"`
	Properties  map[string]*ApiField    `json:"properties"`
	X           map[string]*interface{} `json:"-"`
}

type ApiField struct {
	Type        string                  `json:"type"`
	FullType    string                  `json:"fullType"`
	CanBeNull   bool                    `json:"canBeNull"`
	ReadOnly    bool                    `json:"readOnly"`
	Description string                  `json:"description"`
	Required    bool                    `json:"required"`
	X           map[string]*interface{} `json:"-"`
}

type ApiOperation struct {
	ApiStatus        *ApiStatus             `json:"apiStatus"`
	HttpMethod       string                 `json:"httpMethod"`
	Parameters       *[]*ApiParameter       `json:"parameters"`
	ResponseType     string                 `json:"responseType"`
	NoAuthentication bool                   `json:"noAuthentication"`
	Description      string                 `json:"description"`
	Scopes           *[]string              `json:"scopes"`
	X                map[string]interface{} `json:"-"`
}

type ApiParameter struct {
	Name        string                  `json:"name"`
	DataType    string                  `json:"dataType"`
	ParamType   string                  `json:"paramType"`
	FullType    string                  `json:"fullType"`
	Required    bool                    `json:"required"`
	Description string                  `json:"description"`
	X           map[string]*interface{} `json:"-"`
}

type ApiStatus struct {
	Description string                  `json:"description"`
	Value       string                  `json:"value"`
	X           map[string]*interface{} `json:"-"`
}

type ApiTemplate struct {
	CurrentType *ApiModel
	Types       map[string]*ApiModel
}

func (self *ApiTemplate) JsonTag(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

func (self *ApiTemplate) GoTypeName(model *ApiModel) string {
	result := fmt.Sprintf("%s_%s", strings.ReplaceAll(model.Namespace, ".", ""), model.ID)
	return strings.ToUpper(result[0:1]) + result[1:]
}

func (self *ApiTemplate) GoEnumSymbol(model *ApiModel, enumValue string) string {
	result := strings.ReplaceAll(enumValue, "-", "")
	result = strings.ReplaceAll(result, ".", "")
	result = strings.ReplaceAll(result, ":", "")
	result = strings.ReplaceAll(result, "/", "")
	// prefix numeric value with _
	if _, err := strconv.Atoi(result); err == nil {
		result = "_" + result
	}
	return self.GoTypeName(model) + "_" + result
}

func (self *ApiTemplate) GoName(name string, field *ApiField) string {
	name = strings.ReplaceAll(name, "-", "")
	return strings.ToUpper(name[0:1]) + name[1:]
}

func (self *ApiTemplate) GoType(models map[string]*ApiModel, name string, field *ApiField) string {
	containedType := field.Type
	var tt string
	var list bool
	if strings.HasSuffix(field.Type, "[]") {
		list = true
		containedType = field.Type[0 : len(field.Type)-2]
	}
	if t, ok := models[containedType]; ok {
		tt = self.GoTypeName(t)
	}
	if t, ok := GO_TYPES[containedType]; ok {
		tt = t
	}
	if tt != "" {
		if list {
			return "[]" + tt
		} else {
			return tt
		}
	}
	panic(fmt.Sprintf("Type %s cannot be mapped", field.Type))
}

func (self *ApiTemplate) PadGoType(models map[string]*ApiModel, name string, field *ApiField, padding int) string {
	goType := self.GoType(models, name, field)
	if length := len(goType); length < padding {
		return goType + strings.Repeat(" ", padding-length)
	}
	return goType
}

func (self *ApiTemplate) PadGoName(name string, field *ApiField, padding int) string {
	goName := self.GoName(name, field)
	if length := len(goName); length < padding {
		return goName + strings.Repeat(" ", padding-length)
	}
	return goName
}

func (self *ApiTemplate) NamePadding(model *ApiModel) int {
	max := 0
	for name, val := range model.Properties {
		if nameLength := len(self.GoName(name, val)); nameLength > max {
			max = nameLength
		}
	}
	return max
}

func (self *ApiTemplate) TypePadding(models map[string]*ApiModel, model *ApiModel) int {
	max := 0
	for name, field := range model.Properties {
		if typeLength := len(self.GoType(models, name, field)); typeLength > max {
			max = typeLength
		}
	}
	return max
}

var GO_TYPES = map[string]string{
	"string":                           "string",
	"boolean":                          "bool",
	"datetime":                         "time.Time",
	"date":                             "time.Time",
	"time":                             "time.Time",
	"long":                             "int64",
	"double":                           "float64",
	"uuid":                             "string",
	"ip":                               "string",
	"ipBlock":                          "string",
	"ipv4":                             "string",
	"ipv4Block":                        "string",
	"ipv6":                             "string",
	"ipv6Block":                        "string",
	"map[string]string":                "map[string]string",
	"map[string]long":                  "map[string]int64",
	"text":                             "string",
	"password":                         "string",
	"duration":                         "time.Duration",
	"complexType.UnitAndValue<long>":   "types.UnitAndValueInt64",
	"complexType.UnitAndValue<double>": "types.UnitAndValueFloat64",
}

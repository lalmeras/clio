package types

import (
	"fmt"
	"strconv"
	"strings"
)

// Template generation: main model
// - Name: API name
// - Types: type name to description map
// - SortedModels: ordered types names
// - Imports: needed imports
// - Parameters: TODO
type ApiTemplate struct {
	Name         string
	Types        map[string]*ApiModel
	SortedModels []string
	Imports      []string
	Parameters   ApiOperationsParameters
}

// Template generation: parameter model
// - Order: parameter name used for ordering
// - VarName: parameter global variable name
// - VarType: parameter go type
// - Operation: TODO
// - Parameter: TODO
type ApiOperationParameter struct {
	Order     string
	VarName   string
	VarType   string
	Operation *ApiOperation
	Parameter *ApiParameter
}

type ApiOperationsParameters []*ApiOperationParameter

// Template generation: parameter ordering
func (self ApiOperationsParameters) Less(i, j int) bool { return self[i].Order < self[j].Order }

// Generate json tag to control type serialization/deserialization
// - name: json field name
func (self *ApiTemplate) JsonTag(name string) string {
	return fmt.Sprintf("`json:\"%s\"`", name)
}

// Generate a valid go type name or an ApiModel (replace forbidden chars, capitalize first letter)
// - model: type we want to generate a go type name for
func (self *ApiTemplate) GoTypeName(model *ApiModel) string {
	result := fmt.Sprintf("%s_%s", strings.ReplaceAll(model.Namespace, ".", ""), model.ID)
	return strings.ToUpper(result[0:1]) + result[1:]
}

// Generate valid go enumeration names (replace forbidden chars, concatenate with type name)
// - model: enumerated type
// - enumValue: json enumeration name we want to generate a go type name for
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

// Generate a valid go name (replace forbidden chars, capitalize)
func (self *ApiTemplate) GoName(name string, field *ApiField) string {
	name = strings.ReplaceAll(name, "-", "")
	return strings.ToUpper(name[0:1]) + name[1:]
}

// Generate a valid go type name, handling go native type mapping and list case
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

// Generate type part for a struct field declaration, applying padding for alignment
func (self *ApiTemplate) PadGoType(models map[string]*ApiModel, name string, field *ApiField, padding int) string {
	goType := self.GoType(models, name, field)
	if length := len(goType); length < padding {
		return goType + strings.Repeat(" ", padding-length)
	}
	return goType
}

// Generate name part for a struct field declaration, applying padding for alignment
func (self *ApiTemplate) PadGoName(name string, field *ApiField, padding int) string {
	goName := self.GoName(name, field)
	if length := len(goName); length < padding {
		return goName + strings.Repeat(" ", padding-length)
	}
	return goName
}

// Compute padding for field names
func (self *ApiTemplate) NamePadding(model *ApiModel) int {
	max := 0
	for name, val := range model.Properties {
		if nameLength := len(self.GoName(name, val)); nameLength > max {
			max = nameLength
		}
	}
	return max
}

// Compute padding for type names
func (self *ApiTemplate) TypePadding(models map[string]*ApiModel, model *ApiModel) int {
	max := 0
	for name, field := range model.Properties {
		if typeLength := len(self.GoType(models, name, field)); typeLength > max {
			max = typeLength
		}
	}
	return max
}

// Basic types
var SUPPORTED_TYPES []string = []string{
	"string",
	"password",
	"long",
	"datetime",
	"uuid",
	"ipBlock",
	"boolean",
	"phoneNumber",
	"macAddress",
	"internationalPhoneNumber",
	"coreTypes.AccountId:string",
}

// Native types mapping
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
	"phoneNumber":                      "string",
	"macAddress":                       "string",
	"internationalPhoneNumber":         "string",
	"coreTypes.AccountId:string":       "string",
}

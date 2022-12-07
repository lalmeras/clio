// This is type definition for json deserializtion of OVH API description
// (like https://api.ovh.com/1.0/cloud.json)
package types

// This is the root type of OVH API description
type ApiDescriptionDocument struct {
	ApiVersion   string                 `json:"apiversion"`
	Apis         []*ApiDescription      `json:"apis"`
	ResourcePath string                 `json:"resourcePath"`
	BasePath     string                 `json:"basePath"`
	Models       map[string]*ApiModel   `json:"models"`
	X            map[string]interface{} `json:"-"`
}

// Endpoint description
// Multiple methods are described as separate operation.
type ApiDescription struct {
	Path        string                 `json:"path"`
	Operations  *[]*ApiOperation       `json:"operations"`
	Description string                 `json:"description"`
	X           map[string]interface{} `json:"-"`
}

// Type description
type ApiModel struct {
	ID          string                 `json:"id"`
	Namespace   string                 `json:"namespace"`
	Description string                 `json:"description"`
	Enum        *[]string              `json:"enum"`
	EnumType    string                 `json:"enumType"`
	Properties  map[string]*ApiField   `json:"properties"`
	X           map[string]interface{} `json:"-"`
}

// Type's fields description
type ApiField struct {
	Type        string                 `json:"type"`
	FullType    string                 `json:"fullType"`
	CanBeNull   bool                   `json:"canBeNull"`
	ReadOnly    bool                   `json:"readOnly"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	X           map[string]interface{} `json:"-"`
}

// Endpoint variation (by method)
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

// Endpoint input parameter description
type ApiParameter struct {
	Name        string                 `json:"name"`
	DataType    string                 `json:"dataType"`
	ParamType   string                 `json:"paramType"`
	FullType    string                 `json:"fullType"`
	Required    bool                   `json:"required"`
	Description string                 `json:"description"`
	X           map[string]interface{} `json:"-"`
}

// Api status description (beta, stable, ...)
type ApiStatus struct {
	Description string                 `json:"description"`
	Value       string                 `json:"value"`
	X           map[string]interface{} `json:"-"`
}

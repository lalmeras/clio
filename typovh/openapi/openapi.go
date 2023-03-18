package openapi

type OpenAPI struct {
	Openapi           string              `json:"openapi"`
	Info              *Info               `json:"info"`
	JsonSchemaDialect string              `json:"jsonSchemaDialect"`
	Servers           []Server            `json:"servers"`
	Paths             map[string]PathItem `json:"paths"`
}

type Info struct {
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	Description    string `json:"description"`
	TermsOfService string `json:"termsOfService"`
	Version        string `json:"version"`
}

type Server struct {
	Url         string                    `json:"url"`
	Description string                    `json:"description"`
	Variables   map[string]ServerVariable `json:"variables"`
}

type PathItem struct {
	Ref         string      `json:"$ref,omitempty"`
	Summary     string      `json:"summary,omitempty"`
	Description string      `json:"description,omitempty"`
	Get         *Operation  `json:"get,omitempty"`
	Put         *Operation  `json:"put,omitempty"`
	Post        *Operation  `json:"post,omitempty"`
	Delete      *Operation  `json:"delete,omitempty"`
	Options     *Operation  `json:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty"`
	Patch       *Operation  `json:"patch,omitempty"`
	Trace       *Operation  `json:"trace,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Name            string `json:"name"`
	In              string `json:"in"`
	Description     string `json:"description"`
	Required        bool   `json:"required"`
	Deprecated      bool   `json:"deprecated"`
	AllowEmptyValue bool   `json:"allowEmptyValue"`
}

type Operation struct {
	Summary     string              `json:"summary,omitempty"`
	Description string              `json:"description,omitempty"`
	OperationId string              `json:"operationId,omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty"`
}

type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

type RequestBody struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"mediaType"`
	Required    bool                 `json:"required"`
}

type MediaType struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
}

type Response struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

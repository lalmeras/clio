package types

type ApisTemplate struct {
	Apis     []ApiDescriptor `json:"apis"`
	BasePath string          `json:"basePath"`
}

type ApiDescriptor struct {
	Path        string    `json:"path"`
	Schema      string    `json:"schema"`
	Format      ApiFormat `json:"format"`
	Description string    `json:"description"`
}

type ApiFormat string

const (
	JSON ApiFormat = "json"
	YAML           = "yaml"
)

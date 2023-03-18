package types

type Endpoints []Endpoint
type Endpoint struct {
	Path   string
	Method ApiHttpMethod
}

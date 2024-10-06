package model

type Swagger struct {
	OpenAPI    string              `json:"openapi"`
	Info       Info                `json:"info"`
	Servers    []Server            `json:"servers"`
	Paths      map[string]PathItem `json:"paths"`
	Components Components          `json:"components"`
	Security   []Security          `json:"security,omitempty"`
	Tags       []Tag               `json:"tags,omitempty"`
}

type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type Server struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type PathItem struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
}

type Operation struct {
	Tags        []string            `json:"tags"`
	Summary     *string             `json:"summary"`
	Description *string             `json:"description"`
	Parameters  []Parameter         `json:"parameters,omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses"`
}

type Parameter struct {
	Name        *string `json:"name,omitempty"`
	In          *string `json:"in,omitempty"`
	Description *string `json:"description,omitempty"`
	Required    *bool   `json:"required,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
	Ref         *string `json:"$ref,omitempty"`
	Minimum     *int    `json:"minimum,omitempty"`
	Default     *any    `json:"default,omitempty"`
}

type RequestBody struct {
	Ref      *string               `json:"$ref,omitempty"`
	Required *bool                 `json:"required,omitempty"`
	Content  *map[string]MediaType `json:"content,omitempty"`
}

type Response struct {
	Ref         *string              `json:"$ref,omitempty"`
	Description *string              `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Headers     map[string]Header    `json:"headers,omitempty"`
}

type Header struct {
	Ref         *string `json:"$ref,omitempty"`
	Description *string `json:"description,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

type MediaType struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type        *string           `json:"type,omitempty"`
	Ref         *string           `json:"$ref,omitempty"`
	Items       *Schema           `json:"items,omitempty"`
	Enum        []any             `json:"enum,omitempty"`
	AllOf       []Schema          `json:"allOf,omitempty"`
	AnyOf       []Schema          `json:"anyOf,omitempty"`
	OneOf       []Schema          `json:"oneOf,omitempty"`
	Example     *any              `json:"example,omitempty"`
	Format      *string           `json:"format,omitempty"`
	Default     *any              `json:"default,omitempty"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Description *string           `json:"description,omitempty"`
	Required    []string          `json:"required,omitempty"`
	ReadOnly    *bool             `json:"readOnly,omitempty"`
	WriteOnly   *bool             `json:"writeOnly,omitempty"`
	Minimum     *int              `json:"minimum,omitempty"`
	MaxLength   *int              `json:"maxLength,omitempty"`
}

type Components struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
}

type SecurityScheme struct {
	Type  string `json:"type"`
	Flows any    `json:"flows,omitempty"`
}

type Security map[string]any

type Tag struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

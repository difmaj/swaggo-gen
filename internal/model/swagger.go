package model

type Swagger struct {
	OpenAPI    string                     `json:"openapi"`
	Info       SwaggerInfo                `json:"info"`
	Servers    []SwaggerServer            `json:"servers"`
	Paths      map[string]SwaggerPathItem `json:"paths"`
	Components SwaggerComponents          `json:"components"`
	Security   []SwaggerSecurity          `json:"security,omitempty"`
	Tags       []SwaggerTag               `json:"tags,omitempty"`
}

type SwaggerInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type SwaggerServer struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type SwaggerPathItem struct {
	Get    *SwaggerOperation `json:"get,omitempty"`
	Post   *SwaggerOperation `json:"post,omitempty"`
	Put    *SwaggerOperation `json:"put,omitempty"`
	Delete *SwaggerOperation `json:"delete,omitempty"`
	Patch  *SwaggerOperation `json:"patch,omitempty"`
}

type SwaggerOperation struct {
	Tags        []string                   `json:"tags"`
	Summary     *string                    `json:"summary"`
	Description *string                    `json:"description"`
	Parameters  []SwaggerParameter         `json:"parameters,omitempty"`
	RequestBody *SwaggerRequestBody        `json:"requestBody,omitempty"`
	Responses   map[string]SwaggerResponse `json:"responses"`
}

type SwaggerParameter struct {
	Name        *string        `json:"name,omitempty"`
	In          *string        `json:"in,omitempty"`
	Description *string        `json:"description,omitempty"`
	Required    *bool          `json:"required,omitempty"`
	Schema      *SwaggerSchema `json:"schema,omitempty"`
	Ref         *string        `json:"$ref,omitempty"`
	Minimum     *int           `json:"minimum,omitempty"`
	Default     *any           `json:"default,omitempty"`
}

type SwaggerRequestBody struct {
	Ref      *string                      `json:"$ref,omitempty"`
	Required *bool                        `json:"required,omitempty"`
	Content  *map[string]SwaggerMediaType `json:"content,omitempty"`
}

type SwaggerResponse struct {
	Ref         *string                     `json:"$ref,omitempty"`
	Description *string                     `json:"description,omitempty"`
	Content     map[string]SwaggerMediaType `json:"content,omitempty"`
	Headers     map[string]SwaggerHeader    `json:"headers,omitempty"`
}

type SwaggerHeader struct {
	Ref         *string        `json:"$ref,omitempty"`
	Description *string        `json:"description,omitempty"`
	Schema      *SwaggerSchema `json:"schema,omitempty"`
}

type SwaggerMediaType struct {
	Schema SwaggerSchema `json:"schema"`
}

type SwaggerSchema struct {
	Type        *string                  `json:"type,omitempty"`
	Ref         *string                  `json:"$ref,omitempty"`
	Items       *SwaggerSchema           `json:"items,omitempty"`
	Enum        []any                    `json:"enum,omitempty"`
	AllOf       []SwaggerSchema          `json:"allOf,omitempty"`
	AnyOf       []SwaggerSchema          `json:"anyOf,omitempty"`
	OneOf       []SwaggerSchema          `json:"oneOf,omitempty"`
	Example     *any                     `json:"example,omitempty"`
	Format      *string                  `json:"format,omitempty"`
	Default     *any                     `json:"default,omitempty"`
	Properties  map[string]SwaggerSchema `json:"properties,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Required    []string                 `json:"required,omitempty"`
	ReadOnly    *bool                    `json:"readOnly,omitempty"`
	WriteOnly   *bool                    `json:"writeOnly,omitempty"`
	Minimum     *int                     `json:"minimum,omitempty"`
	MaxLength   *int                     `json:"maxLength,omitempty"`
}

type SwaggerComponents struct {
	Schemas         map[string]SwaggerSchema         `json:"schemas,omitempty"`
	Responses       map[string]SwaggerResponse       `json:"responses,omitempty"`
	Parameters      map[string]SwaggerParameter      `json:"parameters,omitempty"`
	SecuritySchemes map[string]SwaggerSecurityScheme `json:"securitySchemes,omitempty"`
	Headers         map[string]SwaggerHeader         `json:"headers,omitempty"`
}

type SwaggerSecurityScheme struct {
	Type  string `json:"type"`
	Flows any    `json:"flows,omitempty"`
}

type SwaggerSecurity map[string]any

type SwaggerTag struct {
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

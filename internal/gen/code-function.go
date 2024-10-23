package gen

import (
	"fmt"
	"io"
)

// Function is a method in a package
type Function struct {
	Name        string // method name
	Description string // optional

	HTTPMethod string // one of HTTP methods defined in RFC 7231 section 4.3.
	Path       string // URL path template

	Request, Response *Parameter // method parameters
}

// Print writes the method signature to w.
func (m *Function) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "// %s %s\n", m.Name, m.Description)
	_, _ = fmt.Fprintf(w, "func %s(", m.Name)
	_, _ = fmt.Fprintf(w, "%s %T", m.Request.Name, m.Request.Type)

	_, _ = fmt.Fprint(w, ") (")
	_, _ = fmt.Fprintf(w, "%s %T", m.Response.Name, m.Response.Type)
	_, _ = fmt.Fprint(w, ")")

	_, _ = fmt.Fprintf(w, " {\n")
	_, _ = fmt.Fprintf(w, "\treq, err := c.controller.NewRequest(ctx, http.Method%s, %q, request)\n", m.HTTPMethod, m.Path)
	_, _ = fmt.Fprintf(w, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w, "\t\treturn nil, err\n")
	_, _ = fmt.Fprintf(w, "\t}\n")
	_, _ = fmt.Fprintf(w, "\n")
	_, _ = fmt.Fprintf(w, "\tresponse := new(%T)\n", m.Response.Type)
	_, _ = fmt.Fprintf(w, "\t_, err = c.controller.Do(req, response)\n")
	_, _ = fmt.Fprintf(w, "\tif err != nil {\n")
	_, _ = fmt.Fprintf(w, "\t\treturn nil, err\n")
	_, _ = fmt.Fprintf(w, "\t}\n")
	_, _ = fmt.Fprintf(w, "\treturn response, nil\n")
	_, _ = fmt.Fprintf(w, "}\n")
}

// Parameter is an argument or return parameter of a method.
type Parameter struct {
	Name string
	Type any
}

// Print writes the parameter to w.
func (p *Parameter) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "%s %T", p.Name, p.Type)
}

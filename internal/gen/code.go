package gen

import (
	"fmt"
	"io"
)

// File is a Go package
type File struct {
	PkgName   string
	Imports   []string
	Consts    []*Const
	Functions []*Function
}

// Print writes the package name and imports to w.
func (pkg *File) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "package %s\n", pkg.PkgName)

	_, _ = fmt.Fprintf(w, "import (\n")
	for _, imports := range pkg.Imports {
		_, _ = fmt.Fprintf(w, "\t\"%s\"\n", imports)
	}
	_, _ = fmt.Fprintf(w, ")\n\n")

	_, _ = fmt.Fprintf(w, "// Endpoints\n")
	_, _ = fmt.Fprintf(w, "const (\n")
	for _, path := range pkg.Consts {
		_, _ = fmt.Fprintf(w, "\t// %s %s\n", path.Name, path.Description)
		_, _ = fmt.Fprintf(w, "\t%s = \"/%s\"\n", path.Name, path.Value)
	}
	_, _ = fmt.Fprintf(w, ")\n\n")

	for _, function := range pkg.Functions {
		function.Print(w)
	}
}

// _, _ = fmt.Fprintf(w, "// Client is a client\n")
// _, _ = fmt.Fprintf(w, "type Client struct {\n")
// _, _ = fmt.Fprintf(w, "\tcontroller IController\n")
// _, _ = fmt.Fprintf(w, "}\n\n")

// _, _ = fmt.Fprintf(w, "// NewClient returns a new client\n")
// _, _ = fmt.Fprintf(w, "func NewClient(controller *Controller) *Client {\n")
// _, _ = fmt.Fprintf(w, "\treturn &Client{controller: controller}\n")
// _, _ = fmt.Fprintf(w, "}\n\n")

// _, _ = fmt.Fprintf(w, "// Controller is a controller\n")
// _, _ = fmt.Fprintf(w, "type Controller interface {\n")
// _, _ = fmt.Fprintf(w, "\tDo(req *http.Request, resp *http.Response) (*http.Response, error)\n")
// _, _ = fmt.Fprintf(w, "\tNewRequest(ctx context.Context, method, path string, body any) (*Request, error)\n")
// _, _ = fmt.Fprintf(w, "}\n\n")

// Const is a URL path
type Const struct {
	Name        string
	Description string
	Value       string
}

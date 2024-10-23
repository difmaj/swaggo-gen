package gen

import (
	"fmt"
	"io"
)

// Model is a struct in a package
type Models struct {
	Models []*Model
}

// Print writes the struct to w.
func (m *Models) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "package model\n\n")
	for _, model := range m.Models {
		model.Print(w)
	}
}

// Model is a struct in a package
type Model struct {
	Name        string
	Description string
	Fields      []*Field
}

// Print writes the struct to w.
func (m *Model) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "\n")
	_, _ = fmt.Fprintf(w, "type %s struct {\n", m.Name)
	for _, field := range m.Fields {
		field.Print(w)
	}
	_, _ = fmt.Fprintf(w, "}\n")
}

// Field is a field in a struct
type Field struct {
	Name        string
	Description string
	Type        string
	Tags        []*Tag
}

// Print writes the field to w.
func (f *Field) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "\t%s %s", f.Name, f.Type)
	if len(f.Tags) > 0 {
		_, _ = fmt.Fprintf(w, " `")
		for _, tag := range f.Tags {
			tag.Print(w)
		}
		_, _ = fmt.Fprintf(w, "`")
	}
	_, _ = fmt.Fprintf(w, "\n")
}

// Tag is a tag in a field
type Tag struct {
	Name  string
	Value string
}

// Print writes the tag to w.
func (t *Tag) Print(w io.Writer) {
	_, _ = fmt.Fprintf(w, "%s:\"%s\" ", t.Name, t.Value)
}

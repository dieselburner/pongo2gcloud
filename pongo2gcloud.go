package pongo2gcloud

import (
	"fmt"
	"io"
)

// Loader implements pongo2's TemplateLoader interface for templates stored on Google Cloud Storage
type Loader struct {
	ProjectID string
}

// Abs calculates the path to a given template. Whenever a path must be resolved
// due to an import from another template, the base equals the parent template's path.
func (l *Loader) Abs(base, name string) string {
	fmt.Print("Abs: base=", base, " / name=", name)
	return name
}

// Get returns an io.Reader where the template's content can be read from.
func (l *Loader) Get(path string) (io.Reader, error) {
	fmt.Print("Get: path=", path)
	return nil, fmt.Errorf("not implemented")
}

package pongo2gcloud

import (
	"context"
	"io"
	"log"
	"path"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/storage"
)

// Loader implements pongo2's TemplateLoader interface for templates stored on Google Cloud Storage
type Loader struct {
	ProjectID string
}

// Abs calculates the path to a given template. Whenever a path must be resolved
// due to an import from another template, the base equals the parent template's path.
func (l *Loader) Abs(base, name string) string {
	dir, _ := path.Split(base)
	return path.Join(dir, name)
}

// Get returns an io.Reader where the template's content can be read from.
func (l *Loader) Get(name string) (io.Reader, error) {
	var err error
	var project_id string

	// If ProjectID is not set, try to fetch it from Google Cloud Engine
	if len(l.ProjectID) > 0 {
		project_id = l.ProjectID
	} else {
		if !metadata.OnGCE() {
			log.Fatalf("ProjectID needs to be set when running outside of Google Cloud Engine")
		}
		project_id, err = metadata.ProjectID()
		if err != nil {
			log.Fatalf("Could not fetch Google Compute Engine metadata: %v", err)
		} else if project_id == "" {
			log.Fatalf("Could not obtain Google Gloud project ID")
		}
	}

	// Create Google Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Could not create Google Cloud Storage Client")
	}

	// Get template file from project bucket
	reader, err := client.Bucket(project_id + ".appspot.com").Object(name).NewReader(ctx)
	if err != nil {
		log.Fatalf("Could not get template file '%s' from bucket: %v", name, err)
	}

	// Done
	return reader, nil
}

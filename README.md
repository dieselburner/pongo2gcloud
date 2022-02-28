# pongo2gcloud

[![Go Reference](https://pkg.go.dev/badge/github.com/dieselburner/pongo2gcloud.svg)](https://pkg.go.dev/github.com/dieselburner/pongo2gcloud)

[pongo2](https://github.com/flosch/pongo2) provides Django-syntax like templating language for Go applications. One of the great things about pongo2 is that it supports [template sets](https://pkg.go.dev/github.com/flosch/pongo2#TemplateSet) and [custom template loaders](https://pkg.go.dev/github.com/flosch/pongo2#TemplateLoader).

This package adds support for loading templates from [Google Cloud Storage](https://cloud.google.com/storage).

# Usage

Import package:

```go
import "github.com/dieselburner/pongo2gcloud"
```

[Gin](https://github.com/gin-gonic/gin) example, through [modified pongo2gin middleware](https://github.com/dieselburner/pongo2gin):

Set `ProjectID` to your Google Cloud Project ID. If not set, `pongo2gcloud` will autodetect it when running on Google Cloud Engine, otherwise will panic.

```go
template_set := pongo2.NewSet("gcloud", &pongo2gcloud.Loader{
	ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
})

router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
	ContentType: "text/html; charset=utf-8",
	TemplateDir: "templates",
	TemplateSet: template_set,
})
```

package docs

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed docs.go.tmpl
var docsData embed.FS

func DocsData(data any) ([]byte, error) {

	templ, err := template.ParseFS(docsData, "docs.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = templ.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

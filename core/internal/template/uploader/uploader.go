package uploader

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed uploader.go.tmpl
var uploaderData embed.FS

func UplaoderData(data any) ([]byte, error) {

	temp, err := template.ParseFS(uploaderData, "uploader.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, nil)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

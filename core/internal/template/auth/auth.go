package auth

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed auth.go.tmpl
var authData embed.FS

func AuthData(data any) ([]byte, error) {
	temp, err := template.ParseFS(authData, "auth.go.tmpl")
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

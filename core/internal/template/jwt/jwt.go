package jwt

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed jwt.go.tmpl
var jwtData embed.FS

func JwtData(data any) ([]byte, error) {
	temp, err := template.ParseFS(jwtData, "jwt.go.tmpl")
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

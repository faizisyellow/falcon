package env

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed .env.tmpl
var envData embed.FS

func EnvData(data any) ([]byte, error) {
	temp, err := template.ParseFS(envData, ".env.tmpl")
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

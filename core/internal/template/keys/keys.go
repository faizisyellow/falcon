package keys

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed keys.go.tmpl
var keysData embed.FS

func KeysData(data any) ([]byte, error) {

	temp, err := template.ParseFS(keysData, "keys.go.tmpl")
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

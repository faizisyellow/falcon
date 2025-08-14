package air

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed .air.toml.tmpl
var airData embed.FS

func AirData(data any) ([]byte, error) {

	temp, err := template.ParseFS(airData, ".air.toml.tmpl")
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

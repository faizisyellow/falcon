package loggertemplate

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed logger.go.tmpl
var loggerData embed.FS

func LoggerData(any) ([]byte, error) {

	temp, err := template.ParseFS(loggerData, "logger.go.tmpl")
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

package gitignore

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed .gitignore.tmpl
var gitignoreData embed.FS

func GitignoreData(data any) ([]byte, error) {

	temp, err := template.ParseFS(gitignoreData, ".gitignore.tmpl")
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

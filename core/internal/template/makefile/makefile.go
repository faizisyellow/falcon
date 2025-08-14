package makefile

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed makefile.tmpl
var makefileData embed.FS

func MakefileData(any) ([]byte, error) {

	temp, err := template.ParseFS(makefileData, "makefile.tmpl")
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

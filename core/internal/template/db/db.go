package db

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

type DBOption struct {
	DB string
}

//go:embed mysql/mysql.go.tmpl
var dbData embed.FS

func DBData(data any) ([]byte, error) {

	dbOptions := make(map[string]string)
	dbOptions["mysql"] = "mysql/mysql.go.tmpl"

	opt, ok := data.(DBOption)
	if !ok {
		return nil, fmt.Errorf("error unknwon db option")
	}

	// parse file from option
	temp, err := template.ParseFS(dbData, dbOptions[opt.DB])
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

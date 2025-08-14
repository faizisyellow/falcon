package service

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/faizisyellow/falcon/internal/utils"
)

//go:embed service.go.tmpl
var serviceData embed.FS

func ServiceData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(serviceData, "service.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, datas)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

//go:embed users.go.tmpl
var usersData embed.FS

func UsersData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(usersData, "users.go.tmpl")
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)

	err = temp.Execute(b, datas)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

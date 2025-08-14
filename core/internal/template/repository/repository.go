package repository

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/faizisyellow/falcon/internal/utils"
)

//go:embed invitation.go.tmpl
var invitationData embed.FS

func InvitationData(data any) ([]byte, error) {

	temp, err := template.ParseFS(invitationData, "invitation.go.tmpl")
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

//go:embed repository.go.tmpl
var repositoryData embed.FS

func RepositoryData(data any) ([]byte, error) {

	temp, err := template.ParseFS(repositoryData, "repository.go.tmpl")
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

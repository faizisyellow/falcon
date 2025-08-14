package utilstemplate

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed token.go.tmpl
var tokenData embed.FS

func TokenData(data any) ([]byte, error) {
	temp, err := template.ParseFS(tokenData, "token.go.tmpl")
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

//go:embed password.go.tmpl
var passwordData embed.FS

func PasswordData(data any) ([]byte, error) {
	temp, err := template.ParseFS(passwordData, "password.go.tmpl")
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

//go:embed contentContext.go.tmpl
var contentContextData embed.FS

func ContentContextData(data any) ([]byte, error) {
	temp, err := template.ParseFS(contentContextData, "contentContext.go.tmpl")
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

//go:embed pointer.go.tmpl
var pointerData embed.FS

func PointerData(data any) ([]byte, error) {
	temp, err := template.ParseFS(pointerData, "pointer.go.tmpl")
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

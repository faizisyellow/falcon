package api

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/faizisyellow/falcon/internal/utils"
)

//go:embed "api.go.tmpl"
var apiData embed.FS

func ApiData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(apiData, "api.go.tmpl")
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

//go:embed "auth.go.tmpl"
var auth embed.FS

func AuthData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(auth, "auth.go.tmpl")
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

//go:embed "json.go.tmpl"
var jsonData embed.FS

func JsonData(data any) ([]byte, error) {

	temp, err := template.ParseFS(jsonData, "json.go.tmpl")
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

//go:embed "main.go.tmpl"
var mainData embed.FS

func MainData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(mainData, "main.go.tmpl")
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

//go:embed "middlewares.go.tmpl"
var middlewaresData embed.FS

func MiddlewaresData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(middlewaresData, "middlewares.go.tmpl")
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

//go:embed "middlewares_test.go.tmpl"
var middlewaresTestData embed.FS

func MiddlewaresTestData(data any) ([]byte, error) {

	temp, err := template.ParseFS(middlewaresTestData, "middlewares_test.go.tmpl")
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

//go:embed "response.go.tmpl"
var responseData embed.FS

func ResponseData(data any) ([]byte, error) {

	datas := struct {
		Module string
	}{
		Module: utils.GetModuleName(),
	}

	temp, err := template.ParseFS(responseData, "response.go.tmpl")
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

//go:embed "users.go.tmpl"
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

type RouterOpt struct {
	Router string
}

//go:embed "mux.go.tmpl"
var muxData embed.FS

//go:embed "router/chi.go.tmpl"
var chi embed.FS

func MuxData(data any) ([]byte, error) {

	temp, err := template.ParseFS(muxData, "mux.go.tmpl")
	if err != nil {
		return nil, err
	}

	opt, ok := data.(RouterOpt)
	if !ok {
		return nil, fmt.Errorf("error unknown router option type")
	}

	switch opt.Router {
	case "chi":
		_, err = temp.ParseFS(chi, "router/chi.go.tmpl")
		if err != nil {
			return nil, err
		}
	}

	b := new(bytes.Buffer)
	err = temp.Execute(b, opt)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

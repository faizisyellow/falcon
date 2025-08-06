package api

import (
	_ "embed"
)

//go:embed api.go.tmpl
var Api []byte

//go:embed auth.go.tmpl
var Auth []byte

//go:embed errors.go.tmpl
var Errors []byte

//go:embed json.go.tmpl
var Json []byte

//go:embed main.go.tmpl
var Main []byte

//go:embed middlewares.go.tmpl
var Middlewares []byte

//go:embed users.go.tmpl
var Users []byte

package service

import _ "embed"

//go:embed service.go.tmpl
var Service []byte

//go:embed users.go.tmpl
var Users []byte

package utils

import _ "embed"

//go:embed token.go.tmpl
var Token []byte

//go:embed password.go.tmpl
var Password []byte

//go:embed contentContext.go.tmpl
var ContentContext []byte

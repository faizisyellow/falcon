package auth

import _ "embed"

//go:embed auth.go.tmpl
var Auth []byte

package jwt

import _ "embed"

//go:embed jwt.go.tmpl
var Jwt []byte

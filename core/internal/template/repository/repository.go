package repository

import _ "embed"

//go:embed invitation.go.tmpl
var Invitation []byte

//go:embed repository.go.tmpl
var Repository []byte

//go:embed users.go.tmpl
var Users []byte

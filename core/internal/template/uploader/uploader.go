package uploader

import _ "embed"

//go:embed uploader.go.tmpl
var Uploader []byte

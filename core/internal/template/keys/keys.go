package keys

import _ "embed"

//go:embed keys.go.tmpl
var Keys []byte

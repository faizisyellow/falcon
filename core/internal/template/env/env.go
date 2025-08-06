package env

import _ "embed"

//go:embed .env.tmpl
var Env []byte

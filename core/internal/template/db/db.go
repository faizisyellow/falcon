package db

import _ "embed"

//go:embed mysql/mysql.go.tmpl
var Mysql []byte

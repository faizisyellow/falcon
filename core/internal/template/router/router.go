package router

import _ "embed"

//go:embed mux.go.tmpl
var Mux []byte

//go:embed chi/chi.go.tmpl
var ChiRouter []byte

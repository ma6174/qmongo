package url

import (
	"net/url"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/url",

	"queryEscape":   url.QueryEscape,
	"queryUnescape": url.QueryUnescape,

	"URL":             qlang.StructOf((*url.URL)(nil)),
	"parse":           url.Parse,
	"parseRequestURI": url.ParseRequestURI,
	"user":            url.User,
	"userPassword":    url.UserPassword,
}

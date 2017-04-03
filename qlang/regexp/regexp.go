package regexp

import (
	"regexp"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "regexp",

	"match":       regexp.Match,
	"matchReader": regexp.MatchReader,
	"matchString": regexp.MatchString,
	"quoteMeta":   regexp.QuoteMeta,

	"compile":          regexp.Compile,
	"compilePOSIX":     regexp.CompilePOSIX,
	"mustCompile":      regexp.MustCompile,
	"mustCompilePOSIX": regexp.MustCompilePOSIX,
}

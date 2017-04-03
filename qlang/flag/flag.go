package flag

import (
	"flag"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "flag",

	"ContinueOnError": flag.ContinueOnError,
	"ExitOnError":     flag.ExitOnError,
	"PanicOnError":    flag.PanicOnError,

	"CommandLine": flag.CommandLine,
	"ErrHelp":     flag.ErrHelp,
	"Usage":       flag.Usage,

	"arg":           flag.Arg,
	"args":          flag.Args,
	"bool":          flag.Bool,
	"boolVar":       flag.BoolVar,
	"duration":      flag.Duration,
	"durationVar":   flag.DurationVar,
	"float64":       flag.Float64,
	"float64Var":    flag.Float64Var,
	"int":           flag.Int,
	"int64":         flag.Int64,
	"int64Var":      flag.Int64Var,
	"intVar":        flag.IntVar,
	"NArg":          flag.NArg,
	"NFlag":         flag.NFlag,
	"parse":         flag.Parse,
	"parsed":        flag.Parsed,
	"printDefaults": flag.PrintDefaults,
	"set":           flag.Set,
	"string":        flag.String,
	"stringVar":     flag.StringVar,
	"uint":          flag.Uint,
	"uint64":        flag.Uint64,
	"uint64Var":     flag.Uint64Var,
	"uintVar":       flag.UintVar,
	"var":           flag.Var,
	"visit":         flag.Visit,
	"visitAll":      flag.VisitAll,

	"Flag":    qlang.StructOf((*flag.Flag)(nil)),
	"lookup":  flag.Lookup,
	"flagSet": flag.NewFlagSet,
}

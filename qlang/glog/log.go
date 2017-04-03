package glog

import (
	"log"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "glog",

	"Ldate":         log.Ldate,
	"Llongfile":     log.Llongfile,
	"Lmicroseconds": log.Lmicroseconds,
	"Lshortfile":    log.Lshortfile,
	"LstdFlags":     log.LstdFlags,
	"Ltime":         log.Ltime,

	"fatal":     log.Fatal,
	"fatalf":    log.Fatalf,
	"fatalln":   log.Fatalln,
	"flags":     log.Flags,
	"panic":     log.Panic,
	"panicf":    log.Panicf,
	"panicln":   log.Panicln,
	"prefix":    log.Prefix,
	"print":     log.Print,
	"printf":    log.Printf,
	"println":   log.Println,
	"setFlags":  log.SetFlags,
	"setOutput": log.SetOutput,
	"setPrefix": log.SetPrefix,

	"new": log.New,
}

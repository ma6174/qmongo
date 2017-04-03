package main

import (
	"github.com/ma6174/qmongo/qlang/flag"
	"github.com/ma6174/qmongo/qlang/fmt"
	"github.com/ma6174/qmongo/qlang/glog"
	"github.com/ma6174/qmongo/qlang/gopkg.in/mgo.v2"
	"github.com/ma6174/qmongo/qlang/gopkg.in/mgo.v2/bson"
	"github.com/ma6174/qmongo/qlang/net"
	"github.com/ma6174/qmongo/qlang/net/url"
	"github.com/ma6174/qmongo/qlang/regexp"
	"github.com/ma6174/qmongo/qlang/time"

	"qlang.io/cmd/qshell"
	"qlang.io/spec"
)

// -----------------------------------------------------------------------------

func main() {

	spec.Import("bson", bson.Exports)
	spec.Import("flag", flag.Exports)
	spec.Import("fmt", fmt.Exports)
	spec.Import("glog", glog.Exports)
	spec.Import("mgo", mgo.Exports)
	spec.Import("net", net.Exports)
	spec.Import("regexp", regexp.Exports)
	spec.Import("time", time.Exports)
	spec.Import("url", url.Exports)
	qshell.Main(false)
}

// -----------------------------------------------------------------------------

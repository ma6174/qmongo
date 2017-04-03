// +build go1.8

package net

import (
	"net"

	qlang "qlang.io/spec"
)

func init() {
	Exports["Resolver"] = qlang.StructOf((*net.Resolver)(nil))
	Exports["DefaultResolver"] = net.DefaultResolver
}

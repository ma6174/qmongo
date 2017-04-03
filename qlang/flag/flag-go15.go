// +build go1.5

package flag

import "flag"

func init() {
	Exports["unquoteUsage"] = flag.UnquoteUsage
}

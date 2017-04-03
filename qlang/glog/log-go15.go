// +build go1.5

package glog

import "log"

func init() {
	Exports["LUTC"] = log.LUTC
	Exports["output"] = log.Output
}

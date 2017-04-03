// +build go1.8

package url

import "net/url"

func init() {
	Exports["pathEscape"] = url.PathEscape
	Exports["pathUnescape"] = url.PathUnescape
}

package fmt

import (
	"fmt"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "fmt",

	"errorf":   fmt.Errorf,
	"fprint":   fmt.Fprint,
	"fprintf":  fmt.Fprintf,
	"fprintln": fmt.Fprintln,
	"fscan":    fmt.Fscan,
	"fscanf":   fmt.Fscanf,
	"fscanln":  fmt.Fscanln,
	"print":    fmt.Print,
	"printf":   fmt.Printf,
	"println":  fmt.Println,
	"scan":     fmt.Scan,
	"scanf":    fmt.Scanf,
	"scanln":   fmt.Scanln,
	"sprint":   fmt.Sprint,
	"sprintf":  fmt.Sprintf,
	"sprintln": fmt.Sprintln,
	"sscan":    fmt.Sscan,
	"sscanf":   fmt.Sscanf,
	"sscanln":  fmt.Sscanln,
}

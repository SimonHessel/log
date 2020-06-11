package log

import (
	"fmt"

	"github.com/fatih/color"
)

func prefix(prefix string, c color.Attribute) func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		color.New(c).Printf(fmt.Sprintf("[%s] %s\n", prefix, format), a...)
	}
}

var Info = prefix("Info", color.FgGreen)
var Warning = prefix("Warning", color.FgYellow)
var Error = prefix("Error", color.FgRed)

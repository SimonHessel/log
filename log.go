package log

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func prefix(prefix string, color func(interface{}) Value) func(format string, a ...interface{}) {
	return func(format string, a ...interface{}) {
		fmt.Println(color(fmt.Sprintf(fmt.Sprintf("[%s] %s", prefix, format), a...)))
	}
}

var Info = prefix("Info", Green)
var Warning = prefix("Warning", Yellow)
var Error = prefix("Error", Red)

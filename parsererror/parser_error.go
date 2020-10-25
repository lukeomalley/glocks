package parsererror

import (
	"fmt"
	"os"
)

// HadError stores whether or not the parser encountered an error
var HadError = false

// LogMessage reports in stderr an error encountered during parsing
func LogMessage(line int, message string) {
	report(line, "", message)
	HadError = true
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %v] Error at '%s'\n", line, where)
}

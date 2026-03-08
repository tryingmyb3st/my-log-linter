package analyzer

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "myLogLinter",
	Doc:  "my implemetation of linter for logs",
	Run:  run,
}

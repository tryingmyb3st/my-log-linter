package mylinter

import (
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"github.com/golangci/golangci-lint/v2/pkg/golinters/mylinter/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinter("myLinter", "...", []*analysis.Analyzer{
			analyzer.Analyzer,
		}, nil).
		WithLoadMode(goanalysis.LoadModeSyntax)

}

package directives

import (
	"reflect"

	"github.com/Abirdcfly/go-tools/analysis/lint"
	"golang.org/x/tools/go/analysis"
)

func directives(pass *analysis.Pass) (interface{}, error) {
	return lint.ParseDirectives(pass.Files, pass.Fset), nil
}

var Analyzer = &analysis.Analyzer{
	Name:             "directives",
	Doc:              "extracts linter directives",
	Run:              directives,
	RunDespiteErrors: true,
	ResultType:       reflect.TypeOf([]lint.Directive{}),
}

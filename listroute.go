package listroute

import (
	"bytes"
	"go/ast"
	"go/format"
	"log"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "listroute displays all routes in the project files."

var Analyzer = &analysis.Analyzer{
	Name: "listroute",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func Contains[T comparable](target T, list []T) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}

var matcherFunctionsRaw string // -listroute.matcherFunctions="..."

func init() {
	Analyzer.Flags.StringVar(&matcherFunctionsRaw, "matcherFunctions", "GET,POST,PUT,DELETE,PATCH,Static,File,Group", "Matcher functions. You can specify multiple by separating them with ','.")
	log.Println(Analyzer.Flags.Parsed())
}

func run(pass *analysis.Pass) (any, error) {
	matcherFunctions := strings.Split(matcherFunctionsRaw, ",")

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	var err error
	buf := &bytes.Buffer{}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			log.Println("node is not *ast.CallExpr")
		}
		switch fun := callExpr.Fun.(type) {
		case *ast.Ident:
			if !Contains(fun.Name, matcherFunctions) {
				return
			}
		case *ast.SelectorExpr:
			if !Contains(fun.Sel.Name, matcherFunctions) {
				return
			}
		default:
			// do nothing
		}

		err = format.Node(buf, pass.Fset, callExpr)
		if err != nil {
			log.Println("format.Node failed")
			return
		}

		pass.Reportf(callExpr.Pos(), "%v", buf.String())
		buf.Reset()
	})

	return nil, nil
}

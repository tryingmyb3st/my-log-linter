package analyzer

import (
	"go/ast"
	"strings"
	"unicode"

	"github.com/lovelydeng/gomoji"
	"golang.org/x/tools/go/analysis"
)

var sensitiveWords = []string{
	"password:", "password=", "key:", "key=", "token:", "token=",
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {

			if _, ok := n.(*ast.CallExpr); !ok {
				return true
			}

			node, _ := n.(*ast.CallExpr)

			if !strings.Contains(node.Fun.(*ast.SelectorExpr).X.(*ast.Ident).Name, "log") {
				return true
			}

			for _, arg := range node.Args {

				if _, ok := arg.(*ast.BasicLit); !ok {
					leftOperand := arg.(*ast.BinaryExpr).X
					value := leftOperand.(*ast.BasicLit).Value

					if isHasSensitiveData(value) {
						pass.Reportf(leftOperand.(*ast.BasicLit).ValuePos, "logs shouldn't contain credentials")
					}
					continue
				}

				value := arg.(*ast.BasicLit).Value

				if isFirstCapital([]rune(value)[1]) {
					pass.Reportf(arg.(*ast.BasicLit).ValuePos, "first letter should be in lowercase")
				}

				if !isOnlyLatinLetters(value) {
					pass.Reportf(arg.(*ast.BasicLit).ValuePos, "all logs should be in english")
				}

				if isHasSpecialSymbols(value) {
					pass.Reportf(arg.(*ast.BasicLit).ValuePos, "logs shouldn't contain special symbols")
				}
			}
			return true
		})
	}
	return nil, nil
}

func isFirstCapital(str rune) bool {
	return unicode.IsUpper(str)
}

func isOnlyLatinLetters(str string) bool {
	for i, letter := range str {
		if i == 0 || len(str) == i+1 || gomoji.ContainsEmoji(string(letter)) {
			continue
		}

		if !(unicode.In(letter, unicode.Latin, unicode.Number, unicode.Space, unicode.Punct)) {
			return false
		}
	}
	return true
}

func isHasSpecialSymbols(str string) bool {
	if gomoji.ContainsEmoji(str) {
		return true
	}

	for i, letter := range str {
		if i == 0 || len(str) == i+1 {
			continue
		}

		if unicode.In(letter, unicode.Punct) {
			return true
		}
	}
	return false
}

func isHasSensitiveData(str string) bool {
	for _, word := range sensitiveWords {
		if strings.Contains(str, word) {
			return true
		}
	}
	return false
}

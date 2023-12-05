package main

import (
	"fmt"
	"github.com/kushsharma/tinyxpr/pkg/eval"
	"github.com/kushsharma/tinyxpr/pkg/lexer"
	"github.com/kushsharma/tinyxpr/pkg/parser"
)

func main() {
	input := "1 + 3 * (4 + CEIL(9.5))"
	tokens := lexer.Tokenize(input)
	fmt.Println(tokens)
	ast := parser.NewParser(tokens).Parse()
	fmt.Println(ast)
	evaluator := eval.NewEvaluator()

	result := evaluator.Evaluate(ast)
	fmt.Printf("Result: %f\n", result)
}

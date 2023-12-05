package eval

import (
	"fmt"
	"github.com/kushsharma/tinyxpr/pkg/lexer"
	"math"
	"strconv"

	"github.com/kushsharma/tinyxpr/pkg/parser" // Change this to your actual package name
)

// Interpreter represents the interpreter.
type Interpreter struct{}

// NewEvaluator creates a new interpreter.
func NewEvaluator() *Interpreter {
	return &Interpreter{}
}

// Evaluate starts the evaluation process.
func (i *Interpreter) Evaluate(ast *parser.ASTNode) float64 {
	if ast == nil {
		return 0
	}

	switch ast.Type {
	case lexer.INTEGER:
		value, err := strconv.ParseFloat(ast.Value, 64)
		if err != nil {
			fmt.Printf("Error converting value to float64: %s\n", err)
			return 0
		}
		return value
	case lexer.PLUS:
		return i.Evaluate(ast.LeftChild) + i.Evaluate(ast.RightChild)
	case lexer.MULTIPLY:
		return i.Evaluate(ast.LeftChild) * i.Evaluate(ast.RightChild)
	case lexer.CEIL:
		return math.Ceil(i.Evaluate(ast.LeftChild))
	default:
		fmt.Printf("Unsupported operation: %s\n", ast.Value)
		return 0
	}
}

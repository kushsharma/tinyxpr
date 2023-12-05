package parser

import (
	"fmt"
	"github.com/kushsharma/tinyxpr/pkg/lexer" // Change this to your actual package name
)

// ASTNode represents a node in the Abstract Syntax Tree (AST).
type ASTNode struct {
	Type       lexer.TokenType
	Value      string
	LeftChild  *ASTNode
	RightChild *ASTNode
}

// String function of ASTNode
func (a *ASTNode) String() string {
	if a == nil {
		return ""
	}
	if a.LeftChild == nil && a.RightChild == nil {
		return fmt.Sprintf("ASTNode(%s, '%s')", a.Type, a.Value)
	}
	return fmt.Sprintf("ASTNode(%s, '%s', %s, %s)", a.Type, a.Value, a.LeftChild, a.RightChild)
}

// Parser represents the parser.
type Parser struct {
	tokens  []lexer.Token
	current int
}

// NewParser creates a new parser with the given tokens.
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

// Parse starts the parsing process.
func (p *Parser) Parse() *ASTNode {
	return p.parseExpression()
}

// parseExpression parses an expression.
func (p *Parser) parseExpression() *ASTNode {
	return p.parseBinaryOperation(p.parseTerm(), 0)
}

// parseTerm parses a term.
func (p *Parser) parseTerm() *ASTNode {
	return p.parseBinaryOperation(p.parseFactor(), 1)
}

// parseBinaryOperation parses binary operations with given precedence.
func (p *Parser) parseBinaryOperation(left *ASTNode, precedence int) *ASTNode {
	for p.isBinaryOperator(precedence) {
		operator := p.consume()
		right := p.parseFactor()
		left = &ASTNode{
			Type:       operator.Type,
			Value:      operator.Value,
			LeftChild:  left,
			RightChild: right,
		}
	}
	return left
}

// parseFactor parses a factor.
func (p *Parser) parseFactor() *ASTNode {
	token := p.consume()

	switch token.Type {
	case lexer.INTEGER:
		return &ASTNode{Type: lexer.INTEGER, Value: token.Value}
	case lexer.CEIL:
		return &ASTNode{Type: lexer.CEIL, Value: token.Value, LeftChild: p.parseExpression()}
	case lexer.LPAREN:
		expression := p.parseExpression()
		p.consumeExpect(lexer.RPAREN, "Expected closing parenthesis after expression.")
		return expression
	default:
		p.error("Unexpected token: %v", token)
		return nil
	}
}

// consume reads the current token and moves to the next one.
func (p *Parser) consume() lexer.Token {
	if p.current < len(p.tokens) {
		token := p.tokens[p.current]
		p.current++
		return token
	}
	return lexer.Token{Type: lexer.EOF, Value: ""}
}

// consumeExpect consumes the current token if it matches the expected type,
// otherwise, it raises an error.
func (p *Parser) consumeExpect(expectedType lexer.TokenType, errorMessage string) {
	token := p.consume()
	if token.Type != expectedType {
		p.error(errorMessage)
	}
}

// isBinaryOperator checks if the current token is a binary operator with the given precedence.
func (p *Parser) isBinaryOperator(precedence int) bool {
	// we don't care about operator precedence for now

	if p.current < len(p.tokens) {
		operator := p.tokens[p.current]
		if operator.Type == lexer.PLUS || operator.Type == lexer.MULTIPLY {
			return true
		}
	}
	return false
}

// error raises a parsing error with a formatted message.
func (p *Parser) error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	panic(fmt.Sprintf("Parsing error at position %d: %s", p.current, message))
}

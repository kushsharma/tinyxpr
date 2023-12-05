package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// TokenType represents the type of token.
type TokenType int

const (
	ILLEGAL = iota
	EOF
	INTEGER
	PLUS
	MULTIPLY
	LPAREN
	RPAREN
	CEIL
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case INTEGER:
		return "INTEGER"
	case PLUS:
		return "PLUS"
	case MULTIPLY:
		return "MULTIPLY"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case CEIL:
		return "CEIL"
	}
	return ""
}

// Token represents a lexical token.
type Token struct {
	Type  TokenType
	Value string
}

// String function of token
func (t Token) String() string {
	return fmt.Sprintf("Token(%s, '%s')", t.Type, t.Value)
}

// Lexer represents the lexer.
type Lexer struct {
	input  string
	pos    int
	currCh byte
}

// NewLexer creates a new lexer for the given input string.
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar() // initialize currCh
	return lexer
}

// readChar reads the next character from the input string.
func (l *Lexer) readChar() {
	if l.pos < len(l.input) {
		l.currCh = l.input[l.pos]
	} else {
		l.currCh = 0 // ASCII code for null character
	}
	l.pos++
}

// peekChar returns the next character without advancing the position.
func (l *Lexer) peekChar() byte {
	if l.pos < len(l.input) {
		return l.input[l.pos]
	}
	return 0
}

// skipWhitespace skips any whitespace characters.
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.currCh)) {
		l.readChar()
	}
}

// NextToken returns the next token from the input string.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.currCh {
	case 0:
		tok = Token{Type: EOF, Value: ""}
	case '+':
		tok = Token{Type: PLUS, Value: string(l.currCh)}
	case '*':
		tok = Token{Type: MULTIPLY, Value: string(l.currCh)}
	case '(':
		tok = Token{Type: LPAREN, Value: string(l.currCh)}
	case ')':
		tok = Token{Type: RPAREN, Value: string(l.currCh)}
	default:
		if unicode.IsDigit(rune(l.currCh)) {
			tok.Type = INTEGER
			var value strings.Builder
			for unicode.IsDigit(rune(l.currCh)) || l.currCh == '.' {
				value.WriteRune(rune(l.currCh))
				l.readChar()
			}
			tok.Value = value.String()
			return tok
		} else {
			// Handle identifiers (CEIL in this case)
			var value strings.Builder
			for unicode.IsLetter(rune(l.currCh)) {
				value.WriteRune(rune(l.currCh))
				l.readChar()
			}
			identifier := strings.ToUpper(value.String())
			switch identifier {
			case "CEIL":
				tok = Token{Type: CEIL, Value: identifier}
			default:
				tok = Token{Type: ILLEGAL, Value: identifier}
			}
			return tok
		}
	}

	l.readChar()
	return tok
}

// Tokenize takes an input string and returns a slice of tokens.
func Tokenize(input string) []Token {
	var tokens []Token
	lexer := NewLexer(input)

	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)

		if token.Type == EOF {
			break
		}
	}

	return tokens
}

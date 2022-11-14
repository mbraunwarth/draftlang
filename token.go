package main

import "fmt"

type Token struct {
	typ  TokenType
	val  string
	line int
	col  int
}

func (t Token) String() string {
	return fmt.Sprintf("Token{%s, %s, %d, %d}", t.typ, t.val, t.line, t.col)
}

type TokenType int

const (
	// Single char token
	Plus TokenType = iota
	Minus
	Star
	Slash
	Bang
	Equal
	Less
	Greater

	// Braces
	LeftParan
	RightParan
	LeftSquare
	RightSquare
	LeftCurl
	RightCurl

	// Multi char token
	PlusEqual
	MinusEqual
	StarEqual
	BangEqual
	EqualEqual
	LessEqual
	GreaterEqual

	// Literal Token
	NumberLit
	StringLit

	// Keywords
	And
	Comment
	Const
	False
	Fun
	Let
	Nil
	Or
	Print
	Return
	Symbol
	True
)

func (typ TokenType) String() string {
	switch typ {
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case Bang:
		return "Bang"
	case Equal:
		return "Equal"
	case Less:
		return "Less"
	case Greater:
		return "Greater"
	case LeftParan:
		return "LeftParan"
	case RightParan:
		return "RightParan"
	case LeftSquare:
		return "LeftSquare"
	case RightSquare:
		return "RightSquare"
	case LeftCurl:
		return "LeftCurl"
	case RightCurl:
		return "RightCurl"
	case PlusEqual:
		return "PlusEqual"
	case MinusEqual:
		return "MinusEqual"
	case StarEqual:
		return "StarEqual"
	case BangEqual:
		return "BangEqual"
	case EqualEqual:
		return "EqualEqual"
	case LessEqual:
		return "LessEqual"
	case GreaterEqual:
		return "GreaterEqual"
	case NumberLit:
		return "NumberLit"
	case StringLit:
		return "StringLit"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Let:
		return "Let"
	case And:
		return "And"
	case Const:
		return "Const"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case Or:
		return "Or"
	case Symbol:
		return "Symbol"
	case Comment:
		return "Comment"
	case True:
		return "True"
	case Nil:
		return "Nil"
	default:
		return "Unknown"
	}
}

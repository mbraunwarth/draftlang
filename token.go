package main

type Token struct {
	typ  TokenType
	val  string
	line int
	col  int
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
	LeftBrace
	RightBrace
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
	TrueLit
	FalseLit
	NilLit

	// Keywords
	And
	Or
	Let
	Print
	Symbol
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
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
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
	case TrueLit:
		return "TrueLit"
	case FalseLit:
		return "FalseLit"
	case NilLit:
		return "NilLit"
	case Print:
		return "Print"
	case Let:
		return "Let"
	case And:
		return "And"
	case Or:
		return "Or"
	case Symbol:
		return "Symbol"
	default:
		return "Unknown"
	}
}

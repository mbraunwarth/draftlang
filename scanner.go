package main

import "unicode"

type Scanner struct {
	src   string // whole source code as text
	r     rune   // current rune
	line  int    // line counter
	col   int    // column counter
	pos   int    // current position in source string
	start int    // start position of current token
	errs  []ScanError
}

type ScanError struct {
	line int
	col  int
	err  string
}

func NewScanner(src string) *Scanner {
	return &Scanner{
		src:   src,
		r:     rune(src[0]),
		line:  1,
		col:   1,
		pos:   0,
		start: 0,
		errs:  make([]ScanError, 0),
	}
}

func (sc *Scanner) Scan() []Token {
	ts := make([]Token, 0)

	buf := make(chan Token)
	go sc.run(buf)

	for t := range buf {
		ts = append(ts, t)
	}

	return ts
}

func (sc *Scanner) run(buf chan<- Token) {
	for {
		sc.start = sc.pos
		lineStart := sc.col
		switch sc.r {
		case '+':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{PlusEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Plus, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '-':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{PlusEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Plus, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '*':
		case '/':
		case '!':
		case '=':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{EqualEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Equal, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '<':
		case '>':
		case '(':
		case ')':
		case '[':
		case ']':
		case '{':
		case '}':
		case '"':
			// scan here for string literals
		case '\n':
			// reset column counter and increment line count
			sc.col = 0
			sc.line = sc.line + 1
		default:
			// scan here for number literal, symbols or keywords
			switch {
			case unicode.IsNumber(sc.r):
				// TODO support for floating point numbers
				for unicode.IsNumber(sc.peek()) {
					sc.advance()
				}
				buf <- Token{NumberLit, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			case unicode.IsLetter(sc.r):
				for unicode.IsLetter(sc.peek()) {
					sc.advance()
				}
				val := sc.src[sc.start : sc.pos+1]
				// TODO check if val is keyword and if, which
				buf <- Token{Symbol, val, sc.line, lineStart}
			}
		}

		if !sc.advance() {
			break
		}
	}

	close(buf)
}

// advance scanner by one rune, return false if scanner is at end of code.
func (sc *Scanner) advance() bool {
	sc.pos = sc.pos + 1
	sc.col = sc.col + 1

	if len(sc.src) <= sc.pos {
		return false
	}

	sc.r = rune(sc.src[sc.pos])
	return true
}

func (sc *Scanner) peek() rune {
	if len(sc.src) < sc.pos {
		sc.errs = append(sc.errs, ScanError{sc.line, sc.col, "Expected symbol, found EOF."})
		var r rune
		return r
	}
	return rune(sc.src[sc.pos+1])
}

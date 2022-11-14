package main

import (
	"log"
	"unicode"
)

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
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{StarEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Star, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '/':
			if sc.peek() == '/' {
				// take comment lines as one token
				for sc.peek() != '\n' {
					sc.advance()
				}
				buf <- Token{Comment, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Slash, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '!':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{BangEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Bang, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '=':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{EqualEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Equal, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '<':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{LessEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Less, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '>':
			if sc.peek() == '=' {
				sc.advance()
				buf <- Token{GreaterEqual, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			} else {
				buf <- Token{Greater, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			}
		case '(':
			buf <- Token{LeftParan, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case ')':
			buf <- Token{RightParan, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case '[':
			buf <- Token{LeftSquare, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case ']':
			buf <- Token{RightSquare, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case '{':
			buf <- Token{LeftCurl, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case '}':
			buf <- Token{RightCurl, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
		case '"':
			// scan here for string literals
			for sc.peek() != '"' {
				if !sc.advance() {
					log.Fatal("Scan Error: Unterminated string literal.")
				}
			}
			// ignore surrounding quotes
			buf <- Token{StringLit, sc.src[sc.start+1 : sc.pos+1], sc.line, lineStart}
			sc.advance()
		case '\n':
			// reset column counter and increment line count
			sc.col = 0
			sc.line = sc.line + 1
		default:
			// scan here for number literal, symbols or keywords
			switch {
			case unicode.IsNumber(sc.r):
				for unicode.IsNumber(sc.peek()) {
					sc.advance()
				}

				// check if number is float
				if sc.peek() == '.' {
					// skip floating point
					sc.advance()
					// scan remaining number fraction
					for unicode.IsNumber(sc.peek()) {
						sc.advance()
					}
				}

				buf <- Token{NumberLit, sc.src[sc.start : sc.pos+1], sc.line, lineStart}
			case unicode.IsLetter(sc.r):
				for unicode.IsLetter(sc.peek()) {
					sc.advance()
				}

				// cache scanned symbol to check if it is a valid keyword
				val := sc.src[sc.start : sc.pos+1]
				var typ TokenType
				if IsKeyword(val) {
					typ = TypeFromKeyword(val)
				} else {
					typ = Symbol
				}

				buf <- Token{typ, val, sc.line, lineStart}
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

// peek returns the rune right after the current position.
func (sc *Scanner) peek() rune {
	if len(sc.src) < sc.pos {
		sc.errs = append(sc.errs, ScanError{sc.line, sc.col, "Expected symbol, found EOF."})
		var r rune
		return r
	}
	return rune(sc.src[sc.pos+1])
}

package main

func keywords() map[string]TokenType {
	return map[string]TokenType{
		"const":  Const,
		"false":  False,
		"fun":    Fun,
		"let":    Let,
		"print":  Print,
		"return": Return,
		"true":   True,
	}
}

// IsKeyword checks if the given string is a valid keyword.
func IsKeyword(k string) bool {
	ks := keywords()
	_, ok := ks[k]
	return ok
}

// GetTypeFromKeyword returns the TokenType value of the given keyword string.
// The function does not handle any error cases like if the given string is not a
// keyword.
func TypeFromKeyword(k string) TokenType {
	ks := keywords()
	return ks[k]
}

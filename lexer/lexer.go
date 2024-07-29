package lexer

import "strings"

func Tokenize(expr string) []string {
	expr = strings.ReplaceAll(expr, "(", " ( ")
	expr = strings.ReplaceAll(expr, ")", " ) ")
	return strings.Fields(expr)
}
package main

import "strings"

func ConcatenateWithAssign(str ...string) string {
	var sep, s string

	for _, s := range str {
		s += sep + s
		sep = " "
	}

	return s
}

func ConcatenateWithJoin(str ...string) string {
	return strings.Join(str, " ")
}

package minic

import "strings"

func Replace(text, old, new string) string {
	return strings.ReplaceAll(text, old, new)
}
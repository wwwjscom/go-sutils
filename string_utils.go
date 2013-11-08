// A module of string operations I do often, thus mkaing them more readable
package sutils

import "strings"

// Escapes all quotes in the given string
func EscapeAllQuotes(s string) string {
    return strings.Replace(s, "\"", "\\\"", -1)
}

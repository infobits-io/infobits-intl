package template

import (
	"strings"
	"text/template"
	"unicode"
)

// FuncMap returns the template functions.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"pascalCase":     PascalCase,
		"camelCase":      CamelCase,
		"snakeCase":      SnakeCase,
		"upperSnakeCase": UpperSnakeCase,
		"escapeString":   EscapeString,
		"escapeGoString": EscapeGoString,
		"upper":          strings.ToUpper,
		"lower":          strings.ToLower,
		"join":           strings.Join,
		"add":            func(a, b int) int { return a + b },
		"sub":            func(a, b int) int { return a - b },
	}
}

// PascalCase converts a string to PascalCase.
func PascalCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}

	return strings.Join(words, "")
}

// CamelCase converts a string to camelCase.
func CamelCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if len(word) > 0 {
			if i == 0 {
				words[i] = strings.ToLower(word)
			} else {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
	}

	return strings.Join(words, "")
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "_")
}

// UpperSnakeCase converts a string to UPPER_SNAKE_CASE.
func UpperSnakeCase(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}

	return strings.Join(words, "_")
}

// EscapeString escapes a string for use in Dart/TypeScript string literals.
func EscapeString(s string) string {
	var result strings.Builder

	for _, r := range s {
		switch r {
		case '"':
			result.WriteString(`\"`)
		case '\\':
			result.WriteString(`\\`)
		case '\n':
			result.WriteString(`\n`)
		case '\r':
			result.WriteString(`\r`)
		case '\t':
			result.WriteString(`\t`)
		case '$':
			result.WriteString(`\$`)
		default:
			result.WriteRune(r)
		}
	}

	return result.String()
}

// EscapeGoString escapes a string for use in Go string literals.
func EscapeGoString(s string) string {
	var result strings.Builder

	for _, r := range s {
		switch r {
		case '"':
			result.WriteString(`\"`)
		case '\\':
			result.WriteString(`\\`)
		case '\n':
			result.WriteString(`\n`)
		case '\r':
			result.WriteString(`\r`)
		case '\t':
			result.WriteString(`\t`)
		case '`':
			result.WriteString("` + \"`\" + `")
		default:
			result.WriteRune(r)
		}
	}

	return result.String()
}

// splitWords splits a string into words.
func splitWords(s string) []string {
	var (
		words       []string
		currentWord strings.Builder
	)

	for i, r := range s {
		if r == '_' || r == '-' || r == ' ' {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}

			continue
		}

		if unicode.IsUpper(r) && i > 0 {
			prevRune := rune(s[i-1])
			if unicode.IsLower(prevRune) {
				if currentWord.Len() > 0 {
					words = append(words, currentWord.String())
					currentWord.Reset()
				}
			}
		}

		currentWord.WriteRune(r)
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}

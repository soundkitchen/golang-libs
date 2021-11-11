package strings

import (
	"regexp"
	"strings"
)

var (
	matchSnake    = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToCamelCase converts snake case string to camel case.
func ToCamelCase(str string) string {
	return matchSnake.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// ToSnakeCase converts camel case string to snake case.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

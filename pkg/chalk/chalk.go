package chalk

import (
	"fmt"
)

// color is a type alias for string
type color string

// Color constants
const (
	reset  color = "\033[0m"
	red    color = "\033[31m"
	green  color = "\033[32m"
	yellow color = "\033[33m"
	blue   color = "\033[34m"
	purple color = "\033[35m"
	cyan   color = "\033[36m"
	white  color = "\033[37m"
)

func Red(a ...any) string {
	return colorize(red, a)
}

func Green(a ...any) string {
	return colorize(green, a)
}

func Yellow(a ...any) string {
	return colorize(yellow, a)
}

func Blue(a ...any) string {
	return colorize(blue, a)
}

func Purple(a ...any) string {
	return colorize(purple, a)
}

func Cyan(a ...any) string {
	return colorize(cyan, a)
}

func White(a ...any) string {
	return colorize(white, a)
}

// Colorize takes a string and a color and returns the string wrapped in the color
func colorize(color color, a []any) string {
	return fmt.Sprintf("%s%s%s", color, format(a), reset)
}

// Format takes a slice of any type and returns a string of the values separated by a space
func format(a []any) string {
	var s string
	for _, v := range a {
		s += fmt.Sprintf("%v", v)
		s += " "
	}

	s = s[:len(s)-1]
	return s
}

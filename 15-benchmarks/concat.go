package benchmarks

import "strings"

func ConcatPlus(n int) string {
	result := ""

	for i := 0; i < n; i++ {
		result += "x"
	}
	return result
}

func ConcatBuilder(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("x")
	}
	return sb.String()
}

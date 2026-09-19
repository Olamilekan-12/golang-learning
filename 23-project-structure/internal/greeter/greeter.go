package greeter

import (
	"fmt"
	"strings"
)

func shout(s string) string {
	return strings.ToUpper(s) + "!"
}

func Loud(name string) string {
	return shout(Hello(name))
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

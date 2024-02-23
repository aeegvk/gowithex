package iteration

import (
	"fmt"
	"strings"
)

func Repeat(character string, repetition int) string {
	return strings.Repeat(character, repetition)
}

func main() {
	fmt.Println(Repeat("a", 5))
}

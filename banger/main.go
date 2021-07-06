package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	word := os.Args[1]
	l := utf8.RuneCountInString(word)
	char := strings.Repeat("!", l)
	s := char + word + char
	s = strings.ToUpper(s)
	fmt.Println(s)

}

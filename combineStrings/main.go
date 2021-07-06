package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Jeewan" + " " + "Soori"
	fmt.Println("Hello", name)

	refer := "sd#$Sf"
	fmt.Println(utf8.RuneCountInString(refer))
	fmt.Println(len(refer))
}

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	path := `c:\program files\duper super\fun.txt
			 c:\program files\really\funny.png`
	fmt.Println(path)

	name := "İNANÇ"
	msg := `hi ` + name + `
	how are you ?`
	fmt.Println(msg)
	fmt.Print(utf8.RuneCountInString(name), "\n")

	animal := os.Args[1]
	fmt.Println(strings.ToUpper(animal))

	txt := `
	
	The weather looks good.
I should go and play.
	`
	fmt.Println(strings.TrimSpace(txt))

	friend := "inanç           "
	newName := strings.TrimRight(friend, " ")
	fmt.Println(utf8.RuneCountInString(newName))
}

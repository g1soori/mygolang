package main

import (
	"fmt"
	"strings"
)

func main() {
	start, stop := 'A', 'G'

	fmt.Printf("%-10s %-10s %-10s %-10s\n", "literal", "dec", "hex", "encoded")
	fmt.Println(strings.Repeat("-", 50))
	for i := 'A'; i <= stop; i++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -10x\n", i, string(i))
	}
}

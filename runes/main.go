package main

import "fmt"

func main() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	_ = bwords

	for _, v := range words {
		bytes := []byte(v)
		fmt.Println(bytes)

		bwords = append(bwords, []byte(v))
	}

	for _, v := range bwords {
		fmt.Printf("%s\n", v)
	}

	fmt.Printf("\n\n\n")

	const word = "console"

	var bytes []byte

	fmt.Printf("%-10s %-10s %-10s\n", "decimal", "hexa", "binary")
	for _, v := range word {
		fmt.Printf("%-10d %-10[1]x % -10[1]b\n", v)

		bytes = append(bytes, byte(v))
	}

	fmt.Println(bytes)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run maing.go <number>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	fmt.Printf("%5s", "X")

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}
}

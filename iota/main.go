package main

import "fmt"

func main() {
	const _ = iota
	const (
		_ = iota
		monday
		tuesday
		wednesday
	)

	fmt.Println(monday, tuesday, wednesday)
}

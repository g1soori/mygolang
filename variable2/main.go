package main

import "fmt"

func main() {
	i, f := 314, 3.14
	s := "Hello"
	b := true

	sum := 27 + 3.5

	//var on bool
	on, _ := true, true

	fmt.Printf("i: %d f:%g s:%s b:%t\n", i, f, s, b)
	fmt.Println(sum)
	fmt.Println(on)
}

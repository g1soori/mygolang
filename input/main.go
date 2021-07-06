package main

import (
	"fmt"
	"os"
)

func main() {
	count, name := len(os.Args), "jeewan"
	it := false

	fmt.Printf("%#v\n", os.Args)
	fmt.Println(count)
	fmt.Println("my path is ", os.Args[0])
	fmt.Printf("Tyoe is %v\n", name)

	fmt.Printf("%t\n", it)

	hello := "Hello world"
	fmt.Printf("\"Hello world\"\n")
	fmt.Printf("%q\n", hello)

	fmt.Printf("Type of %s is %[1]T\n", "hello00")

	var ratio float64 = 3 / 2.0
	fmt.Println(ratio)
	fmt.Printf("%T\n", ratio)
}

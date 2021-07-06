package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	fmt.Println(arg)
	cel, _ := strconv.ParseFloat(arg, 64)
	far := cel*1.8 + 32
	fmt.Printf("Temperature is %g\n", far)
}

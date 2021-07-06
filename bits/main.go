package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	num, _ := strconv.Atoi(arg)
	fmt.Printf("%08b\n", num)
}

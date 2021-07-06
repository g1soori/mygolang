package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arg := os.Args
	const usage = "[Usage] gp run maing.go 2"

	if len(arg) != 2 {
		fmt.Println(usage)
		return
	}

	if n, err := strconv.Atoi(arg[1]); err != nil {
		fmt.Printf("Wrong input %q\n", arg[1])
	} else if n < 1 {
		fmt.Println("invalid number")
	} else if n > 17 {
		fmt.Println("R-Rated")
	} else if n < 13 {
		fmt.Println("PG-Rated")
	} else {
		fmt.Println("PG-R Rated")
	}
}

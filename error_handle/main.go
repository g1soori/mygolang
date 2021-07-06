package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const f2m = 0.3048

	f, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("Wrong input:", err)
	} else {
		m := f * f2m
		fmt.Printf("%f\n", m)
	}
}

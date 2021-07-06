package main

import (
	"fmt"
	"time"
)

func main() {

	switch t := time.Now().Hour(); true {

	case t < 12:
		fmt.Print("Morning\n")
	case t >= 12 && t < 18:
		fmt.Print("Afternoon\n")
	default:
		fmt.Print("Night\n")
	}
}

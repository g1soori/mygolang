package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	n := "2 4 6 1 3 5"

	ns := strings.Fields(n)

	var ni []int

	for _, v := range ns {
		i, err := strconv.Atoi(v)
		if err == nil {
			ni = append(ni, i)
		}
	}

	fmt.Printf("nums: %d\n", ni)

	even := ni[:3]
	fmt.Printf("Even no: %d\n", even)

	odd := ni[3:]
	fmt.Printf("Even no: %d\n", odd)

	fmt.Printf("last 2 no: %d\n", ni[len(ni)-2:])

	fmt.Printf("Even last: %d\n", even[len(odd)-1:])
}

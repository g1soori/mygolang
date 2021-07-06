package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var (
		friends []string
		//distance []float64
		//data []uint
	)

	friends = []string{"Lahiru", "Dum", "Harsha"}

	//fmt.Printf("names: %s ", friends)
	fmt.Printf("names: %10T %5d %5t\n", friends, len(friends), friends == nil)

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	A := strings.Split(namesA, ", ")

	fmt.Printf("%T\n", A)
	fmt.Printf("%T\n", namesB)

	sort.Strings(A)
	sort.Strings(namesB)

	if len(A) == len(namesB) {
		fmt.Println("length is equal")

	}

	png, header := []byte{'P', 'N', 'G'}, []byte{}
	var name []byte

	header = append(header, png...)
	name = append(name, png...)

	if bytes.Equal(header, png) {
		fmt.Println("Slices are equal")
	}

	fmt.Printf("%s\n", name)
}

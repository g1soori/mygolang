package main

import "fmt"

func main() {
	names := []string{"jeewan", "soori", "John"}
	names = append(names, names[1:]...)
	fmt.Println(names)
	names = append(names[0:1], "mike", "steve")
	fmt.Println(names, cap(names))
	names = names[0:5]

	fmt.Println(names[0:cap(names)])

	fmt.Println(names[0:5], len(names))

	names = names[4:5]
	fmt.Println(names, len(names), cap(names))
}

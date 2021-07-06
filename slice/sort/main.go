package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	fmt.Println()
	fmt.Println("Middle items  :", items[5:8])
	sort.Strings(items[5:8])
	fmt.Println("Sorted  :", items)

}

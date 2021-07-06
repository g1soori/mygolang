package main

import (
	"fmt"
)

func main() {
	phn := map[string]string{
		"Soori": "98829",
		"Weer":  "32134",
	}

	prodAvail := map[int]bool{
		1: true,
		2: false,
		3: true,
	}

	phns := map[string][]string{
		"Amal":  {"934134", "3243"},
		"Kasun": {"234243", "32432"},
	}

	busketId := map[int]map[int]int{
		100: {617841573: 4, 576872813: 2},
		101: {576872813: 5, 657473833: 20},
	}

	fmt.Println(phn["Soori"])
	fmt.Println("Is prod ID 1 exist: ", prodAvail[1])
	fmt.Println("Amal's second no is: ", phns["Amal"][1])
	fmt.Println("How many of 576872813 the customer 101 is going to buy? ", busketId[101][576872813])
}

package main

import "fmt"

//have to import fmt again here even though its already imported in other files
//its because import clause is file scope not package scope

func bye() {
	fmt.Println("Bye........")
}

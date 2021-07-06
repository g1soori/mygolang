package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

func main() {

	g := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 30}, genre: "sandbox"},
	}

	msg := `Inanc's game store has 3 games.

	> list   : lists all the games
	> quit   : quits`

	fmt.Println(msg)

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		switch in.Text() {
		case "list":
			for _, v := range g {
				fmt.Printf("%+v %-20s\n", v, "("+v.genre+")")
			}
		case "quit":
			return
		default:
			fmt.Println("wrong input")
		}

	}

}

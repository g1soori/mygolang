package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [...][3]string{
		{"awesome", "fantastic", "great"},
		{"sad", "bad", "lost"},
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run <your name> <positive|negative>")
		return
	}
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(len(moods[0]))
	if os.Args[2] == "positive" {
		fmt.Println(moods[0][i])
	} else if os.Args[2] == "negative" {
		fmt.Println(moods[1][i])
	} else {
		fmt.Println("invalid mood")
	}

}

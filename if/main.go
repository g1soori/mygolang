package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	user, password := "jack", 1888
	const (
		usage   = "Usage: [username] [password]"
		denied  = "Access denied for %q \n"
		granted = "Access granted for %q \n"
	)

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	pwd, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(denied)
		return
	}

	if os.Args[1] == user && pwd == password {
		fmt.Printf(granted, user)
	} else {
		fmt.Printf(denied, os.Args[1])
	}

}

package main

import (
	"fmt"
	"path"
)

func main() {

	_, file := path.Split("/var/spool/mail")
	fmt.Println(file)
}

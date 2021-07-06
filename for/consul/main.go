package main

import (
	"fmt"
	"os"
)

const output = `
  - id: %s
    address: %[1]s
    name: 
    port: 80
    checks:
      - name: Check HTTP Port
        interval: 5s
        http: http://%[1]s:80
	`

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Enter space separated IP address")
		return
	}

	for _, v := range os.Args[1:] {
		fmt.Printf(output, v)
		fmt.Println()
	}

}

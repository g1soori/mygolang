package main

import "fmt"

//fmt is file scoped. It can only be used within this file

const ok = true

//ok and main func are package scoped. Visible to all the files in the package

func main() {
	var name = "Jeewan"
	//name is a blocked scoped. It can't be seen outside this block

	fmt.Println(name, ok)

}

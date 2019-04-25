package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "world"
	fmt.Println(x == y)

	var z string = "hello"
	fmt.Println(x == z)
}
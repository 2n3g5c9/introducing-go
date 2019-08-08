package main

import "fmt"

func half(number int) (int, bool) {
	return number/2, number%2==0
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
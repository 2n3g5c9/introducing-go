package main

import "fmt"

func half(number int) (int, bool) {
	halfNumber := number / 2
	if (number % 2) == 0 {
		return halfNumber, true
	} else {
		return halfNumber, false
	}
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
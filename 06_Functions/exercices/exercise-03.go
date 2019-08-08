package main

import "fmt"

func greatest(numbers ...int) int {
	greatestNumber := numbers[0]
	for _, number := range numbers {
		if number > greatestNumber {
			greatestNumber = number
		}
	}
	return greatestNumber
}

func main() {
	fmt.Println(greatest(11, 91, 45, 53, 67, 94))
}
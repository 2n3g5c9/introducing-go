package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) time.Time {
	return <-time.After(d)
}

func main() {
	go func() {
		Sleep(5 * time.Second)
		fmt.Println("Hi from the sub goroutine")
	}()

	fmt.Println("Hi from the main goroutine")

	var input string
	fmt.Scanln(&input)
}
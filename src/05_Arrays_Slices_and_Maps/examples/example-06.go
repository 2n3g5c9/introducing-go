package main

import "fmt"

func main() {
	i := make(map[string]int)
	i["key"] = 10
	fmt.Println(i["key"])
}
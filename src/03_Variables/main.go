package main

import "fmt"

// Scope
var k string = "Hello, World"

func main() {
	var a string = "Hello, World"
	fmt.Println(a)

	var b string
	b = "Hello, World"
	fmt.Println(b)

	var c string
	c = "first"
	fmt.Println(c)
	c = "second"
	fmt.Println(c)

	var d string = "hello"
	var e string = "world"
	fmt.Println(d == e)

	var f string = "hello"
	var g string = "hello"
	fmt.Println(f == g)

	h := "Hello, World"
	fmt.Println(h)

	var i = "Hello, World"
	fmt.Println(i)

	j := 5
	fmt.Println(j)

	// How to Name a Variable
	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	l()

	// Constants
	const m string = "Hello, World"
	fmt.Println(m)

	// Defining multiple variables
	var (
		n = 5
		o = 10
		p = 15
	)
	fmt.Println(n, o, p)

	// An example program
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func l() {
	fmt.Println(k)
}

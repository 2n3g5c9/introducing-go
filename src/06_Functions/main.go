package main

import "fmt"

// Your Second Function
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Variadic Functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Closure
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// defer, panic and recover
func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	// Your Second Function
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))

	someOtherName := []float64{98,93,77,82,83}
	fmt.Println(average(someOtherName))

	fmt.Println(add(1, 2, 3))
	ys := []int{1,2,3}
	fmt.Println(add(ys...))

	// Closure
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	// Recursion
	fmt.Println(factorial(5))

	// Pointers
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	// defer, panic, and recover
	defer second()
	first()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

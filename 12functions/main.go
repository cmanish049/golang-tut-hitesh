package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions in Go")

	foo()
	bar()
	fmt.Println("Sum of 2 and 3 is", sum(2, 3))

	total, status := proSum(2, 3, 4, 5, 6)
	fmt.Println("Sum of 2, 3, 4, 5, 6 is", total)
	fmt.Println("Status is", status)
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func sum(x int, y int) int {
	return x + y
}

func proSum(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Success"
}

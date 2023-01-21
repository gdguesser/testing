package main

import "fmt"

func main() {
	someValue := SomeValue()

	addition := Add(3, 2)

	multiplication := Multiply(5, -6)

	fmt.Printf("Some value: %s, Addition: %d, Multiplication: %d\n", someValue, addition, multiplication)
}

func SomeValue() string {
	return "expected value"
}

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

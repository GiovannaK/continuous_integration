package main

import "fmt"

func main() {
	result := sum(10, 10)
	fmt.Println("The sum is:", result)
}

func sum(a, b int) int {
	return a + b
}

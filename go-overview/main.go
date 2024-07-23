package main

import "fmt"

func main() {
	x := 42
	var y *int
	y = &x

	fmt.Printf("The value of x: %v", x)
	*y = 10
	fmt.Printf("The value of x: %v", x)
}

package main

import "fmt"

func main() {
	if x := 10; x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	y := 5

	switch y {
	case 1: // y == 1
		fmt.Println("y is 1")
	case 5: // y == 5
		fmt.Println("y is 5")
	default:
		fmt.Println("y is not 1 or 5")
	}

	// for
	//
	//

	// 1. for as a for
	//
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// 2. for as for range
	//

	// 3. for as while
	//
	//
	j := 5
	for j < 5 {
		fmt.Println(j)
		j++
	}
}

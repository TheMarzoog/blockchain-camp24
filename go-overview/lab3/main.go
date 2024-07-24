package main

import "fmt"

func main() {
	// Arrays fix size, type
	//
	var arr [5]int

	fmt.Println(arr)
	arr[0] = 1
	fmt.Println(arr)

	for _, v := range arr {
		fmt.Printf("Value %d\n", v)
	}

	// Sices fix type
	//
	sl := []int{1, 2, 3}

	sl = append(sl, 4, 5)

	fmt.Println(sl)

	capitals := map[string]string{
		"Saudi Arabia": "Riyadh",
		"France":       "Paris",
		"Japan":        "Tokyo",
	}

	capitals["Kuwait"] = "Kuwait City"

	delete(capitals, "Japan")

	for key, value := range capitals {
		fmt.Printf("The key %v, The value %v\n", key, value)
	}
}

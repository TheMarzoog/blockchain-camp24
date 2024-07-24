package main

import "fmt"

func main() {
	favoriteNumbers := [5]int{3, 7, 9, 14, 21}
	fmt.Printf("Favorite Numbers: %v\n", favoriteNumbers)

	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "date", "orange")
	fmt.Println("Fruits:", fruits)

	capitals := map[string]string{
		"Saudi Arabia": "Riyadh",
		"France":       "Paris",
		"Japan":        "Tokyo",
	}
	capitals["Kuwait"] = "Kuwait City"
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}
	delete(capitals, "France")

	fmt.Printf("capitals: %v", capitals)
}

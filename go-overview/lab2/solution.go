package main

import "fmt"

func main() {
	number := 5
	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	day := "Weday"
	switch day {
	case "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println("It's a weekday.")
	case "Friday", "Saturday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day.")
	}

	number2 := 5
	factorial := 1
	for i := 1; i <= number2; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d is %d\n", number2, factorial)
}

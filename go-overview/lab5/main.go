package main

import (
	"fmt"
	"lab5/model/duck"
)

func main() {
	eimad := duck.New("Eimad", 2, "Yellow", 3, 12.5)

	eimad.ChangeColor("Red")

	fmt.Printf("Eimad : %#v\n\n", eimad)

	fmt.Println(eimad)

}

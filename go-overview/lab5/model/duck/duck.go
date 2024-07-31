package duck

import "fmt"

type Duck struct {
	Name   string
	Age    int
	Color  string
	Size   int
	Weight float64
}

func New(name string, age int, color string, size int, weight float64) Duck {
	return Duck{
		Name:   name,
		Age:    age,
		Color:  color,
		Size:   size,
		Weight: weight,
	}
}

func (d *Duck) ChangeColor(newColor string) {
	d.Color = newColor
}

func (d *Duck) String() string {
	return fmt.Sprintf("Name: %v, Age: %v, Color: %v", d.Name, d.Age, d.Color)
}

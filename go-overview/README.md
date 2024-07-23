# Go Programming Language Overview
Go is a statically typed, compiled language that offers an excellent balance between simplicity and performance. Designed for efficiency and ease of use, it is ideal for developing scalable and high-performance applications.

## Go Key Features
- Simple Syntax
- Concurrency Support
- Fast Compilation
- Garbage Collection
- Static Typing

## Hello, world!

Let's write a simple "Hello, World!" program in Go.

1. Create a `main.go` file.
2. Write the following code.
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
3. Running the program
```sh
go build main.go
./main
## Or use the following command
go run main.go
```

## Variable Declaration and Initialization
In Go, variables can be declared and initialized in several ways.
### Declaration with var
Syntax
```go
var variableName type
variableName = value
```
Example
```go
var name string
name = "Mohammed"
var age int = 20
```
### Type Inference
Syntax
```go
var variableName = value
```
Example
```go
var greeting = "Hello, World!" // Type inferred as string
var number = 42                // Type inferred as int
```
### Short Variable Declaration
Syntax
```go
variableName := value
```
Example
```go
message := "Go is awesome!" // Type inferred as string
count := 100               // Type inferred as int
```

### Multiple Varialbe Declarations
Example
```go
var name, age := "Mohammed", 20
```

### Zero Values
If a variable is declared without an initial value, it is automatically assigned a zero value based on its type.
Example
```go
var count int       // count is initialized to 0
var text string     // text is initialized to an empty string ""
var flag bool       // flag is initialized to false
```

## Pointers
#### Pointer Operators

1. **Address-of Operator (`&`)**:
  - The `&` operator is used to obtain the memory address of a variable.

2. **Dereferencing Operator (`*`)**:
  - The `*` operator is used to access the value stored at the memory address.

#### Syntax

To declare a pointer, use the `*` symbol before the type:
```go
var ptr *int
```

#### Example
```go
var x int = 10
var ptr *int

// Get the address of x using the & operator
ptr = &x

// Print the address stored in ptr
fmt.Println("Address of x:", ptr)

// Print the value stored at the address in ptr using the * operator
fmt.Println("Value of x:", *ptr)

// Change the value at the address in ptr
*ptr = 20

// Print the new value of x
fmt.Println("New value of x:", x)
```

## Control Statements

### if and else

#### Syntax
```go
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
```
#### Example
```go
number := 10

	if number%2 == 0 {
    fmt.Println("The number is even.")
	} else {
    fmt.Println("The number is odd.")
	}
}
```
Other way
```go
if number := 10; number %2 == 0 {
	fmt.Println("The number is even.")
} else {
	fmt.Println("The number is odd.")
}

```

### switch
#### Syntax
```go
switch value {
case option1:
    // code to execute if value == option1
case option2:
    // code to execute if value == option2
default:
    // code to execute if value does not match any option
}
```

## Functions
#### Syntax
One return
```go
func functionName(parameterName parameterType) returnType {
    // function body
    return returnValue
}
```
Multiple returns
```go
func functionName(parameterName parameterType) (returnType1, returnType2) {
    // function body
    return returnValue1, returnValue2
}
```
#### Example
One return
```go
func add(a int, b int) int {
    return a + b
}
```
Multiple returns
```go
func divide(a int, b int) (int, error) {
	if b == 0 {
    return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
```

## Arrays and Slices

### Arrays
Arrays in Go are fixed-size, ordered collections of elements of the same type. The size of an array is part of its type.

#### Syntax
```go
var arrayName [size]elementType
```
#### Example
```go
arr := [5]int{10, 20, 30, 40, 50}
fmt.Println("Array:", arr)
```

### Slices
Slices are more flexible and powerful than arrays. They are dynamically-sized, and they provide a convenient and efficient way to work with sequences of elements.

#### Syntax
```go
var sliceName []elementType
```
#### Example
```go
// Create and initialize a slice
slice := []int{10, 20, 30, 40, 50}
fmt.Println("Slice:", slice)

// Append values to the slice
slice = append(slice, 60, 70)
fmt.Println("Modified Slice:", slice)

// Slice a slice
subSlice := slice[1:4]
fmt.Println("Sub-Slice:", subSlice)
```

## Maps

Maps in Go are unordered collections of key-value pairs. They provide a fast and efficient way to look up data based on a custom key.

### Syntax:
```go
var mapName map[keyType]valueType
```

## Loops
Go only has a for loop, but the for does it all.

### for as a for
#### Syntax
```go
for initialization; condition; post {
    // code to execute in each iteration
}
```
#### Exmaple
```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

### for as a for range
#### Syntax
```go
for index, value := range collection {
    // code to execute for each element
}
```
#### Example:
```go
numbers := []int{1, 2, 3 , 4, 5}
for i, v := numbers {
	fmt.Printf("Index: %d, Value %d\n", i, v)
}
```

### for as a while
#### Syntax
```go
for condition {
    // code to execute while condition is true
}
```

#### Example
```go
count := 0
for count < 5 {
	fmt.Println("Count: ", count)
	count++
}
```
## Structs
#### Syntax
```go
type StructName struct {
    FieldName1 FieldType1
    FieldName2 FieldType2
    // more fields...
}
```
#### Example
```go
type Duck struct {
    Name  string
    Age   int
    Color string
}
```
## Methods
#### Syntax
```go
func (receiverType receiverName) methodName(parameterName parameterType) returnType {
    // method body
    return returnValue
}
```
#### Example
```go
// Method to display the details of the Duck
// This method has a value receiver, so it gets a copy of the Duck struct
func (d Duck) DisplayDetails() {
    fmt.Printf("Name: %s, Age: %d, Color: %s\n", d.Name, d.Age, d.Color)
}

// Method to change the color of the Duck
// This method has a pointer receiver, so it can modify the original Duck struct
func (d *Duck) ChangeColor(newColor string) {
    d.Color = newColor
}
```

## THE END

For a more comprehensive overview of Go, be sure to check out the [Tour of Go](https://tour.golang.org/). The Tour of Go is an interactive introduction to the Go programming language, providing detailed explanations and hands-on exercises to help you deepen your understanding of Go's features and best practices.

Thank you for following along with this tutorial, and happy coding!

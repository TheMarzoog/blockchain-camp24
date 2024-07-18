# Go Programming Language Overview:
Go is a statically typed, compiled language that offers an excellent balance between simplicity and performance. Designed for efficiency and ease of use, it is ideal for developing scalable and high-performance applications.

## Go Key Features:
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
TODO

## Control Statements

### if and else

#### Syntax:
```go
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}
```

### switch
#### Syntax:
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

## Arrays and Slices
TODO
## Loops
Go only have a for loop, but the for does it all.

### for as a for
TODO
### for as a for range
TODO
### for as a while
TODO

## Functions
TODO
## Structs
TODO
## Methods
TODO


## THE END
TODO

## Lab 4: Structs, Methods, and Go Modules

### Objective
Learn to define and use structs and methods to create and manipulate custom data types, and understand how to use Go modules.

### Tasks
1. **Set Up Go Module**
  - Create a new directory for the lab.
  - Initialize a new Go module called `lab4` using the following command
  ```sh
  go mod init lab4
  ```

2. **Define a Struct**
  - Create a directory called `model` and within it, create a directory called `horse`.
  - Create a file named `horse.go` in the `model/horse` directory.
  - Define a `Horse` struct with fields for name, breed, age, and speed.

3. **Add Methods to the Struct**
  - Create a constructor function `New` that returns a new horse.
  - Add a `String` method to print a horse's details.
  - Add a method `UpdateSpeed` to update the speed of the horse and return an error if the speed is negative.

4. **Use the Struct and Methods**
  - Create a `main.go` file in the `lab4` directory.
  - Import the `horse` package and use the `Horse` struct and its methods.

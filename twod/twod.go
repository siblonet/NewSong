package twod

import "fmt"

func Towd() {
	fmt.Println("memory address")

	// Declare a variable
	var value int = 42

	// Get the memory address of the variable using the '&' operator
	pointer := &value

	// Print the memory address (pointer) of the variable
	fmt.Printf("Memory Address: %p\n", pointer)

	// You can also directly print the pointer without using fmt.Printf
	// fmt.Println("Memory Address:", pointer)
}

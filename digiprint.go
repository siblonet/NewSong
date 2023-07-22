package main

import "fmt"

func Digiprint() {
	for char := 'a'; char <= 'q'; char++ {
		fmt.Println(string(char))
	}
	fmt.Println() // Print a newline after printing the alphabets
}

package main

func IterativeFactorial(nb int) int {
	// Factorial of 0 is 1.
	if nb == 0 {
		return 1
	}

	// Initialize the result to 1.
	result := 1

	// Calculate the factorial iteratively.
	for i := 1; i <= nb; i++ {
		// Check for overflow.
		if result > (1<<31-1)/i {
			return 0
		}
		result *= i
	}

	return result
}

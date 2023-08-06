package main

func LeftorSift(nb int) int {
	if nb < 0 || nb > 4 {
		return 0 // Return 0 for negative values (error)
	}

	result := 1
	for i := 1; i <= nb; i++ {
		if result > (1<<31-1)/i {
			return 0 // Return 0 if the factorial overflows
		}
		result *= i
	}

	return result
}

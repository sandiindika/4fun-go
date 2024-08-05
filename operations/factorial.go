package operations

func Factorial(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 || n == 1 {
		return 1
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

package operations

func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	series := make([]int, n)
	series[0] = 0

	if n > 1 {
		series[1] = 1
		for i := 2; i < n; i++ {
			series[i] = series[i-1] + series[i-2]
		}
	}
	return series
}

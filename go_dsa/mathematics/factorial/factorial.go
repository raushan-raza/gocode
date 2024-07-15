package fact

func Factorial(n int) int {
	res := 0
	for i := 5; i <= n; i = i * 5 {
		res = res + n/i
	}

	return res
}

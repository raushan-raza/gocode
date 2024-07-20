package fact

import (
	"fmt"
)

func Factorial(n int) int {
	res := 0
	for i := 5; i <= n; i = i * 5 {
		res = res + n/i
	}

	return res
}

func Prime(n int) {

	var flag int = 0

	if n == 2 || n == 3 {
		fmt.Println(n)
		flag = 1
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flag = 1
			break
		}
	}

	if flag != 1 {
		fmt.Println(n)
	}
}

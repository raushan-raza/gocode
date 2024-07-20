package main

import "fmt"

func main() {

	var x, n, result int

	fmt.Println("Enter x and n to calculate x^n")
	fmt.Scan(&x, &n)

	result = 1

	//neive solution
	// for i := 0; i < n; i++ {
	// 	result = result * x
	// }

	//binary iterative logic solution
	for n > 0 {
		if n%2 != 0 {
			result = result * x
		}
		n = n / 2
		x = x * x
	}

	fmt.Println("power of x^n = ", result)
}

// to calculate trailing zeros in factorial of number N
package main

import (
	"fmt"
	"mathematics/MyPackages"
)

func main() {
	var N int
	fmt.Println("enter N")
	fmt.Scan(&N)
	trailing_zero := MyPackages.Factorial(N)

	fmt.Printf("number of tariling zeros in factorial of %d = %d\n", N, trailing_zero)
}

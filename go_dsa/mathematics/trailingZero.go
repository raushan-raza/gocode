// to calculate trailing zeros in factorial of number N
package main

import (
	"fmt"
	fact "go_dsa/factorial"
)

func main() {
	var N int
	fmt.Println("enter N")
	fmt.Scan(&N)
	trailing_zero := fact.Factorial(N)

	fmt.Printf("number of tariling zeros in factorial of %d = %d\n", N, trailing_zero)
}

// Find all prime numbers less than or equal to given number n

package main

import (
	"fmt"
	"mathematics/MyPackages"
	"os"
)

func main() {
	var n int
	fmt.Println("Please enater a Number")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Please enter a valid number greater than 1")
		os.Exit(0)
	}

	for i := 2; i <= n; i++ {
		MyPackages.Prime(i)
	}

}

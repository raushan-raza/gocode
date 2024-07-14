package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2 int
	fmt.Println("enter n1 and n2")
	fmt.Scan(&n1, &n2)

	var smallest_num int

	if n1 == 0 || n2 == 0 {
		fmt.Println("enter valid numbers")
		os.Exit(0)
	}

	if n1 < n2 {
		smallest_num = n1
	} else {
		smallest_num = n2
	}

	// for i := 1; i <= smallest_num; i++ {
	// 	if n1%i == 0 && n2%i == 0 {
	// 		GCD = i
	// 	}
	// }

	for smallest_num > 0 {
		if n1%smallest_num == 0 && n2%smallest_num == 0 {
			break
		}
		smallest_num--
	}
	fmt.Printf("greatest common divisor id %d\n", smallest_num)
}

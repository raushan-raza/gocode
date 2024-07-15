package main

import (
	"fmt"
)

func main() {
	var n1, n2, lcm int

	fmt.Println("enter n1 and n2 to calculate lcm")
	fmt.Scan(&n1, &n2)

	if n1 > n2 {
		lcm = n1
	} else {
		lcm = n2
	}

	for i := 2; i > 0; i++ {
		lcm = lcm * i
		if lcm%n1 == 0 && lcm%n2 == 0 {
			break
		}
		lcm = lcm / i
	}

	fmt.Printf("LCM of %d and %d is = %d\n", n1, n2, lcm)
}

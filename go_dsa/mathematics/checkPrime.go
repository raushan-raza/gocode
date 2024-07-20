package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	flag := 0
	fmt.Println("Enter numcer to check if prime")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("not prime")
		os.Exit(0)
	} else if n == 2 || n == 3 {
		fmt.Println("prime number")
		os.Exit(0)
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flag = 1
			break
		}
	}

	if flag == 1 {
		fmt.Println("not prime")
	} else {
		fmt.Println("prime")
	}
}

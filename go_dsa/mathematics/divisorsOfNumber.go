package main

import "fmt"

func main() {
	var N, i int
	fmt.Println("enter number")
	fmt.Scan(&N)

	for i = 1; i*i < N; i++ {
		if N%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}

	for ; i >= 1; i-- {
		if N%i == 0 {
			fmt.Printf("%d\n", N/i)
		}
	}
}

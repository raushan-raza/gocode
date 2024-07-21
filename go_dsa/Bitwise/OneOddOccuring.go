package main

import (
	Bitpackage "Bitwise/MyBitPackages"
	"fmt"
)

func main() {
	var size int
	fmt.Println("enter the size of slice")
	fmt.Scan(&size)

	slice := make([]int, size)
	fmt.Println("Enter slice of size %d", size)
	for i := 0; i < size; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Printf("that one odd element is %d\n", Bitpackage.FindOddElement(slice))
}

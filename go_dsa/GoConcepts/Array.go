package main

import "fmt"

func main() {
	// Basic declaration
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println("arr1:", arr1)

	// Declaration with initialization
	var arr2 = [3]int{4, 5, 6}
	fmt.Println("arr2:", arr2)

	// Short declaration
	arr3 := [3]int{7, 8, 9}
	fmt.Println("arr3:", arr3)

	// Implicit length
	arr4 := [...]int{10, 11, 12}
	fmt.Println("arr4:", arr4)

	// Iterating with for loop
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("arr4[%d] = %d\n", i, arr4[i])
	}

	// Iterating with for range loop
	for index, value := range arr4 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Multi-dimensional array
	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	// 2-dimensional array
	// Using shorthand declaration
	// arr := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
	//     {"C++", "Go", "HTML"}}

	// creating a copy of an array by value
	// arr := arr1

	// Creating a copy of an array by reference
	// arr := &arr1

	fmt.Println("matrix:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

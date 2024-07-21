package main

import (
	"fmt"
	"mathematics/MyPackages"
)

func main() {

	var x, k int
	fmt.Println("please enter x and k")
	fmt.Scan(&x, &k)

	res := MyPackages.Isset(x, k)

	if res {
		fmt.Println("SET")
	} else {
		fmt.Println("Not SET")
	}
}

package main

import "fmt"

var x int = 42
var y string = "james bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}

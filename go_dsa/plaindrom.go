// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message
//find if plaindrome if x >=0

package main

import "fmt"

func main() {
	var x int
	fmt.Println("enter value")
	fmt.Scan(&x)
	//   if x < 10{
	//       fmt.Printf ("given %d is a plandrome\n",x)
	//       return
	//   }
	temp := x
	reverse := 0
	for temp > 0 {
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}
	if reverse != x {
		fmt.Printf("given %d is anot  drome\n", x)
	} else {
		fmt.Printf("given %d is a plandrome\n", x)
	}

	return
}

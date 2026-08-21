package main

import (
	"fmt"
)

func main() {
	//	n := 20
	//	for i := 0; i <= n; i++ {
	//		println(i)
	//	}
	x := 505
	o := x
	r := 0
	for x > 0 {
		d := x % 10
		r = r*10 + d
		x = x / 10

	}
	if r == o {
		fmt.Println(r, "palindrome")
	} else {
		fmt.Println("not palindrome")
	}

}

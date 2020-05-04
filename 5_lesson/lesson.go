package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n) // 100

	// fmt.Println(&n) // 0xc00008a008

	// var p *int = &n

	// fmt.Println(p) // 0xc00008a008

	// fmt.Println(*p) // 100
}

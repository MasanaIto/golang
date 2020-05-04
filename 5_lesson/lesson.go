package main

import "fmt"

func main() {

	// make は ポインタを返さない
	s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int

	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int

	ch := make(chan int) // chan int
	fmt.Printf("%T\n", ch)

	// new は ポインタを返す
	var p *int = new(int)
	fmt.Printf("%T\n", p) // *int

	var st = new(struct{}) // *struct {}
	fmt.Printf("%T\n", st)
}

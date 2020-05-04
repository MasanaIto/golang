package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// struct の Person を書き換えたい場合、Stringerを使う
func (p Person) String() string {
	return fmt.Sprintf("My Name is %v.", p.Name)
}

func main() {
	masana := Person{"Masana", 20}
	fmt.Println(masana) // My Name is Masana.
}

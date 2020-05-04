package main

import (
	"fmt"
)

/*
Vertex ...
*/
type Vertex struct {
	X, Y int
}

/*
Plus ... Q1. 以下に、7と表示されるメソッドを作成してください。
*/
func (v Vertex) Plus() int {
	return v.X + v.Y
}

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Plus()) // 7
// }

// Q2 X is 3! Y is 4! と表示されるStringerを作成してください。
func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v) // X is 3! Y is 4!
}

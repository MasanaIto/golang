package main

import "fmt"

type Vertex struct {
	// 宣言は外部読み込みする場合は大文字でかく
	X int
	Y int
	// X, Y int でもok
	S string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)        // {1 2}
	fmt.Println(v.X, v.Y) // 1 2
	v.X = 100
	fmt.Println(v.X, v.Y) // 100 2

	v2 := Vertex{X: 1}
	fmt.Println(v2) // {1 0}

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3) // {1 2 test}

	v4 := Vertex{}
	fmt.Println(v4) // {0 0 }

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5) // main.Vertex &{0 0 }

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6) // *main.Vertex &{0 0 }

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7) // *main.Vertex &{0 0 }

}

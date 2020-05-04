package main

import "fmt"

/*
Vertex ... 座標X, Yを扱う
*/
type Vertex struct {
	X, Y int
}

/*
Area ... 値レシーバー
*/
func (v Vertex) Area() int {
	return v.X * v.Y
}

/*
Scale ... ポインタレシーバー
アスタリスクをつけることで、Vertexの値の中身を書き換えることができる
*/
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}

	// メソッドは .（ドット）で呼び出せる
	fmt.Println(v.Area()) // 12
	v.Scale(10)
	fmt.Println(v.Area()) // 1200
}

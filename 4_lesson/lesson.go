package main

import (
	"fmt"
	"os"
)

// defer ... main関数の処理が終わったら呼び出される

func foo() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// lesson.go ファイルの 100byte分だけ出力する
func main() {
	foo()
	file, _ := os.Open("./lesson.go")
	// fileをオープンすると、クローズする処理を入れなくてはならない
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

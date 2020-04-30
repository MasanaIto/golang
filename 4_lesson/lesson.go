package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}

	// if文を省略した場合
	// result2 は if文の中身だけでしか使えない
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("greet 2")
	}

	num := 6
	if num%2 == 0 {
		fmt.Println("by 2") // by2
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 10, 20
	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}

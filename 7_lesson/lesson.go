package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	// close
	close(c)
}

// rangeとcloseで、結果がわかると随時channelに渡す処理

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	// range
	for i := range c {
		fmt.Println(i)
		// 実行結果
		// 1
		// 3
		// 6
		// 10
		// 15
	}
}

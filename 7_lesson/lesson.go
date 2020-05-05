package main

import "fmt"

// バッファ...記憶単位間のデータ転送において一時的にデータを記憶すること
func main() {
	// チャネルのバッファは2まで
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}

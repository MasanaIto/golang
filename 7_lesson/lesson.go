package main

import (
	"fmt"
	"time"
)

func goroutine(s string) {
	// defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	go goroutine("world")
	normal("hello")
	// wg.Wait()
}

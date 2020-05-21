package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	// 2020-05-21 23:24:45.647641 +0900 JST m=+0.000063878
	fmt.Println(t.Format(time.RFC3339))
	// 2020-05-21T23:24:45+09:00
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// 2020 May 21 23 24 45
}

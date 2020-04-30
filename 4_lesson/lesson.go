package main

import "fmt"

func thirdPartyConnectDB() {
	// プログラム強制終了する, 本来は自分でpanicを書かない
	panic("Unable to connect database!")
}

func save() {
	// panicをキャッチして、recoverによりプログラムの強制終了を起こさせない
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}

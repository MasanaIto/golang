package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2020-04-01", "2020-04-05", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 1)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 5)
	fmt.Println(mva)
}

package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	// test.log 実行結果
	// 2020/05/01 00:59:32 lesson.go:20: logging
	// 2020/05/01 00:59:32 lesson.go:21: string test

	log.Println("logging")              // 2020/05/01 00:50:13 logging
	log.Printf("%T %v", "test", "test") // 2020/05/01 00:50:13 string test

	// log.Fatalln("error")
}

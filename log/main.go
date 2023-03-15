package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ログの設定
// GC の際に書き込まれない事がある
// さらに複雑な事がしたければサードパーティー制のパッケージを使う
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	log.Println("start main")
	_, err := os.Open("nothing")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	fmt.Println("finished")

}

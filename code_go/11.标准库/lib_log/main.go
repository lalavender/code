package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("defer")
	log.Println("this a log!")
	//log.Fatalln("this is a fatal log")
	//fmt.Println("over")

	// 设置日志格式
	i := log.Flags()
	fmt.Println("i is :", i)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("this is a new log")

	s := log.Prefix()
	fmt.Println("s is :", s)
	log.SetPrefix(">>> ")
	log.Println("this is a prefix log")

	//写入文件
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln("打开日志文件异常")
	}
	log.SetOutput(f)
	log.Println("日志已写入文件")

	//new log
	logger := log.New(f, ">>> ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("this is a new logger")

}

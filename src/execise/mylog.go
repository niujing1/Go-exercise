package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger //记录所有日志
	Info *log.Logger //重要的信息
	Warning *log.Logger //需要注意的信息
	Error *log.Logger //非常严重的问题
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file :", err)
	}

	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate | log.Ltime |log.Lshortfile)

	Info = log.New(os.Stdout, "Info: ", log.Ldate | log.Ltime | log.Lshortfile)

	Warning = log.New(os.Stdout, "Warning: ", log.Ldate | log.Ltime | log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	Trace.Println("I have somthing to stadard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("something has failed")
}

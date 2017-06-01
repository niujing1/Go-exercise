package main

import (
	"log"
)

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main(){
	log.Println("message") //Println写到标准的日志记录器

	log.Fatalln("fatal message") // fatal Message在调用Println之后会接着调用os.Exit(1)

	log.Panicln("panic message") //Panicln在调用Println之后会调用panic()
}

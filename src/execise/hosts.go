package main

import (
	"net/http"
	"io"
	"log"
)


func helloServer(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello, go~")
}

func main() {
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}



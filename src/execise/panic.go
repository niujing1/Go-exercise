package main

import (
	"os"
	"fmt"
)

func main(){
	defer func(){
		fmt.Println("aaa")
	}()
	fmt.Println("bb")
	fmt.Println("cc")
	panic("fatal error")

	_, err := os.Create("/tmp/file")
	if err := recover(); err != nil {
		//panic(err)
		fmt.Println(err)
		fmt.Println("haha")
	}

	fmt.Println(err)
	//fmt.Println("aaaa")

}


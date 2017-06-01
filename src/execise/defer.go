package main

import "fmt"

func main(){
	defer fmt.Println("echo delay exec")
	defer fmt.Println("echo delay exec2")
	fmt.Println("write last but output first")
	fmt.Println("write last but output first~~")

	for i := 1; i < 4; i++{
		defer func(n int){
			fmt.Printf("\n%d\n",n)
		}(i*2) //得到6 4 2
	}
}

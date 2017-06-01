package main

import "fmt"

func main(){
	var con string

	switch con{
		default:
			fmt.Println("Unkonwn language")
		case "Python":
			fmt.Println("an interreted language")
		case "Go":
			fmt.Println("A compiled language")
	}
}

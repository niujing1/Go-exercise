package main

import (
	"fmt"
)


func main() {
	slice := []int{1, 2, 3, 4}
//迭代每一个元素
	for index, value := range slice {
		res, _ := fmt.Printf("Index: %d Value:%d\n", index, value)
		fmt.Println(res)
		fmt.Sprintf("Index: %d Value:%d\n", index, value)
	}
}

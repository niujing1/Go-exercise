package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//分配一个逻辑处理器来给调度器使用
	runtime.GOMAXPROCS(1)
	
	//wg用来等待goroutine完成，计数加2表示要等待2个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutine")

	//创建一个go routine打印输出1-10
	go func() {
		//在函数退出时调用Done来通知main函数工作完成
		defer wg.Done()

		//显示1-10三次
		for j := 1; j <= 3; j++ {
			for i := 1; i <= 10; i++ {
				fmt.Println(i)
			}
		}
	}()

	//声明一个匿名函数并创建一个goroutine
	go func() {
		//在函数退出调用时调用Done来通知main函数工作完成
		defer wg.Done()

		//显示a-i 3次
		for i := 0; i< 3; i++ {
			for char := 'a'; char < 'a'+10 ; char++ {
				fmt.Printf("%d ", char)
			}
		}
	}()

	//等待goroutine结束
	fmt.Println("waiting for end~~~")
	wg.Wait()

	fmt.Println("\nEnded..")
}

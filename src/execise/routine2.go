package main

import (
	"fmt"
	"sync"
	"runtime"
)


var wg sync.WaitGroup //等待程序完成

func main() {
	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(4)

	wg.Add(2) //设置等待的goroutine线程数为2

	//创建2个goroutine
	go printNum("A")
	go printNum("B")

	//等待goroutine结束
	fmt.Println("waiting for end~~~")
	wg.Wait() //它会一直阻塞到goroutine结束

	fmt.Println("Ending~~~")
}

func printNum(prefix string) {
	defer wg.Done()

	for i := 0; i < 5 ; i++ {
		fmt.Printf("%s%d ", prefix, i)
	}
	fmt.Println("Completed~")
}

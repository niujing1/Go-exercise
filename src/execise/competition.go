package main

//测试竞争状态

import (
	"fmt"
	"sync"
	"runtime"
	"sync/atomic"
)

//用来等待程序结束
var wg sync.WaitGroup
//var counter int
var counter int64

func main() {
	wg.Add(2)

	//创建2个goroutine
	go incCounter(1)
	go incCounter(2)

	//等待goroutine结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter用来增加包里counter变量的值
func incCounter(id int) {
	//在函数退出时调用Done来通知main函数工作已完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		//捕获counter的值
		//value := counter
		//安全的对counter+1
		atomic.AddInt64(&counter, 1)

		//goroutine从线程退出并放回队列
		runtime.Gosched()

		//增加本地value变量的值
		//value++

		//将值保存回counter
		//counter = value
	}
}

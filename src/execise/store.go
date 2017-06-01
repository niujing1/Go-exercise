package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"sync"
)

var (
	shutdown int64 //shutdown是通知正在执行的goroutine停止工作的标志
	wg sync.WaitGroup //wg 用来等待程序结束
)

func main() {
	wg.Add(2) //计数器加2表示要等待两个线程的结束

	//创建2个goroutine
	go doWork("A")
	go doWork("B")

	//给定goroutine执行的时间
	time.Sleep(1 * time.Second)

	//改停止工作了，安全的设置shutdown标志
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}


//doWork用来模拟之前执行工作的goroutine
//检测之前的shutdown标志来检测是否需要提前终止
func doWork(name string) {
	//在函数退出时调用Done来通知main函数是否工作完成
	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		//要停止工作了？
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shuttung down %s\n", name)
			break
		}
	}
}

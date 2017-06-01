package main

import (
	"fmt"
	"sync"
	"runtime"
)

var (
	counter int //所有变量都要更新这个值
	wg sync.WaitGroup //wg用来等待程序结束
	mutex sync.Mutex //用来定义一段临界区代码
)

func main() {
	wg.Add(2) //计数器加2表示等待2个goroutine

	//创建2个goroutine
	go incCounter(1)
	go incCounter(2)

	//等待goroutine结束
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

//使用互斥锁来同步并保证安全访问，增加包里的counter变量的值
func incCounter(id int) {
	defer wg.Done() //在函数退出时通知main函数，工作完成

	for count := 0; count < 2; count++ {
		//同一时刻，只允许一个goroutine进入
		mutex.Lock()
		{
			//捕获Counter的值
			value := counter

			//当前goroutine从线程退出，并放回队列
			runtime.Gosched()

			//增加本地value的值
			value++

			//将值保存会counter
			counter = value
		}

		mutex.Unlock()
		//释放锁，允许其他正在等待的goroutine进入临界区
	}
}

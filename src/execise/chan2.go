package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup //使用wg等待程序结束
)

func init() {
	//init会先于main函数执行
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//创建一个无缓冲的通道
	court := make(chan int)
	wg.Add(2)

	//启动两个选手
	go player("a", court)
	go player("b", court)

	//发球
	court <- 1

	//等待游戏结束
	wg.Wait()
}

//player模拟一个选手正在打球
func player(name string, court chan int) {
	defer wg.Done()

	for {
		//等待球被击打过来
		ball, ok := <- court

		if !ok {
			//如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		//选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court) //关闭通道，表示我们输了
			return
		}

		//显示击球数，并将击球数+1
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		//将球打向对手
		court <- ball
	}
}

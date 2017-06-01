package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

//用户定义一个用户类型
type user struct {
	name string
	email string
}

//notify 是使用指针接收值实现
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// main是程序创建的入口
func main(){
	//创建一个user类型的值，并发送通知
	u := user{"nj", "nj@126.com"}
	sendNotification(&u)//user does not implement notifier (notify method has pointer receiver)这里必须要使用&u否则就会报错，接收类型是指针，传参的时候要用指针
}

func sendNotification(n notifier) {
	n.notify()
}

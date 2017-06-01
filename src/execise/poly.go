package main

//多态的例子
import (
	"fmt"
)

//定义通知类行为的接口
type notifier interface {
	notify()
}

//user定义用户类型
type user struct {
	name string
	email string
}

//user使用指针接收者实现notifier接口
func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

//admin定义用户类型
type admin struct {
	name string
	email string
}

//admin使用指针类型接收者实现notifier接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main(){
	//创建一个user并传值给sendNotification
	bill := user {
		name: "bill",
		email: "bill@qq.com",
	}
	sendNotification(&bill)

	//创建一个admin，传递给sendNotification
	lisa := admin {
		name: "lisa",
		email: "lisa@qq.com",
	}
	sendNotification(&lisa)
}

func sendNotification(n notifier){
	n.notify()
}





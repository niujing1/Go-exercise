package main

//将一个内部类型嵌入到外部类型，内部类型和外部类型的关系

import (
	"fmt"
)

//加入一个interface
type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

//用纸张接收值实现notify
func (u *user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

//admin代表一个有权限的用户
type admin struct{
	user //嵌入类型
	level string
}

//admin实现自己的notify
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	/*
	bill := user{"bill", "bill@qq.com"}
	ad := admin{
		user: bill,
		level: "super",
	}
	*/
	//上边注释的可以写成
	ad := admin{
		user: user{
			name: "bill",
			email: "email",
		},
		level: "super",
	}

	//用于实现接口的内部类型的方法也会被提示到外部类型
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}

package main

//将一个内部类型嵌入到外部类型，内部类型和外部类型的关系

import (
	"fmt"
)

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

	//可以直接调用内部类型的方法
	ad.user.notify()
	//也可以把内部类型的方法提升到外部类型
	ad.notify()
}

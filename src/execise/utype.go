package main

import (
	"fmt"
)

//在程序中定义一个user类型
type user struct {
	name string
	email string
}

// notify 通过值接收者实现了一个方法
//使用值接收时，方法接收到的是值的一个副本
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

//changeEmail使用指针接收者实现了一个方法
//使用指针接收时，共享的是指针指向的值，而不是副本
func (u *user) changeEmail(email string) {
	u.email = email
}

//main函数是程序的主入口
func main() { 
	//user类型的值可以用来调用使用值接收者声明的方法
	bill := user{"bill", "bill@126.com"}
	bill.notify()

	//指向user类型的指针，也可以用来调用 使用值接收者声明的方法
	lisa := &user{"lisa", "lisa@qq.com"}//声明一个指向user变量的指针
	lisa.notify()//调用notify时，go编译器会先把指针解引用为(*lisa),就符合了notify方法的值定义

	//user类型的值可以用来调用使用值接收者声明的方法
	bill.changeEmail("bill@qq.com")
	bill.notify()

	//指向user类型的指针可以用来调用使用指针接收者声明的方法
	//使用值来调用指针类型时，go编译器会先转为(&lisa)指针
	lisa.changeEmail("lisa@126.com")
	lisa.notify()
}

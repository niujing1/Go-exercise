package main

import (
	"fmt"
	"runtime"
)

func init() {
	//代码包初始化
	fmt.Printf("Map: %v\n", m)
	info = fmt.Sprintf("OS: %s, Arch: %s\n ", runtime.GOOS, runtime.GOARCH)
}

//非局部变量，map声明并初始化
var m = map[int]string{1: "A", 2: "B", 3: "C"}
var info string

func main() {
	fmt.Println(info)
}

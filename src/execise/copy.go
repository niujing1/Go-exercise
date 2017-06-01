package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	//将字符串写入到Bufer
	b.Write([]byte("Hello"))

	//使用Fprintf将字符串拼接到Buffer
	fmt.Fprintf(&b, "world~")

	//将Buffer的内容输出到标准输出stdout
	io.Copy(os.Stdout, &b)
}

package main

//使用io.Reader io.Writer实现一个简单版本的curl
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//init会在main函数之前调用
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curl <url>")
		os.Exit(-1)
	}
}

//主函数
func main() {
	//从web服务器得到响应
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	//从Body copy到stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

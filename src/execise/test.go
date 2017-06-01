package main

import (
	"os/exec"
	"fmt"
	"io"
)

func main(){
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang")//创建一个exec.Cmd类型的值
	//创建一个可以获取此命令输出的管道：
	stdout0, err := cmd0.StdoutPipe() //stdoutPipe返回一个输出管道，代表把输出管道的值赋给stdout0它的类型是：io.ReaderCloser后者是一个扩展了io.Reader接口的接口类型，并定义了可关闭的数据读取行为
	if err != nil {
		fmt.Printf("Error: can't obtain the stdout pipe for command No.0: %s\n", err)
		return
	}

	//exec.Cmd上的Start方法可以启动命令
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The Command No.0 can't be start up: %s\n", err)
		return
	}

	//有了上述启动就可以调用Read方法来获取命令的输出了
	output0 := make([]byte, 30)
	n, err := stdout0.Read(output0)
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe: %s", err)
		return
	}

	fmt.Printf("read : %s\n", output0[:n])//若命令的输出小于output0的长度，n就代表实际输出 否则就等于output0的长度，证明数据一次没有读出全部

	var outputBuf0 bytes.Buffer

	for {
		tmpOutput := make([]byte, 5)
		n, err := stdout0.Read(tmpOutput)
		if err != nil {
			if err == io.EOF {
				break //数据读取完毕
			} else {
				fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
				return 
			}
		}

		if n > 0 {
			outputBuf0.Write(tmpOutput[:n])
		}
	}

	fmt.Printf("%s\n", outputBuf0.String())
}

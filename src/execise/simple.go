package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//准备从标准输入读入数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("An error Occured: %s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		//使用切片操作最后的 \n
		name := input[:len(input)-1]
		fmt.Printf("%s , what can I do for you~\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error Occured: %s\n", err)
			continue
		}

		input = input[:len(input)-1]
		input = strings.ToLower(input) //全部转为小写
		//fmt.Println(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			// 正常退出。
			os.Exit(0)
		default:
		        fmt.Println(len(input))
			fmt.Println("Sorry, I didn't catch you.")

		}

	}
}

package chatbot

import (
	"fmt"
	"strings"
)

//代表针对中文的延时级聊天机器人
type simpleCN struct{
	name string
	talk Talk
}

//NewSimpleCN 用于创建针对中文的演示级别的聊天机器人
func NewSimpleCN(name string, talk Talk) Chatbot {
	return &simpleCN{
		name: name,
		talk: talk,
	}
}

//Name 是Chatbot接口的实现的一部分
func (robot *simpleCN) Name() string{
	return robot.name
}

//Begin 是Chatbot接口的实现的一部分
func (robot *simpleCN) Begin() (string, error){
	return "请输入你的名字：", nil
}

// Hello是talk接口实现的一部分
func (robot *simpleCN) Hello(userName string) string{
	userName = strings.TrimSpace(userName)

	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}

	return fmt.Sprintf("你好，%s~ 我可以为你做些什么吗？", userName)
}

//talk 是Talk接口实现的一部分
func (robot *simpleCN) Talk(heart string) (saying string, end bool, er error){
	heart = strings.TrimSpace(heart)

	if robot.talk != nil {
		return robot.talk.Talk(heart)
	}

	switch heart {
		case "":
			return
		case "没有","再见":
			saying = "再见"
			end = true
			return
		default:
			saying = "对不起，没听懂你说的什么"
			return 
	}
}

func (robot *simpleCN) ReportError(err error) string{
	return fmt.Sprintf("发生了一个错误：%s\n", err)
}

//end是chatbot接口实现的一部分
func (robot *simpleCN) End() error {
	return nil
}

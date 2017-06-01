package chatbot

import (
	"fmt"
	"strings"
)

//simpleEN 代表针对英文的演示级聊天机器人
type simpleEN struct{
	name string
	talk Talk
}

//NewSimpleEN 用于创建针对英文的演示级聊天机器人
func NewSimpleEN(name string, talk Talk) Chatbot{
	return &simpleEN{
		name: name,
		talk: talk,
	}
}

// Name是chatbot接口实现的一部分
func (robot *simpleEN) Name() string {
	return robot.name
}

//Begin是Chatbot接口实现的一部分
func (robot *simpleEN) Begin() (string, error){
	return "please input your name:", nil
}

//Hello是Talk接口实现的一部分
func (robot *simpleEN) Hello(userName string) string{
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}

	return fmt.Sprintf("Hello, %s~, what can I do for you?", userName)
}

//Talk是Talk接口实现的额一部分
func (robot *simpleEN) Talk(heart string)(saying string, end bool, err error){
	heart = strings.TrimSpace(heart)
	if robot.talk != nil {
		return robot.talk.Talk(heart)
	}

	switch heart{
		case "" :
			return
		case "nothing", "bye":
			saying = "Bye"
			end = true
			return
		default:
			saying = "Sorry, I did't catch you."
			return
	}
}

//ReportError是chatbot接口实现的一部分
func (robot *simpleEN) ReportError(err error) string{
	return fmt.Sprintf("An error occured: %s\n", err)
}

//End是Chatbot接口实现的一部分
func (robot *simpleEN) End() error{
	return nil
}


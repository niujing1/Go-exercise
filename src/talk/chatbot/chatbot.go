package chatbot

import "errors"

//Talk 定义了聊天的接口类型

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}


//chatbot 定义了聊天机器人的类型
type Chatbot interface{
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var(
	ErrInvalidChatName = errors.New("Invalid chat name")
	ErrInvalidChatbot = errors.New("Invalid chatbot name")
	ErrExistingChatbot = errors.New("Existing chatbot") //同名机器人重复注册错误
)

var chatbotMap = map[string]Chatbot{}

//Register用于注册聊天机器人

func Register(chatbot Chatbot) error{
	if chatbot == nil {
		return ErrInvalidChatbot
	}

	name := chatbot.Name()
	if name == "" {
		return ErrInvalidChatbot
	}

	if _, ok := chatbotMap[name]; ok{
		return ErrExistingChatbot
	}
	chatbotMap[name] = chatbot
	return nil
}

//Get用于获取指定名称的聊天机器人
func Get(name string) Chatbot{
	return chatbotMap[name]
}

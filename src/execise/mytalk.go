type myTalk string

func (talk *myTalk) Hello(userName string) string{
	//...
}

func (talk *myTalk) Talk(heart string)(saying string, end bool, err error){
	//...
}

var talk Talk = new (myTalk) //myTalk并不是接口Talk的实现，*myTalk才是 接口类型的变量可以被赋予任何实现类型的值

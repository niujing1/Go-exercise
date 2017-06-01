package main

import(
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	//Author []string
	Author []str `json:"Authort"`
	Price  float64
	Ispublished bool
}

//{}->struct []->[]string `json:"x"`->映射字段 
type str struct {
	A string
	B string
	C string
}

func main() {
	//若b的值是json，在下边使用的时候就要用 json.Unmarshal([]byte(b)),
	//先把string转换为切片
	b := []byte(`{
		  "title" : "go programming",
		  "Authort" : [{"a":"aa", "b":"bb", "c":"cc"}],
		  "Price" : 99,
		  "IsPublished" : true
		}`)

	//先创建一个目标类型的实例对象，用于存放解码后的值
	var book Book
	err := json.Unmarshal(b, &book) //反序列化b，将结果存储到v
	if err != nil {
		fmt.Println("error in tranlating,", err.Error())
		return
	}
	fmt.Println(book)
}



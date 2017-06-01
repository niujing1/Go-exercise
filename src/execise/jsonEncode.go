package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main(){
	//创建一个保存键值对的映射
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home" : "aaaaa",
		"phone" : "123132323",
	}

	//将这个映射序列化到json
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Println("Error : ", err)
		return
	}

	fmt.Println(string(data))
}

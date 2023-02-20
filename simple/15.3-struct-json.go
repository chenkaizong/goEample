package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City   string `json:"classname"` //可以通过备注修改json的返回名
	Detail string
}

func (address Address) getCity() string {
	return address.City
}

type Person struct {
	Address //嵌套,表示具有city 和detail 字段
	Name    string
	Age     int64
	hobby   [6]string
	map1    map[string]string
	city    string
}

func main() {
	person := Person{
		Name:    "name",
		Age:     16,
		hobby:   [6]string{"WO", "LA", "BAI"},
		map1:    map[string]string{"key1": "value1", "key2": "value2"},
		Address: Address{City: "福州", Detail: "baid"},
		city:    "开发", //与嵌套属性重复时返回自身的
	}

	json_str, _ := json.Marshal(person)

	fmt.Printf("%s\n", json_str) //首字母必须是大写才能输出

	var person2 Person
	unjson_str := `{"City":"福州","Detail":"baid","Name":"name","Age":16}`
	error := json.Unmarshal([]byte(unjson_str), &person2)
	if error == nil {
		fmt.Printf("%#v", person2)
	}
}

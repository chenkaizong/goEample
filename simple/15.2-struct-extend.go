package main

import "fmt"

type Address struct {
	city   string
	detail string
}

func (address Address) getCity() string {
	return address.city
}

type Person struct {
	Address //嵌套,表示具有city 和detail 字段
	name    string
	age     int64
	hobby   [6]string
	map1    map[string]string
	city    string
}

func main() {
	person := Person{
		name:    "name",
		age:     16,
		hobby:   [6]string{"WO", "LA", "BAI"},
		map1:    map[string]string{"key1": "value1", "key2": "value2"},
		Address: Address{city: "福州", detail: "baid"},
		city:    "开发", //与嵌套属性重复时返回自身的
	}

	fmt.Println(person.city)   //开发
	fmt.Println(person.detail) //baid

	person.Address.city = "北京"
	fmt.Printf("%+v\n", person)
	fmt.Println(person.getCity()) //嵌套即可实现继承

}

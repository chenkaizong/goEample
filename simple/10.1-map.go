package main

import "fmt"

func main() {
	// 1、使用make方式
	var map1 = make(map[int]int)
	map1[1] = 20
	map1[2] = 33

	var map2 = make(map[string]string)
	map2["a"] = "aaa"
	map2["b"] = "bbb"

	// 2、默认方式
	var user = map[string]string{
		"a": "aaa",
		"b": "bbb",
	}
	fmt.Println(user)

	// 3、短标签方式
	user2 := map[string]string{
		"a": "aaa",
		"b": "bbb",
	}

	fmt.Println(user2)

	// 4、遍历
	for key, value := range user2 {
		fmt.Println(key, value)
	}

	// 5、查找key是否存在
	v, ok := user2["a"]
	fmt.Println(v, ok)

	// 6、删除
	delete(user2, "a")

}

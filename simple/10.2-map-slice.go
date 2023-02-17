package main

import "fmt"

func main() {
	// 3表示长度，6表示容量
	var list = make([]map[string]string, 3, 6)
	list[0] = map[string]string{"username": "chen", "age": "20"}
	list[1] = map[string]string{"username": "chen2", "age": "40"}
	list[2] = map[string]string{"username": "chen3", "age": "40"}

	fmt.Println(list)

	var userinfo = make([]map[string][]string, 3, 6)
	userinfo[0] = map[string][]string{"username": {"chen", "age", "20"}}

	fmt.Println(userinfo)
	userinfo2 := userinfo
	userinfo2[0]["username"] = []string{"chen", "age", "2220"}
	fmt.Println(userinfo)

}

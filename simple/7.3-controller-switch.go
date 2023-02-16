package main

import "fmt"

func main() {
	// 1.switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的，从上到下逐一测试，知道匹配为止
	// 2.匹配项后面也不需要再加break
	a := "哈"
	switch a {
	case "哈":
		fmt.Println("1")
	case "哈哈":
		fmt.Println("2")
	// 这里可以有多个case语句
	default:
		fmt.Println("3")
	}

}

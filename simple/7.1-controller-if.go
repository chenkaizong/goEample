package main

import "fmt"

func main() {
	age := 75
	if age > 70 {
		fmt.Println("大于75")
	} else if age > 60 {
		fmt.Println("大于60")
	} else {
		fmt.Println("小于60")
	}

	if age1 := 75; age1 > 70 { //age1只在当前作用域有用
		fmt.Println("大于75")
	}

}

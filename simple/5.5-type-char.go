package main

import "fmt"

func main() {
	//1、 单引号，这不是一个基础类型 uint8
	var a byte = 'a'
	fmt.Printf("%T,%d", a, a) //uint8,97

	//2、 %c 可以输出对应字符
	var b uint = 99
	fmt.Println(fmt.Sprintf("%c", b))

	//3、 utf8 是一个 rune 类型  是 int32
	var c rune = '国'
	fmt.Printf("%d,%T\n", c, c) //22269,int32

	// 4、遍历
	str := "我a是b中c国d人"
	//循环遍历,这样是乱码
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c", str[i])
	// }
	// 使用range才是正确的
	for _, k := range str {
		fmt.Printf("%c", k)
	}

	// 5、修改字符串,直接修改是不行的， 只能转换成rune数组修改再转成字符串
	runeStr := []rune(str)
	runeStr[3] = '哈'
	fmt.Printf("%s", string(runeStr))

}

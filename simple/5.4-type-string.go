package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	// 1、字符串
	var str = "haha"
	fmt.Printf("type is %T,size %d\n", str, unsafe.Sizeof(str))
	// 2、转义 \n,\t 等，反转义 \\ 代表 \   \" 表示 \

	// 3、反引号定义多行字符串

	var str2 = `asdfasdf
	asdfasdf
	asdfasdf
	"asdfasdf"`

	fmt.Printf("%s \n type is %T,size %d\n", str2, str2, unsafe.Sizeof(str2))

	//4、常用字符串方法
	str3 := "我"
	str4 := "是"
	str5 := "我 - 是 - 中"

	//4.1 一个汉字三个字节
	fmt.Println(len(str3))
	//4.2 拼接
	fmt.Println(str3 + str4)                      //我是
	fmt.Println(fmt.Sprintf("%s,%s", str3, str4)) //S开头的表示返回值
	//4.3分割
	str6 := strings.Split(str5, "-")
	fmt.Println(str6)
	//4.4连接
	fmt.Println(strings.Join(str6, "*"))
	str7 := []string{"php", "golang", "java"}
	fmt.Println(strings.Join(str7, "*"))

	// 4.5 是否包含
	fmt.Println(strings.Contains(str5, "我"))

	// 4.6 是否有前缀或后缀
	fmt.Println(strings.HasPrefix(str5, "我"))
	fmt.Println(strings.HasSuffix(str5, "我"))

	// 4.7 查看出现的位置
	strings.Index(str5, str4)
	strings.LastIndex(str5, str4)

}

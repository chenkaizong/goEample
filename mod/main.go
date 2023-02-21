package main

import (
	"fmt"

	"example.com/m/calc"
	"github.com/shopspring/decimal"
)

// 1、 go mod init example.com/m
// 2、 新建文件夹即可创建包，通过引入example.com/m/calc，就可以用首字母大写来调用
// 3、 第三方包 go get github.com/shopspring/decimal
// 4、 go vendor 表示把mod 放到vendor文件夹，优先从这里查找模块
func main() {
	// fmt.Println(calc.aaa) 这里肯定是报错的，不能是小写的
	fmt.Println(calc.Add(16, 20))

	a := decimal.NewFromFloat32(64.3)
	b := decimal.NewFromInt(43)
	fmt.Println(a.Add(b))

}

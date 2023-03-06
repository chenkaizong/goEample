package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 第一种方法
func Open(path string) {

	file, err := os.Open(path)
	defer file.Close() //记得defer close
	if err != nil {
		fmt.Println("打开文件失败")
	}
	var strSlice []byte //存放所有读取内容

	for {
		var tmp = make([]byte, 128) //一次性读取内容
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("执行完毕")
			break
		} else if err != nil {
			fmt.Println("读取内容失败")
			break
		}
		strSlice = append(strSlice, tmp[:n]...)
	}

	fmt.Println(string(strSlice))
}

// 第二种方法
func bufioOpen(path string) {
	file, err := os.Open(path)
	defer file.Close() //记得defer close
	if err != nil {
		fmt.Println("打开文件失败")
	}
	reader := bufio.NewReader(file)
	var strResult string
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取错误")
			break
		}
		strResult += str
	}

	fmt.Println(strResult)

}

// 3一次性
func utilOpen(path string) {
	fileByte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("错误")
	}
	fmt.Print(string(fileByte))
}

// https://zhuanlan.zhihu.com/p/260661233
func main() {
	// Open("23demo/main.txt")
	// bufioOpen("23demo/main.txt")
	utilOpen("23demo/main.txt")

}

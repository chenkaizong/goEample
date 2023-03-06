package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 第一种方法
// os.O_RDONLY、 只读
// os.O_WRONLY   只写
// os.O_RDWR	读写
// os.O_APPEND：当向文件中写入内容时，把新内容追加到现有内容的后边。
// os.O_CREATE：当给定路径上的文件不存在时，创建一个新文件。
// os.O_EXCL：需要与os.O_CREATE一同使用，表示在给定的路径上不能有已存在的文件。
// os.O_SYNC：在打开的文件之上实施同步 I/O。它会保证读写的内容总会与硬盘上的数据保持同步。
// os.O_TRUNC：如果文件已存在，并且是常规的文件，那么就先清空其中已经存在的任何内容。

func Write(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("错误")
	}
	var writeBtye []byte
	writeBtye = []byte("bbbb")
	file.Write(writeBtye)

}

// 第二种方法
func bufioWrite(path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println("错误")
	}

	write := bufio.NewWriter(file)
	write.WriteString("basdijf;aoisdjfaisdf;ajisdofa;sdifj")
	write.Flush() //缓存 需要调用

}

// 第三种3 一次性
func utilOpen(path string) {
	err := ioutil.WriteFile(path, []byte("baidoaisdj;foaisjdfasdf"), 0600)
	if err != nil {
		fmt.Println("错误")
	}

}

// https://zhuanlan.zhihu.com/p/260661233
func main() {
	// Write("23demo/test.txt")
	// bufioWrite("23demo/test.txt")
	utilOpen("23demo/test2.txt")

}

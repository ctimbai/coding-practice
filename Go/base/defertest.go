/******************************************************************************
 *	File Name: defertest.go
 *	Author: 公众号: 5分钟学Go
 *****************************************************************************/

package main

import (
	"fmt"
	"os"
)

func main() {
	GetFileContentDefer()
	DeferTest()
}

func DeferTest() {
	defer FirstCall()
	defer SecondCall()
	fmt.Println("defer test in main")
}

func FirstCall() {
	fmt.Println("first call defer")
}

func SecondCall() {
	fmt.Println("second call defer")
}

// 没有使用defer
func GetFileContentNoDefer() int {
	file, _ := os.Open("./file.txt")
	data := make([]byte, 100)
	file.Read(data)
	dataStr := string(data)
	fmt.Println(dataStr)
	if dataStr == "hello" {
		file.Close()
		return 1
	} else {
		file.Close() // 多次file.Close()
		return 0
	}
}

// 使用defer
func GetFileContentDefer() int {
	file, _ := os.Open("./file.txt")
	defer file.Close() // 只用一次
	data := make([]byte, 100)
	file.Read(data)
	dataStr := string(data)
	fmt.Println(dataStr)
	if dataStr == "hello" {
		return 1
	} else {
		return 0
	}
}

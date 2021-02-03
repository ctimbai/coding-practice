/******************************************************************************
 *	File Name: localtest.go
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	a int
	b float32
	c string
}

func main() {
	var i int
	var f float32
	var s string
	var arr = [...]int{1, 2, 3, 4}
	var slice = []int{1, 2, 3, 4}
	var user User

	fmt.Println(unsafe.Sizeof(i))     // 8
	fmt.Println(unsafe.Sizeof(f))     // 4
	fmt.Println(unsafe.Sizeof(s))     // 16 string包括指向字符串的指针域和字符串长度域
	fmt.Println(unsafe.Sizeof(arr))   // 32 数组固定长度
	fmt.Println(unsafe.Sizeof(slice)) // 24 切片包括指向切片的指针域、容量域和长度域
	fmt.Println(unsafe.Sizeof(user))  // 32

	fmt.Println(unsafe.Alignof(i))       // 8
	fmt.Println(unsafe.Alignof(f))       // 4
	fmt.Println(unsafe.Alignof(user))    // 8
	fmt.Println(unsafe.Offsetof(user.b)) // 8
	fmt.Println(unsafe.Offsetof(user.c)) // 16
}

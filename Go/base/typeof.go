/******************************************************************************
 *	File Name: typeof.go
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "hello"
	var b = 1.2
	var c int32 = 1
	var d int = 1
	if int(c) == d {
		fmt.Println("true") // true
	} else {
		fmt.Println("false")
	}
	var e = int(c) + d
	fmt.Println(e) // 2
	var r rune = 12
	var t byte = 2
	var i int32 = r + c
	fmt.Println(i)                 // 13
	fmt.Println(reflect.TypeOf(r)) // int32
	fmt.Println(reflect.TypeOf(t)) // uint8
	fmt.Println(reflect.TypeOf(a)) // string
	fmt.Println(reflect.TypeOf(b)) // float64
}

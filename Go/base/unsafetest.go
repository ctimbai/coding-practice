/******************************************************************************
 *	File Name: localtest.go
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

package main

import (
    "fmt"
    "unsafe"
)

var i int = 10
var j = 2

func main() {
    i := 1
    fmt.Println(i)

    a := "hello"
    fmt.Println(unsafe.Sizeof(a))
}

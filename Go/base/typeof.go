/******************************************************************************
 *	File Name: typeof.go
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

package main

import (
    "reflect"
    "fmt"
)

func main() {
    var a  = "hello"
    var b = 1.2
    var c int32 = 1
    var d int = 1
    if int(c) == d {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
    var e = int(c) + d
    fmt.Println(e)
    var r rune = 12
    var t byte = 2
    var i int32 = r + c
    fmt.Println(i)
    fmt.Println(reflect.TypeOf(r))
    fmt.Println(reflect.TypeOf(t))
    fmt.Println(reflect.TypeOf(a))
    fmt.Println(reflect.TypeOf(b))
}

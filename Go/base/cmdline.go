/******************************************************************************
 *	File Name: cmdline.go
 *	Author: 公众号: CloudDeveloper
 *	Created Time: Sat 26 Jan 2019 04:47:25 PM CST
 *****************************************************************************/

package main

import (
    "fmt"
    "strings"
    "os"
)

/*
func main() {
    var s, sep string
    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        fmt.Println(s)
        sep = " "
    }
    fmt.Println(s)
}
*/

func main() {
    var s, sep string
    fmt.Println(os.Args[0])
    for i, arg := range os.Args[0:] {
        s += sep + arg 
        fmt.Println(i, s)
        sep = " "
    }
    fmt.Println(s)
    fmt.Println(strings.Join(os.Args[0:], " "))
    fmt.Println(os.Args[0:])
}

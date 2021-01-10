/******************************************************************************
 *	File Name: swap.go
 *	Author: 公众号: Linux云计算网络
 *****************************************************************************/

package main

import "fmt"

func swap(p, q *int) {
    t := *p
    *p = *q
    *q = t
}

func main() {
    p, q := 2, 5
    swap(&p, &q)
    fmt.Println(p, q)

    x := new(int)
    y := new(int)
    fmt.Println(x, y)
    fmt.Println(x == y)

    new, old := 1, 2
    fmt.Println(new, old)
}

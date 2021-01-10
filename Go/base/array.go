package main

import (
	"fmt"
	"strings"
)

func main()  {
	var a [5]int // 数组的长度也是其类型的一部分
	var b []int  // 这种写法就不是数组，是切片
	fmt.Println(a, b)

	// 数组的初始化
	a1 := [3]int{1,2,3}
	a2 := [2]string{"hello", "world"}
	fmt.Println(a1, a2)

	// 字符串和字符串数组
	s1 := "hellogo"
	fmt.Println(s1)
	fmt.Printf("%c, %c\n", s1[0], s1[1])

	// 使用 strings.Fields 将字符串转换为字符串数组
	s2 := strings.Fields(s1)
	fmt.Println(s2[0])
}

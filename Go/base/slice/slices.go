package main

import (
	"fmt"
	"strings"
)

func main() {
	var a []int // 切片可以看成是大小动态变化的数组
	fmt.Println(a)

	// 切片初始化
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := []string{
		"hello",
		"world", // 不能漏掉这里的,
	}
	fmt.Println(a1, a2)

	// 结构体切片
	a3 := []struct{
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(a3)

	// 切片赋值和范围
	b1 := a1[:] // 由上界和下界包含，是一个半开区间，包括第一个元素，但排除最后一个元素，下界默认是0，上界默认是切片的长度
	fmt.Println(b1)
	b1 = a1[:6] // 包含前 6 个元素
	fmt.Println(b1)
	b1 = a1[3:] // 包含从 4 开始到结尾的元素
	fmt.Println(b1)
	b1 = a1[1:4] // 包含 2,3,4 三个元素
	fmt.Println(b1)

	// 改变切片元素，影响源切片
	b2 := a1[1:6]
	fmt.Println(b2)
	b2[0] = 11 // 切片并不存储任何数据，它只是是底层数组的一个引用，更改切片的元素会修改其底层数组中对应的元素
	fmt.Println(a1)
	fmt.Println(b1) // 与它共享底层数组的切片也会观测到这些修改

	// 切片的空间变化
	fmt.Println(len(a1)) // len() 是切片的长度
	fmt.Println(cap(a1)) // cap() 是切片的空间，当长度增加到和空间相同时，空间会增加到当前长度的 2 倍

	// append 增加元素查看 len 和 cap 的变化，可以看到是 2 倍的变化
	var a4 []int

	for i := 0; i < 10; i++ {
		a4 = append(a4, i)
		fmt.Println(a4, len(a4), cap(a4))
	}

	// nil 切片
	var a5 []int // nil 切片的长度和容量为0且没有底层数组
	if a5 == nil {
		fmt.Println("a5 == nil")
	}

	// make 创建切片
	a6 := make([]int, 5)
	printSlice(a6)
	a6 = make([]int, 5, 6)
	printSlice(a6)
	a7 := a6[:2]
	printSlice(a7)
	a7[1] = 3
	printSlice(a7)
	a8 := a7[2:5]
	printSlice(a8)

	// 切片可以包含任何类型，甚至其他切片，也即切片的切片
	str := [][]string{
		[]string{"1", "2"},
		[]string{"3", "4"},
		[]string{"5", "6"},
	}
	fmt.Println(str, str[1])
	for i := 0; i < len(str); i ++ {
		fmt.Printf("%s\n", strings.Join(str[i], " "))
	}

	// range 遍历切片
	for i, v := range a1 {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}

	// i 可以省略
	for _, v := range a1 {
		fmt.Printf("%d ", v)
	}

	// value 可以忽略
	for i := range a1 {
		fmt.Print(i)
	}
}

func printSlice(s []int) {
	fmt.Printf("%v len = %d cap = %d\n", s, len(s), cap(s))
}

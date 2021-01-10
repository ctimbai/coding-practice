package main

import "fmt"

func main()  {
	// 指针指向数值
	i, j := 2, 3

	p := &i
	*p = 10
	fmt.Println(*p)

	p = &j
	fmt.Println(*p)

	// 指针指向字符串
	s := "hello"
	q := &s

	fmt.Println(*q)

	// 指针指向slice
	a := []int{1,2,3}
	pa := &a
	fmt.Println(*pa)
}